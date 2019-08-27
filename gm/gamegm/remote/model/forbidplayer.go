package model

type ForbidPlayer struct {
	PlayerId     int64  `json:"playerId"`
	ForbidReason string `json:"forbidReason"`
	ForbidName   string `json:"forbidName"`
	ForbidTime   int64  `json:"forbidTime"`
}
