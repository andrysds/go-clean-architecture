package servicemock

import (
	"github.com/andrysds/go-clean-architecture/entity"
	"github.com/andrysds/go-clean-architecture/test/entitymock"
)

// FriendServiceMock is service.FriendService mock
type FriendServiceMock struct {
	Err error
}

// GetFriends is service.FriendService.GetFriends mock
func (m *FriendServiceMock) GetFriends() ([]*entity.Friend, error) {
	if m.Err != nil {
		return []*entity.Friend{}, m.Err
	}
	return entitymock.Friends, nil
}

// CreateFriend is service.FriendService.CreateFriend mock
func (m *FriendServiceMock) CreateFriend(*entity.Friend) (*entity.Friend, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return entitymock.FriendAndrys, nil
}

// UpdateFriend is service.FriendService.UpdateFriend mock
func (m *FriendServiceMock) UpdateFriend(*entity.Friend) (*entity.Friend, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return entitymock.FriendBudi, nil
}

// DeleteFriend is service.FriendService.DeleteFriend mock
func (m *FriendServiceMock) DeleteFriend(id int64) error {
	return m.Err
}
