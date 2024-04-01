package interaction

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"github.com/opentracing/opentracing-go"
	"math/rand"
	"strings"
)

func AllCommands(fds *session.FDSConnection) []event.Event {
	return []event.Event{
		&ping{fds},
		&admin{fds},
		&help{fds},
		&play{fds},
		&profile{fds},
		&teams{fds},
		&vcteams{fds},
	}
}

type (
	cmd struct {
		fds *session.FDSConnection
	}

	admin   cmd
	help    cmd
	ping    cmd
	play    cmd
	profile cmd
	teams   cmd
	vcteams cmd
	revoke  cmd
)

func (t *vcteams) Content() any {
	return &discordgo.ApplicationCommand{
		Name:        "vcteams",
		Description: "Generate random teams from the members in your voice channel",
		Version:     "v1.0.1",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "teams",
				Description: "Number of teams (default: 2)",
				Type:        discordgo.ApplicationCommandOptionInteger,
				MinValue:    &two,
				Required:    false,
			},
			{
				Name:        "members",
				Description: "Number of members per team (takes priority over teams)",
				Type:        discordgo.ApplicationCommandOptionInteger,
				MinValue:    &two,
				Required:    false,
			},
		},
	}
}

func (t *vcteams) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ opentracing.Span) error {
	options := i.ApplicationCommandData().Options
	teamsInput := utils.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "teams"
	})
	membersInput := utils.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "members"
	})

	var teamAmount int64
	memberAmount := int64(2)
	if teamsInput != nil {
		teamAmount = teamsInput.IntValue()
	}
	if membersInput != nil {
		memberAmount = membersInput.IntValue()
	}

	userID := i.Member.User.ID
	guildID := i.GuildID
	voiceState, err := s.State.VoiceState(guildID, userID)
	if err != nil {
		return fmt.Errorf("you need to be in a voice channel to use this command")
	}
	voiceChannelID := voiceState.ChannelID
	guild, err := s.State.Guild(guildID)
	if err != nil {
		return err
	}

	var members []string
	for _, vs := range guild.VoiceStates {
		if vs.ChannelID == voiceChannelID {
			member, err := s.GuildMember(guildID, vs.UserID)
			if err != nil {
				return err
			}
			members = append(members, member.Nick)
		}
	}

	var teams [][]string
	if memberAmount != 0 {
		teams = make([][]string, len(members)/int(memberAmount))
	} else {
		teams = make([][]string, int(teamAmount))
	}

	rand.Shuffle(len(members), func(i, j int) {
		members[i], members[j] = members[j], members[i]
	})
	for i, player := range members {
		teams[i%len(teams)] = append(teams[i%len(teams)], player)
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				(&embedTeams{teams}).Content().(*discordgo.MessageEmbed),
			},
		},
	})
}

var two = 2.0

func (t *teams) Content() any {
	return &discordgo.ApplicationCommand{
		Name:        "teams",
		Description: "Generate random teams from input",
		Version:     "v1.0.3",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "players",
				Description: "List of players seperated by a space",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "teams",
				Description: "Number of teams (default: 2)",
				Type:        discordgo.ApplicationCommandOptionInteger,
				MinValue:    &two,
				Required:    false,
			},
			{
				Name:        "members",
				Description: "Number of members per team (takes priority over teams amount)",
				Type:        discordgo.ApplicationCommandOptionInteger,
				MinValue:    &two,
				Required:    false,
			},
		},
	}
}

func (t *teams) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ opentracing.Span) error {
	options := i.ApplicationCommandData().Options
	playersStr := utils.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "players"
	}).StringValue()
	teamsInput := utils.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "teams"
	})
	membersInput := utils.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "members"
	})
	players := strings.Split(playersStr, " ")
	playerAmount := len(players)

	var teamAmount int64
	memberAmount := int64(2)
	if teamsInput != nil {
		teamAmount = teamsInput.IntValue()
	}
	if membersInput != nil {
		memberAmount = membersInput.IntValue()
	}

	var teams [][]string
	if memberAmount != 0 {
		teams = make([][]string, playerAmount/int(memberAmount))
	} else {
		teams = make([][]string, int(teamAmount))
	}

	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})
	for i, player := range players {
		teams[i%len(teams)] = append(teams[i%len(teams)], player)
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				(&embedTeams{teams}).Content().(*discordgo.MessageEmbed),
			},
		},
	},
	)
}

func (p *profile) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, sp opentracing.Span) error {
	option := i.ApplicationCommandData().Options[0]
	var id string
	if option == nil {
		id = i.Member.User.ID
	} else {
		id = option.UserValue(s).ID
	}

	member, err := p.fds.Member(sp, id)
	if err != nil {
		return err
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				(&embedProfile{member}).Content().(*discordgo.MessageEmbed),
			},
		},
	})
}

func (p *profile) Content() any {
	return &discordgo.ApplicationCommand{
		Name:        "profile",
		Description: "View the profile of a member",
		Version:     "v1.0.0",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "member",
				Description: "The member to view the profile of",
				Required:    false,
			},
		},
	}
}

func (p *play) Content() any {
	return &discordgo.ApplicationCommand{
		Name:        "play",
		Description: "Ask the server to play any gamemode with you",
		Version:     "v1.1.0",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "gamemode",
				Description: "The gamemode you want to play",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Value: ChoiceGamemodeScrims,
						Name:  "Scrims",
					},
					{
						Value: ChoiceGamemodeBridgeQueue,
						Name:  "Bridge Queue",
					},
					{
						Value: ChoiceGamemodeBedwars,
						Name:  "BedWars",
					},
					{
						Value: ChoiceGamemodeOther,
						Name:  "Other",
					},
				},
			},
			{
				Name:        "description",
				Description: "A description of what you want to play (e.g. submode, players, map etc.)",
				Type:        discordgo.ApplicationCommandOptionString,
				MaxLength:   512,
				Required:    false,
			},
		},
	}
}

const (
	ChannelLTP                = "❓│looking-to-play" // TODO: replace with env
	ChoiceGamemodeScrims      = "scrims"
	ChoiceGamemodeBridgeQueue = "bridge_queue"
	ChoiceGamemodeBedwars     = "bedwars"
	ChoiceGamemodeOther       = "other"
)

func (p *play) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ opentracing.Span) error {
	options := i.ApplicationCommandData().Options
	gamemode := utils.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "gamemode"
	})
	description := utils.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "description"
	})

	channels, err := s.GuildChannels(i.GuildID)
	if err != nil {
		return err
	}
	ltpChannel := utils.Find(channels, func(c *discordgo.Channel) bool {
		return c.Name == ChannelLTP // TODO: per guild channel name
	})
	if ltpChannel == nil {
		return fmt.Errorf("LTP channel not found")
	}
	if i.ChannelID != ltpChannel.ID {
		return fmt.Errorf("command can only be used in %v", ltpChannel.Mention())
	}

	var roleName string // TODO: per guild role names
	switch gamemode.Value {
	case ChoiceGamemodeScrims:
		roleName = "Notify Bridge Scrims"
	case ChoiceGamemodeBridgeQueue:
		roleName = "Notify Bridge Queue"
	case ChoiceGamemodeBedwars:
		roleName = "Notify BedWars"
	case ChoiceGamemodeOther:
		roleName = "Notify Random"
	}
	roles, err := s.GuildRoles(i.GuildID)
	if err != nil {
		return err
	}
	gamemodeRole := utils.Find(roles, func(r *discordgo.Role) bool {
		return r.Name == roleName
	})
	if gamemodeRole == nil {
		return fmt.Errorf("gamemode role not found")
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("By: %s to: %s", i.Member.Mention(), gamemodeRole.Mention()),
			Embeds: []*discordgo.MessageEmbed{
				(&embedPlay{
					member: i.Member,
					to:     gamemodeRole,
					desc:   description,
					mode: func() (name string) {
						switch gamemode.Value {
						case ChoiceGamemodeScrims:
							name = "Scrims"
						case ChoiceGamemodeBridgeQueue:
							name = "Bridge Queue"
						case ChoiceGamemodeBedwars:
							name = "BedWars"
						case ChoiceGamemodeOther:
							name = "Something"
						}
						return
					}(),
				}).Content().(*discordgo.MessageEmbed),
			},
		},
	})
}

func (a *admin) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ opentracing.Span) error {
	options := i.ApplicationCommandData().Options
	var message *discordgo.MessageSend
	for _, option := range options {
		switch option.Name {
		case "embed":
			switch option.StringValue() {
			case "verify":
				message = &discordgo.MessageSend{
					Embeds: []*discordgo.MessageEmbed{
						new(embedVerify).Content().(*discordgo.MessageEmbed),
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								(&buttonVerify{a.fds}).Content().(*discordgo.Button),
							},
						},
					},
				}
			}
		}
	}

	_, err := s.ChannelMessageSendComplex(i.ChannelID, message)
	if err != nil {
		return err
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message was sent to channel",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

var adminPerms = int64(discordgo.PermissionAdministrator)

func (a *admin) Content() any {
	return &discordgo.ApplicationCommand{
		Name:                     "admin",
		Description:              "Admin utilities",
		Version:                  "v1.0.3",
		DefaultMemberPermissions: &adminPerms,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "embed",
				Description: "Write an embed to a channel",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "verify",
						Value: "verify",
					},
				},
			},
		},
	}
}

func (h *help) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ opentracing.Span) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				(&embedHelp{s, func() map[string]*discordgo.ApplicationCommand {
					m := make(map[string]*discordgo.ApplicationCommand)
					for _, cmd := range AllCommands(h.fds) {
						content := cmd.Content().(*discordgo.ApplicationCommand)
						m[content.Name] = content
					}
					return m
				}()}).Content().(*discordgo.MessageEmbed),
			},
		},
	})
}

func (h *help) Content() any {
	return &discordgo.ApplicationCommand{
		Name:        "help",
		Description: "Get help",
		Version:     "v1.0.0",
	}
}

func (p *ping) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ opentracing.Span) error {
	latency := s.HeartbeatLatency().Milliseconds()
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("The bots latency is %vms", latency),
		},
	})
}

func (p *ping) Content() any {
	return &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping the bot",
		Version:     "v1.2.0",
	}
}
