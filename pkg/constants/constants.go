package constants

import "os"

var (
	DbHost     = os.Getenv("DB_ENDPOINT")
	DbPort     = os.Getenv("DB_PORT")
	DbDatabase = os.Getenv("DB_NAME")
	DbUser     = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
)

var (
	ErrorInvalidID      = "invalid ID"
	ErrorInvalidData    = "invalid data"
	ErrorInvalidEmail   = "invalid email"
	ErrorCouldNotDelete = "could not delete"
)
