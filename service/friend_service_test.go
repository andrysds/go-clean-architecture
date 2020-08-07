package service

import (
	"errors"
	"testing"

	"github.com/andrysds/go-clean-architecture/entity"
	"github.com/andrysds/go-clean-architecture/test/entitymock"
	"github.com/andrysds/go-clean-architecture/test/repositorymock"
	"github.com/stretchr/testify/assert"
)

var errFromFriendRepository = errors.New("error from FriendRepository")

func TestFriendServiceGetFriends(t *testing.T) {
	cases := []struct {
		friendRepositoryErr error
		friends             []*entity.Friend
		err                 error
	}{
		// context: got error from FriendRepository
		{

			friendRepositoryErr: errFromFriendRepository,
			friends:             []*entity.Friend{},
			err:                 errFromFriendRepository,
		},
		// context: positive case
		{

			friendRepositoryErr: nil,
			friends:             entitymock.Friends,
			err:                 nil,
		},
	}

	for _, c := range cases {
		service := &FriendService{
			FriendRepository: &repositorymock.FriendRepositoryMock{
				Err: c.friendRepositoryErr,
			},
		}
		friends, err := service.GetFriends()
		assert.Equal(t, c.friends, friends)
		assert.Equal(t, c.err, err)
	}
}

func TestFriendServiceCreateFriend(t *testing.T) {
	cases := []struct {
		friendRepositoryErr error
		friend              *entity.Friend
		err                 error
	}{
		// context: got error from FriendRepository
		{
			friendRepositoryErr: errFromFriendRepository,
			friend:              nil,
			err:                 errFromFriendRepository,
		},
		// context: positive case
		{

			friendRepositoryErr: nil,
			friend:              entitymock.FriendAndrys,
			err:                 nil,
		},
	}

	for _, c := range cases {
		service := &FriendService{
			FriendRepository: &repositorymock.FriendRepositoryMock{
				Err: c.friendRepositoryErr,
			},
		}
		input := entitymock.FriendAndrys
		input.ID = 0
		friend, err := service.CreateFriend(input)
		assert.Equal(t, c.friend, friend)
		assert.Equal(t, c.err, err)
	}
}

func TestFriendServiceUpdateFriend(t *testing.T) {
	cases := []struct {
		friendRepositoryErr error
		friend              *entity.Friend
		err                 error
	}{
		// context: got error from FriendRepository
		{

			friendRepositoryErr: errFromFriendRepository,
			friend:              nil,
			err:                 errFromFriendRepository,
		},
		// context: positive case
		{

			friendRepositoryErr: nil,
			friend:              entitymock.FriendBudi,
			err:                 nil,
		},
	}

	for _, c := range cases {
		service := &FriendService{
			FriendRepository: &repositorymock.FriendRepositoryMock{
				Err: c.friendRepositoryErr,
			},
		}
		friend, err := service.UpdateFriend(entitymock.FriendAndrys)
		assert.Equal(t, c.friend, friend)
		assert.Equal(t, c.err, err)
	}
}

func TestFriendServiceDeleteFriend(t *testing.T) {
	cases := []struct {
		friendRepositoryErr error
		err                 error
	}{
		// context: got error from FriendRepository
		{

			friendRepositoryErr: errFromFriendRepository,
			err:                 errFromFriendRepository,
		},
		// context: positive case
		{

			friendRepositoryErr: nil,
			err:                 nil,
		},
	}

	for _, c := range cases {
		service := &FriendService{
			FriendRepository: &repositorymock.FriendRepositoryMock{
				Err: c.friendRepositoryErr,
			},
		}
		err := service.DeleteFriend(1)
		assert.Equal(t, c.err, err)
	}
}
