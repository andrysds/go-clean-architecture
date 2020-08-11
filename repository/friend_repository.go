package repository

import (
	"database/sql"

	"github.com/andrysds/go-clean-architecture/entity"
)

// FriendRepository is repository for providing and modifying friends data
type FriendRepository struct {
	DB *sql.DB
}

// FindAll retrieves all friends data
func (r *FriendRepository) FindAll() ([]*entity.Friend, error) {
	query := "SELECT id, name, birthday, created_at, updated_at FROM friends"
	rows, err := r.DB.Query(query)
	if err != nil {
		return []*entity.Friend{}, err
	}

	friends := []*entity.Friend{}
	for rows.Next() {
		row := &entity.Friend{}
		rows.Scan(
			&row.ID,
			&row.Name,
			&row.Birthday,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		friends = append(friends, row)
	}

	return friends, nil

}

// Create inserts a row of friend data to friends table
func (r *FriendRepository) Create(friend *entity.Friend) (*entity.Friend, error) {
	query := "INSERT INTO friends (name, birthday, created_at, updated_at) VALUES (?, ?, ?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(
		friend.Name,
		friend.Birthday,
		friend.CreatedAt,
		friend.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	friend.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return friend, nil
}

// Update updates a row of friend data from friends table
func (r *FriendRepository) Update(friend *entity.Friend) (*entity.Friend, error) {
	query := "UPDATE friends SET name=?, birthday=?, created_at=?, updated_at=? WHERE id=?"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		friend.Name,
		friend.Birthday,
		friend.CreatedAt,
		friend.UpdatedAt,
		friend.ID,
	)
	if err != nil {
		return nil, err
	}

	return friend, nil
}

// Delete deletes a row of friend data from friends table
func (r *FriendRepository) Delete(id int64) error {
	query := "DELETE FROM friends WHERE id=?"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
