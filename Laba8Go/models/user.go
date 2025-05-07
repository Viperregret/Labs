package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"gte=0,lte=130"`
}
