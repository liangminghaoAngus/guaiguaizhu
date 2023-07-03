package data

type Npc struct {
	ID       int    `json:"id"`
	Type     int    `json:"type"`
	Name     string `json:"name"`
	Intro    string `json:"intro"`
	Position string `json:"position"`
	Image    string `json:"image"`
}
