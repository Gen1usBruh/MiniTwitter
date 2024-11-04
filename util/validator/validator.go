package validator

import (
	"github.com/Gen1usBruh/MiniTwitter/internal/storage/models"
	"github.com/Gen1usBruh/MiniTwitter/internal/storage/postgres/sqlc"
	"github.com/asaskevich/govalidator"
)

func ValidateUserSignUp(user postgresdb.CreateUserParams) bool {
	if len(user.Username) == 0 || len(user.Email) == 0 || len(user.Password) == 0 {
		return false
	}
	return govalidator.IsEmail(user.Email) && (len(user.Username) < 20)
}

func ValidateUserSignIn(user models.SignInStruct) bool {
	if len(user.Email) == 0 || len(user.Password) == 0 {
		return false
	}
	return govalidator.IsEmail(user.Email)
}