package domain

type Coordinate struct {
	Id          int      `json:"id"`
	Image       string   `json:"image"`
	UserId      string   `json:"userId"`
	CreatedAt   JsonTime `json:"createdAt"`
	Favorites   int      `json:"favorites"`
	IsFavorited bool     `json:"isFavorited"`
}
