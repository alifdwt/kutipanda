package mapper

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/models"
)

type userMapper struct {
}

func NewUserMapper() *userMapper {
	return &userMapper{}
}

func (m *userMapper) ToUserResponse(request *models.User) *responses.UserResponse {
	return &responses.UserResponse{
		ID:              request.ID,
		Email:           request.Email,
		Username:        request.Username,
		FullName:        request.FullName,
		ProfileImageURL: request.ProfileImageURL,
		Description:     request.Description,
		// BirthDate:       request.BirthDate,
	}
}

func (m *userMapper) ToUserResponses(requests *[]models.User) []responses.UserResponse {
	var userResponses []responses.UserResponse
	for _, request := range *requests {
		userResponses = append(userResponses, *m.ToUserResponse(&request))
	}

	if len(userResponses) > 0 {
		return userResponses
	}

	return []responses.UserResponse{}
}
