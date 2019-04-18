package data

type Student struct {
	UserID   string `json:"user_id" db:"user_id"`
	Age      int    `json:"age" db:"age"`
	UserName string `json:"user_name" db:"user_name"`
}
