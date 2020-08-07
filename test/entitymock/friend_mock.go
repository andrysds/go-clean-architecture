package entitymock

import (
	"github.com/andrysds/go-clean-architecture/entity"
)

// FriendAndrys is a *entity.Friend mock with name:Andrys
var FriendAndrys *entity.Friend = &entity.Friend{
	ID:        1,
	Name:      "Andrys",
	Birthday:  "2020-02-14",
	CreatedAt: 1596781398923,
	UpdateAt:  1596781398923,
}

// FriendBudi is a *entity.Friend mock with name:Budi
var FriendBudi *entity.Friend = &entity.Friend{
	ID:        2,
	Name:      "Budi",
	Birthday:  "2020-03-14",
	CreatedAt: 1596781398923,
	UpdateAt:  1596781398923,
}

// Friends is a mock []*entity.Friend mock
var Friends = []*entity.Friend{
	FriendAndrys,
	FriendBudi,
}
