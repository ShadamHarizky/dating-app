package database

import (
	"fmt"

	"github.com/ShadamHarizky/dating-app/domain/repository"

	"github.com/ShadamHarizky/dating-app/domain/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

// closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

// This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.User{}).Error
}
