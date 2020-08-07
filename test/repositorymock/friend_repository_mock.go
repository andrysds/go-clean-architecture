package repositorymock

import (
	"github.com/andrysds/go-clean-architecture/entity"
	"github.com/andrysds/go-clean-architecture/test/entitymock"
)

// FriendRepositoryMock is repository.FriendRepository mock
type FriendRepositoryMock struct {
	Err error
}

// FindAll is service.FriendRepository.FindAll mock
func (m *FriendRepositoryMock) FindAll() ([]*entity.Friend, error) {
	if m.Err != nil {
		return []*entity.Friend{}, m.Err
	}
	return entitymock.Friends, nil
}

// Create is service.FriendRepository.Create mock
func (m *FriendRepositoryMock) Create(*entity.Friend) (*entity.Friend, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return entitymock.FriendAndrys, nil
}

// Update is service.FriendRepository.Update mock
func (m *FriendRepositoryMock) Update(*entity.Friend) (*entity.Friend, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return entitymock.FriendBudi, nil
}

// Delete is service.FriendRepository.Delete mock
func (m *FriendRepositoryMock) Delete(int64) error {
	return m.Err
}
