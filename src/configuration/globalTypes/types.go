package globaltypes

type UserRole string

const (
	RoleManager UserRole = "manager"
	RoleTeacher UserRole = "teacher"
	RoleStudent UserRole = "student"
)
