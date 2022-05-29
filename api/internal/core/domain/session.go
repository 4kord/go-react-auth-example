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

func (s *Session) ValidateExpiry() bool {
    return !time.Now().UTC().After(s.ExpiresAt)
}

func (s *Session) ValidateIp(requestIp string) bool {
    return s.Ip == requestIp

} 
