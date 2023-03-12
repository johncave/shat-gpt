package main

import "time"

type WSLeaderboard struct {
	EventName   string              `json:"event_name`
	Leaderboard []WSLeaderboardUser `json:"leaderboard"`
	Timestamp   time.Time           `json:"timestamp"`
}

type WSLeaderboardUser struct {
	UserName string `json:"username"`
	Score    int    `json:"score"`
	Rank     int    `json:"rank"`
}
