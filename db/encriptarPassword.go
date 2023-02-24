package db

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutina que me permite encriptar la password antes de ser grabada en la BD */
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
