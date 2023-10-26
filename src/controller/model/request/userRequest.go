package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,containsany=!@$#*"`
	Name     string `json:"name" binding:"max=128"`
	Admin    bool   `json:"admin" binding:"boolean"`
}
