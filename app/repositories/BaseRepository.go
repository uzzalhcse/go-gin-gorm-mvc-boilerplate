package repositories

import (
	"errors"
	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (repo *BaseRepository) GetAll(dest interface{}) error {
	result := repo.db.Find(dest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *BaseRepository) GetByID(dest interface{}, id string) error {
	result := repo.db.First(dest, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *BaseRepository) GetWhere(dest interface{}, condition string, args ...interface{}) error {
	result := repo.db.Where(condition, args...).Find(dest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *BaseRepository) GetFirst(dest interface{}, condition string, args ...interface{}) error {
	result := repo.db.Where(condition, args...).First(dest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *BaseRepository) Create(data interface{}) error {
	result := repo.db.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *BaseRepository) Update(data interface{}) error {
	result := repo.db.Save(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no record was updated")
	}
	return nil
}

func (repo *BaseRepository) Delete(data interface{}) error {
	result := repo.db.Delete(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no record was deleted")
	}
	return nil
}

// SetDB sets the GORM database connection for the repository
func (repo *BaseRepository) SetDB(db *gorm.DB) {
	repo.db = db
}

// GetDB returns the GORM database connection used by the repository
func (repo *BaseRepository) GetDB() *gorm.DB {
	return repo.db
}
