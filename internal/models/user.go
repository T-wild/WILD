package models

import "time"

type User struct {
	BasicModel

	Userid      int64     `json:"userid"`        // User ID        string    `json:"name"`          // Name
	Username    string    `json:"username"`      // Username
	Password    string    `json:"password"`      // Password
	LastLoginAt time.Time `json:"last_login_at"` // Last login time
}
