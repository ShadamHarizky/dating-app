package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ShadamHarizky/dating-app/domain/repository"
	redisRepo "github.com/ShadamHarizky/dating-app/infrastructure/redis"
	"github.com/go-redis/redis/v8"
)

type SwipeService struct {
	rdb redisRepo.RedisDb
	us  repository.UserRepository
}

type SwipeServiceInterface interface {
	RecordSwipe(userID, profileID int, direction string) error
	GetSwipeCount(userID int, direction string) (int, error)
	GetSwipedProfiles(ctx context.Context, userID int, direction string) ([]string, error)
}

func NewSwipeService(rdb redisRepo.RedisDb, us repository.UserRepository) (*SwipeService, error) {
	return &SwipeService{
		rdb: rdb,
		us:  us,
	}, nil
}

func (u *SwipeService) RecordSwipe(userID, profileID int, direction string) error {
	key := direction + ":" + strconv.Itoa(userID)
	score := float64(time.Now().Unix())

	expiration := 24 * time.Hour

	user, err := u.us.GetUser(uint64(userID))
	if err != nil {
		return err
	}

	if user.Premium {
		err := u.rdb.Client.ZAdd(context.Background(), key, &redis.Z{
			Score:  score,
			Member: strconv.Itoa(profileID),
		}).Err()

		if err != nil {
			return err
		}

		return u.rdb.Client.Expire(context.Background(), key, expiration).Err()

	} else {
		countRightSwipe, err := u.GetSwipeCount(int(userID), "right")
		if err != nil {
			return err
		}

		countLeftSwipe, err := u.GetSwipeCount(int(userID), "left")
		if err != nil {
			return err
		}

		total := countRightSwipe + countLeftSwipe
		if total < 10 {
			err := u.rdb.Client.ZAdd(context.Background(), key, &redis.Z{
				Score:  score,
				Member: strconv.Itoa(profileID),
			}).Err()

			if err != nil {
				return err
			}

			return u.rdb.Client.Expire(context.Background(), key, expiration).Err()
		}

		return errors.New("your account already reached limit")

	}
}

func (u *SwipeService) GetSwipeCount(userID int, direction string) (int, error) {
	key := direction + ":" + strconv.Itoa(userID)

	fmt.Println("key", key)

	count, err := u.rdb.Client.ZCount(context.Background(), key, "-inf", "+inf").Result()
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (u *SwipeService) GetSwipedProfiles(ctx context.Context, userID int, direction string) ([]string, error) {
	key := direction + ":" + strconv.Itoa(userID)

	// Retrieve profile IDs from the sorted set
	profiles, err := u.rdb.Client.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve profile IDs: %v", err)
	}

	return profiles, nil
}
