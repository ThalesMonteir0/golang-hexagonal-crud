package domain

import "golang.org/x/crypto/bcrypt"

type UserDomain struct {
	ID       int
	Name     string
	Email    string
	Password string
	Active   bool
}

func (ud *UserDomain) EncryptPassword() error {
	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	ud.Password = string(passwordEncrypted)
	return nil
}
