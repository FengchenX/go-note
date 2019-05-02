package entity

type Movie struct {
	ID       string `json:"id"`
	VideoID  string `json:"video_id"`
	Name     string `json:"name"`
	Creator  string `json:"creator"`
	CreateAt string `json:"create_at"`
	Thumb    string `json:"thumb"`
	Describe string `json:"describe"`
}

type FreeMovie struct {
	ID      string `json:"id"`
	MovieID string `json:"movie_id"`
}

type PayMovie struct {
	ID      string  `json:"id"`
	MovieID string  `json:"movie_id"`
	Money   float64 `json:"money"`
	Vip     int     `json:"vip"`
}
