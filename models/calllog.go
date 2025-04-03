package models

type CallLog struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Caller    string `json:"caller"`
	Receiver  string `json:"receiver"`
	Duration  int    `json:"duration"`
	Timestamp string `json:"timestamp"`

	Voicemails []Voicemail `gorm:"foreignKey:Caller;references:Caller"`
}
