package entity

type Movie struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Creator     string `json:"creator"`
	CreateAt    string `json:"create_at"`
	Thumb       string `json:"thumb"`
	Describe    string `json:"describe"`
	MainPlayers string `json:"main_players"`
	Types       string `json:"types"`
	Year        int    `json:"year"`
}

type MovieVideo struct {
	ID       string `json:"id"`
	Meta     string `json:"meta"`
	MovieID  string `json:"movie_id"`
	Creator  string `json:"creator"`
	CreateAt string `json:"create_at"`
}
