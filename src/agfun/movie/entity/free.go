package entity

// free movie
// swagger:model
type FreeMovie struct {
	ID          string `json:"id"`
	MovieID     string `json:"movie_id"`
	FreeVideoID string `json:"free_video_id"`
}
