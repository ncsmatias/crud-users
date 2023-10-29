package response

type ProfessorResponse struct {
	ID         string `json:"id"`
	Department string `json:"department"`
	UserID     string `json:"user_id"`
}
