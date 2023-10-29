package request

type ProfessorRequest struct {
	Department string `json:"department" binding:"required"`
	UserID     string `json:"user_id" binding:"required"`
}
