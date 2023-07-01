package helpers

import "github.com/asaskevich/govalidator"

// IsValidEmail adalah fungsi untuk memvalidasi alamat email
func IsValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}
