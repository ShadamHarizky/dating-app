package service

import (
	"testing"

	"github.com/ShadamHarizky/dating-app/domain/entity"
	"github.com/stretchr/testify/assert"
)

var (
	saveUserRepo                func(*entity.User) (*entity.User, map[string]string)
	getUserRepo                 func(userId uint64) (*entity.User, error)
	getUsersRepo                func(excludeId uint64) ([]entity.User, error)
	getUserEmailAndPasswordRepo func(*entity.User) (*entity.User, map[string]string)
	purchasePremium             func(userId uint64) (*entity.User, error)
)

type fakeUserService struct{}

func (u *fakeUserService) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return saveUserRepo(user)
}
func (u *fakeUserService) GetUser(userId uint64) (*entity.User, error) {
	return getUserRepo(userId)
}
func (u *fakeUserService) GetUsers(userId uint64) ([]entity.User, error) {
	return getUsersRepo(userId)
}
func (u *fakeUserService) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return getUserEmailAndPasswordRepo(user)
}

func (u *fakeUserService) PurchasePremium(userId uint64) (*entity.User, error) {
	return purchasePremium(userId)
}

var userServiceFake UserServiceInterface = &fakeUserService{} //this is where the real implementation is swap with our fake implementation

func TestSaveUser_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	saveUserRepo = func(user *entity.User) (*entity.User, map[string]string) {
		return &entity.User{
			ID:        1,
			FirstName: "victor",
			LastName:  "steven",
			Email:     "steven@example.com",
			Password:  "password",
		}, nil
	}
	user := &entity.User{
		ID:        1,
		FirstName: "victor",
		LastName:  "steven",
		Email:     "steven@example.com",
		Password:  "password",
	}
	u, err := userServiceFake.SaveUser(user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.FirstName, "victor")
	assert.EqualValues(t, u.LastName, "steven")
	assert.EqualValues(t, u.Email, "steven@example.com")
}

func TestGetUser_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	getUserRepo = func(userId uint64) (*entity.User, error) {
		return &entity.User{
			ID:        1,
			FirstName: "victor",
			LastName:  "steven",
			Email:     "steven@example.com",
			Password:  "password",
		}, nil
	}
	userId := uint64(1)
	u, err := userServiceFake.GetUser(userId)
	assert.Nil(t, err)
	assert.EqualValues(t, u.FirstName, "victor")
	assert.EqualValues(t, u.LastName, "steven")
	assert.EqualValues(t, u.Email, "steven@example.com")
}

func TestGetUsers_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	getUsersRepo = func(uint64) ([]entity.User, error) {
		return []entity.User{
			{
				ID:        1,
				FirstName: "victor",
				LastName:  "steven",
				Email:     "steven@example.com",
				Password:  "password",
			},
			{
				ID:        2,
				FirstName: "kobe",
				LastName:  "bryant",
				Email:     "kobe@example.com",
				Password:  "password",
			},
		}, nil
	}
	users, err := userServiceFake.GetUsers(uint64(2))
	assert.Nil(t, err)
	assert.EqualValues(t, len(users), 2)
}

func TestGetUserByEmailAndPassword_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	getUserEmailAndPasswordRepo = func(user *entity.User) (*entity.User, map[string]string) {
		return &entity.User{
			ID:        1,
			FirstName: "victor",
			LastName:  "steven",
			Email:     "steven@example.com",
			Password:  "password",
		}, nil
	}
	user := &entity.User{
		ID:        1,
		FirstName: "victor",
		LastName:  "steven",
		Email:     "steven@example.com",
		Password:  "password",
	}
	u, err := userServiceFake.GetUserByEmailAndPassword(user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.FirstName, "victor")
	assert.EqualValues(t, u.LastName, "steven")
	assert.EqualValues(t, u.Email, "steven@example.com")
}
