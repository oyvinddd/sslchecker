package domain

import (
	"net/url"
	"time"
)

type Domain struct {
	URL url.URL `json:"url"`

	IssuedAt time.Time `json:"issued_at"`

	ExpiresAt time.Time `json:"expires_at"`
}
