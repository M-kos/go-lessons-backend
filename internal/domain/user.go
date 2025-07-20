package domain

import "time"

type User struct {
	ID         int64
	Email      string
	FullName   string
	CreateTime time.Time
}
