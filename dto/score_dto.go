package dto

type ScoreRequest struct {
	KrsID int `json:"krs_id" binding:"required" validate:"numeric"`
	Score int `json:"score"`
}

type ScoreResponse struct {
	ID    int `json:"id"`
	KrsID int `json:"krs_id"`
	Score int `json:"score"`
}
