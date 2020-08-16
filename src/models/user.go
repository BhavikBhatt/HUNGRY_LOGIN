package models

type User struct {
    Username string `json:"Field Str"`
	Ciphertext string `json:"Field Str"`
	Email string `json:"Field Str"`
	Name string `json:"Field Str"`
	Age int `json:"Field Int"`
}