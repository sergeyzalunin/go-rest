package models

type User struct {
	ID                int
	Email             string
	EncryptedPassword string `sql`
}
