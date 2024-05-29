package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	recordSwipeService       func(userID, profileID int, direction string) error
	getSwipeCountService     func(userID int, direction string) (int, error)
	getSwipedProfilesService func(ctx context.Context, userID int, direction string) ([]string, error)
)

type fakeSwipeService struct{}

func (u *fakeSwipeService) RecordSwipe(userID, profileID int, direction string) error {
	return recordSwipeService(userID, profileID, direction)
}
func (u *fakeSwipeService) GetSwipeCount(userID int, direction string) (int, error) {
	return getSwipeCountService(userID, direction)
}
func (u *fakeSwipeService) GetSwipedProfiles(ctx context.Context, userID int, direction string) ([]string, error) {
	return getSwipedProfilesService(ctx, userID, direction)
}

var swipeServiceFake SwipeServiceInterface = &fakeSwipeService{} //this is where the real implementation is swap with our fake implementation

func TestRecordSwipe_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	recordSwipeService = func(userID, profileID int, direction string) error {
		return nil
	}

	err := swipeServiceFake.RecordSwipe(1, 12, "direction")
	assert.Nil(t, err)
}

func TestGetSwipeCount_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	getSwipeCountService = func(userID int, direction string) (int, error) {
		return 1, nil
	}

	u, err := swipeServiceFake.GetSwipeCount(1, "left")
	assert.Nil(t, err)
	assert.EqualValues(t, u, 1)
}

func TestGetSwipedProfile_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	getSwipedProfilesService = func(ctx context.Context, userID int, direction string) ([]string, error) {
		return []string{"123", "321"}, nil
	}

	userSwiped, err := swipeServiceFake.GetSwipedProfiles(context.Background(), 1, "left")
	assert.Nil(t, err)
	assert.EqualValues(t, len(userSwiped), 2)
}
