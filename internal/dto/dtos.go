package dto

import (
	"time"
)

type Duration struct {
	Days    int `json:"days" binding:"min=0"`    // Días, mínimo 0
	Hours   int `json:"hours" binding:"min=0"`   // Horas, mínimo 0
	Minutes int `json:"minutes" binding:"min=0"` // Minutos, mínimo 0
}

type QRRequest struct {
	Text   string `json:"text" binding:"required"`
	ExpiresIn Duration `json:"expires_in" binding:"required"`
}

type QRResponse struct {
	ID             string     `json:"id"`
	Text           string     `json:"text"`
	Format         string     `json:"format"`
	Size           int        `json:"size"`
	CreatedAt      time.Time  `json:"created_at"`
	ExpiresAt      *time.Time `json:"expires_at,omitempty"`
	LastDownloaded *time.Time `json:"last_downloaded,omitempty"`
	DownloadCount  int        `json:"download_count"`
}

func (d Duration) ToDuration() time.Duration {
    return time.Duration(d.Days)*24*time.Hour +
        time.Duration(d.Hours)*time.Hour +
        time.Duration(d.Minutes)*time.Minute
}