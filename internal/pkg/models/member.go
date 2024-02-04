package models

type DailyResponse struct {
	XP        float64 `json:"xp"`
	Level     int     `json:"level"`
	Levelup   bool    `json:"levelup"`
	Needed    float64 `json:"needed"`
	Gained    float64 `json:"gained"`
	Streak    int     `json:"streak"`
	WithBonus float64 `json:"with_bonus"`
}

type DiscordVerifyInput struct {
	ID   string `json:"id"`
	Nick string `json:"nick"`
	Name string `json:"name"`
}

type DiscordVerifyResponse struct {
	Actual string `json:"actual"`
}
