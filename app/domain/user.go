package domain

type User struct {
	Id     string `json:"uid"`
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Age    int    `json:"age"`
}
