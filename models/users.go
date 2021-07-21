package modles

import "time"

type Users struct {
	// id
	Id int64 `json:"id"`
	// name
	Name string `json:"name"`
	// email
	Email string `json:"email"`
	// email_verified_at
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	// password
	Password string `json:"password"`
	// remember_token
	RememberToken string `json:"remember_token"`
	// created_at
	CreatedAt time.Time `json:"created_at"`
	// updated_at
	UpdatedAt time.Time `json:"updated_at"`
}

func (model Users) TableName() string {
	return "users"
}
