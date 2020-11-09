package user

//User is the structure for users table
type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
