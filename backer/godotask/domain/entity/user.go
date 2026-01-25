package entity

type User struct {
	ID       uint   
	Email    string
	Username string
	Role     string 
	PasswordHash string
}
