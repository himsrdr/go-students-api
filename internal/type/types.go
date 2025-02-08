package types

type Student struct {
	Id    int64
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}

type Studentupdate struct {
	Email string `json:"email" validate:"required"`
}
