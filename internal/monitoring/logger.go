package monitoring

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"time"
)

func LogEvent(i *discordgo.InteractionCreate, latency time.Duration) {
	data := map[string]any{
		"time":     time.Now().Format(time.RFC3339),
		"type":     i.Type,
		"name":     utils.InteractionName(i),
		"latency":  latency.String(),
		"guild":    i.GuildID,
		"channel":  i.ChannelID,
		"member":   i.Member.User.ID,
		"username": i.Member.User.Username,
		"version":  i.Version,
	}
	j, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	println(string(j))
}
