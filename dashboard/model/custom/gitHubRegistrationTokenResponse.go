package custom

import (
	"time"
)

type GitHubRegistrationTokenResponse struct {
	Token      string     `json:"token"`
	ExpriresAt *time.Time `json:"expires_at"`
}
