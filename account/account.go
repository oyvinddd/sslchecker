package account

import "github.com/google/uuid"

type Account struct {

	// ID the account ID
	ID uuid.UUID `json:"id"`

	// Email the account email
	Email string `json:"email"`
}
