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
