package auth

type PasswordService interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}
