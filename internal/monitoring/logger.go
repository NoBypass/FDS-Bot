package monitoring

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"time"
)

func LogEvent(i *discordgo.InteractionCreate, sp opentracing.Span, latency time.Duration, errs ...error) {
	var e string
	if len(errs) > 0 && errs[0] != nil {
		e = errs[0].Error()
	}

	data := map[string]any{
		"time":           time.Now().Format(time.RFC3339),
		"type":           i.Type,
		"name":           utils.InteractionName(i),
		"latency":        latency.String(),
		"guild":          i.GuildID,
		"channel":        i.ChannelID,
		"member":         i.Member.User.ID,
		"username":       i.Member.User.Username,
		"interaction_id": i.Interaction.ID,
		"error":          e,
		"trace_id":       sp.Context().(jaeger.SpanContext).TraceID().String(),
	}

	sp.LogKV(
		"type", data["type"],
		"name", data["name"],
		"guild", data["guild"],
		"channel", data["channel"],
		"member", data["member"],
		"username", data["username"],
		"interaction_id", data["interaction_id"],
	)

	j, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	println(string(j))
}
