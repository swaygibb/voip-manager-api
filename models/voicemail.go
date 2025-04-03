package models

type Voicemail struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Caller    string `json:"caller"`
	Receiver  string `json:"receiver"`
	Duration  int    `json:"duration"`
	Timestamp string `json:"timestamp"`
	Returned bool   `json:"returned"`
	Transcription string `json:"transcription"`
}
