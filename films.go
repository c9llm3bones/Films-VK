package Films

type Film struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FilmsList struct {
	Id     int
	FilmId int
	ListId int
}
