package entity

// Friend is a struct for friend entity
type Friend struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Birthday  string  `json:"birthday"`
	CreatedAt float64 `json:"created_at"`
	UpdatedAt float64 `json:"updated_at"`
}
