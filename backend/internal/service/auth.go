package service

import (
	"errors"
	"strconv"

	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/auth"
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/mapper"
	"github.com/alifdwt/kutipanda-backend/internal/repository"
	customAuth "github.com/alifdwt/kutipanda-backend/pkg/auth"
	"github.com/alifdwt/kutipanda-backend/pkg/hashing"
	"github.com/alifdwt/kutipanda-backend/pkg/logger"
)

type authService struct {
	user   repository.UserRepository
	hash   hashing.Hashing
	log    logger.Logger
	token  customAuth.TokenManager
	mapper mapper.UserMapping
}

func NewAuthService(user repository.UserRepository, hash hashing.Hashing, log logger.Logger, token customAuth.TokenManager, mapper mapper.UserMapping) *authService {
	return &authService{user, hash, log, token, mapper}
}

func (s *authService) Register(input *auth.RegisterRequest) (*responses.UserResponse, error) {
	hashing, err := s.hash.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	input.Password = hashing

	user, err := s.user.CreateUser(input)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(user)

	return mapper, nil
}

func (s *authService) Login(input *auth.LoginRequest) (*responses.Token, error) {
	res, err := s.user.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("error while get user: " + err.Error())
	}

	err = s.hash.ComparePassword(res.Password, input.Password)
	if err != nil {
		return nil, errors.New("error while compare password: " + err.Error())
	}

	accessToken, err := s.createAccessToken(res.ID, res.Email, res.Username, res.FullName)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.createRefreshToken(res.ID, res.Email, res.Username, res.FullName)
	if err != nil {
		return nil, err
	}

	return &responses.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) RefreshToken(req auth.RefreshTokenRequest) (*responses.Token, error) {
	res, err := s.token.ValidateToken(req.RefreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	userID, err := strconv.Atoi(res)
	if err != nil {
		return nil, errors.New("invalid refresh token type")
	}

	user, err := s.user.GetUserById(userID)
	if err != nil {
		return nil, errors.New("invalid refresh token user")
	}

	newToken, err := s.createAccessToken(user.ID, user.Email, user.Username, user.FullName)
	if err != nil {
		return nil, errors.New("invalid access token")
	}

	newRefreshToken, err := s.createRefreshToken(user.ID, user.Email, user.Username, user.FullName)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	return &responses.Token{
		AccessToken:  newToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *authService) createAccessToken(id int, email string, username string, fullname string) (string, error) {
	res, err := s.token.NewJwtToken(id, email, username, fullname, "access")
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *authService) createRefreshToken(id int, email string, username string, fullname string) (string, error) {
	res, err := s.token.NewJwtToken(id, email, username, fullname, "refresh")
	if err != nil {
		return "", err
	}

	return res, nil
}
