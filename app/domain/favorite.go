package domain

type Favorite struct {
	Id           int    `json:"id"`
	Flag         bool   `json:"flag"`
	Coordinateid string `json:"coordinateId"`
	UserId       string `json:"userId"`
}
