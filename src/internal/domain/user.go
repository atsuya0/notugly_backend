package domain

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Sex  int    `json:"sex"`
	Age  int    `json:"age"`
}
