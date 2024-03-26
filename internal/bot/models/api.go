package models

type VerifyRequest struct {
	ID   string `json:"id"`
	Nick string `json:"nick"`
	Name string `json:"name"`
}

type VerifyResponse struct {
	Actual string `json:"actual"`
}

type MemberResponse struct {
	DiscordID   string  `json:"discord_id"`
	Name        string  `json:"name"`
	Nick        string  `json:"nick"`
	XP          float64 `json:"xp"`
	LastDailyAt string  `json:"last_daily_at"`
	Level       int     `json:"level"`
	Streak      int     `json:"streak"`
}

type LoginRequest struct {
	Pwd string `json:"pwd"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LeaderboardResponse []struct {
	DiscordID string  `json:"discord_id"`
	Level     int     `json:"level"`
	XP        float64 `json:"xp"`
}
