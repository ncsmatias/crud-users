package globaltypes

type UserRole string

const (
	RoleManager   UserRole = "manager"
	RoleProfessor UserRole = "professor"
	RoleStudent   UserRole = "student"
)
