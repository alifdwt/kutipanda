package repository

import (
	"fmt"

	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/song"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type songRepository struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) *songRepository {
	return &songRepository{db: db}
}

func (r *songRepository) CreateSong(userId int, request song.CreateSongRequest) (*models.Song, error) {
	var songModel models.Song

	db := r.db.Model(&songModel)

	songModel.Title = request.Title
	songModel.Lyrics = request.Lyrics
	songModel.ReleaseDate = request.ReleaseDate
	songModel.AlbumImageUrl = request.AlbumImageUrl
	songModel.CountryID = request.CountryID

	songModel.UserID = userId
	songModel.Slug = fmt.Sprintf("%s-%d", slug.Make(request.Title), request.ReleaseDate.Year())

	result := db.Debug().Create(&songModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &songModel, nil
}

func (r *songRepository) GetSongAll(limit int, offset int, order string) (*[]models.Song, error) {
	var songs []models.Song

	db := r.db.Model(&songs)

	result := db.Debug().
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&songs)
	if result.Error != nil {
		return nil, result.Error
	}

	return &songs, nil
}

func (r *songRepository) GetSongById(id int) (*models.Song, error) {
	var songModel models.Song

	db := r.db.Model(&songModel)

	result := db.Debug().Where("id = ?", id).First(&songModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &songModel, nil
}

func (r *songRepository) GetSongBySlug(slug string) (*models.Song, error) {
	var songModel models.Song

	db := r.db.Model(&songModel)

	result := db.Debug().Where("slug = ?", slug).First(&songModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &songModel, nil
}

func (r *songRepository) FindSongBySlug(slug string, limit int, offset int, order string) (*[]models.Song, error) {
	var songs []models.Song

	db := r.db.Model(&songs)

	result := db.Debug().
		Where("slug ILIKE ?", slug).
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&songs)
	if result.Error != nil {
		return nil, result.Error
	}

	return &songs, nil
}

func (r *songRepository) UpdateSong(userId int, songId int, updatedSong song.UpdateSongRequest) (*models.Song, error) {
	var song models.Song

	db := r.db.Model(&song)

	checkSongById := db.Debug().Where("id = ? AND user_id = ?", songId, userId).First(&song)
	if checkSongById.Error != nil {
		return nil, checkSongById.Error
	}

	song.Title = updatedSong.Title
	song.Lyrics = updatedSong.Lyrics
	song.ReleaseDate = updatedSong.ReleaseDate
	song.AlbumImageUrl = updatedSong.AlbumImageUrl
	song.CountryID = updatedSong.CountryID

	song.Slug = fmt.Sprintf("%s-%d", slug.Make(updatedSong.Title), updatedSong.ReleaseDate.Year())

	updateSong := db.Debug().Updates(&song)
	if updateSong.Error != nil {
		return nil, updateSong.Error
	}

	return &song, nil
}

func (r *songRepository) DeleteSong(songId int) (*models.Song, error) {
	var songModel models.Song

	db := r.db.Model(&songModel)

	checkSongById := db.Debug().Where("id = ?", songId).First(&songModel)
	if checkSongById.Error != nil {
		return nil, checkSongById.Error
	}

	deleteSong := db.Debug().Delete(&songModel)
	if deleteSong.Error != nil {
		return nil, deleteSong.Error
	}

	return &songModel, nil
}
