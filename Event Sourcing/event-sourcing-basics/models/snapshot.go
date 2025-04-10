package models

import "time"

type Snapshot struct {
	ID			int64		`json:"id" gorm:"primaryKey"`
	AccountID 	string 		`json:"account_id"`
	Balance		int64		`json:"balance"`
	Version		int			`json:"version"`
	Timestamp	time.Time	`json:"timestamp" gorm:"autoCreateTime"`
}
