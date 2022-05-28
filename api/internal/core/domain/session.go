package domain

import "time"

type Session struct {
    Id int
    UserId int
    RefreshToken string
    Ip string
    ExpiresAt time.Time
    CreatedAt time.Time
}

func (s *Session) Validate(ip string) bool {
    if time.Now().UTC().After(s.ExpiresAt) {
        return false
    }

    if s.Ip != ip {
        return false
    }

    return true
}
