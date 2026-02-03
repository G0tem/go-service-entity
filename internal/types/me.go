package types

import "time"

type SuccessResponseMe struct {
	UserID      string    `json:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	Permissions []string  `json:"permissions"`
	Exp         time.Time `json:"exp"`
}
