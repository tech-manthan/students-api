package types

type Student struct {
	Id    int64  `json:"id"`
	Name  string `validate:"required" json:"name"`
	Email string `validate:"required,email" json:"email"`
	Age   int    `validate:"required" json:"age"`
}
