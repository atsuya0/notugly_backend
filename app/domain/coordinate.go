package domain

type Coordinate struct {
	Id          int      `json:"id"`
	ImageName   string   `json:"imageName"`
	UserId      string   `json:"userId"`
	CreatedAt   JsonTime `json:"createdAt"`
	Favorites   int      `json:"favorites"`
	IsFavorited bool     `json:"isFavorited"`
}
