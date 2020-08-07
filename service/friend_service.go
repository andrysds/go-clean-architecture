package service

import (
	"github.com/andrysds/go-clean-architecture/entity"
)

// FriendUseCase defines all friend CRUD use cases
type FriendUseCase interface {
	GetFriends() ([]*entity.Friend, error)
	CreateFriend(*entity.Friend) (*entity.Friend, error)
	UpdateFriend(*entity.Friend) (*entity.Friend, error)
	DeleteFriend(int64) error
}

// FriendService implements FriendUseCase
type FriendService struct {
	FriendRepository interface {
		FindAll() ([]*entity.Friend, error)
		Create(*entity.Friend) (*entity.Friend, error)
		Update(*entity.Friend) (*entity.Friend, error)
		Delete(int64) error
	}
}

// GetFriends implements FriendUseCase.GetFriends
func (s *FriendService) GetFriends() ([]*entity.Friend, error) {
	friends, err := s.FriendRepository.FindAll()

	// write other business logic/rules here

	return friends, err
}

// CreateFriend implements FriendUseCase.CreateFriend
func (s *FriendService) CreateFriend(friend *entity.Friend) (*entity.Friend, error) {
	created, err := s.FriendRepository.Create(friend)

	// write other business logic/rules here

	return created, err
}

// UpdateFriend implements FriendUseCase.UpdateFriend
func (s *FriendService) UpdateFriend(friend *entity.Friend) (*entity.Friend, error) {
	updated, err := s.FriendRepository.Update(friend)

	// write other business logic/rules here

	return updated, err
}

// DeleteFriend implements FriendUseCase.DeleteFriend
func (s *FriendService) DeleteFriend(id int64) error {
	err := s.FriendRepository.Delete(id)

	// write other business logic/rules here

	return err
}
