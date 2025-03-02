package user

import "time"

type User struct {
    ID        int            `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Email     string         `json:"email" db:"email"`
	Role      string         `json:"role" db:"role"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
}