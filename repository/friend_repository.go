package repository

import (
	"database/sql"
	"strconv"

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
		row := map[string]string{}
		rows.Scan(
			row["id"],
			row["name"],
			row["birthday"],
			row["created_at"],
			row["updated_at"],
		)

		id, _ := strconv.ParseInt(row["id"], 10, 64)
		createdAt, _ := strconv.ParseFloat(row["created_at"], 64)
		updatedAt, _ := strconv.ParseFloat(row["updated_at"], 64)

		friends = append(friends, &entity.Friend{
			ID:        id,
			Name:      row["name"],
			Birthday:  row["birthday"],
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
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
		friend.ID,
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
