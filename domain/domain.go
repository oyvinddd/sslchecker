package domain

import (
	"net/url"
	"time"
)

type Domain struct {
	URL url.URL `json:"url"`

	IssuedAt time.Time `json:"issued_at"`

	ExpiresAt time.Time `json:"expires_at"`

	Webhook *string `json:"webhook,omitzero"`
}

func New(domainURL string) (*Domain, error) {
	u, err := url.Parse(domainURL)
	if err != nil {
		return nil, err
	}
	return &Domain{URL: u}
}
