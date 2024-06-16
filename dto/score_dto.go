package dto

type ScoreRequest struct {
	UserID int `json:"user_id" binding:"required" validate:"numeric"`
	KrsID  int `json:"krs_id" binding:"required" validate:"numeric"`
	Score  int `json:"score"`
}

type ScoreResponse struct {
	ID         int      `json:"id"`
	UserID     int      `json:"user_id"`
	KrsID      int      `json:"krs_id"`
	Score      int      `json:"score"`
	// Pagination Paginate `json:"pagination"`
}
