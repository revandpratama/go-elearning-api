package dto

type KRSRequest struct {
	UserID    int    `json:"user_id" binding:"required"`
	SubjectID int    `json:"subject_id" binding:"required"`
	Status    string `json:"status" binding:"required"`
}

type KRSResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	SubjectID int    `json:"subject_id"`
	Status    string `json:"status"`
}
