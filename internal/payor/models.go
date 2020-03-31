package payor

import "time"

// User is ...
type User struct {
	ID        string    `json:"id" gorm:"primary_key;"`
	Userame   string    `json:"username" gorm:"not null"`
	Password  string    `json:"password"`
	APIKey    string    `json:"api_key"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}
