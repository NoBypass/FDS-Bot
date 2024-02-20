package commands

import (
	"bytes"
	"fmt"
	"github.com/NoBypass/fds/pkg/api"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
)

type Profile struct {
	api  *api.Client
	face font.Face
}

func (p *Profile) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	options := i.ApplicationCommandData().Options
	var id string
	if len(options) == 0 {
		id = i.Member.User.ID
	} else {
		id = options[0].UserValue(s).ID
	}

	member, err := p.api.Member(id)
	if err != nil {
		return err
	}

	dcMember, err := s.GuildMember(i.GuildID, id)
	if err != nil {
		return err
	}

	msg, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:        "profile.png",
				ContentType: "image/png",
				Reader: func() io.Reader {
					const W = 200
					const H = 14
					img := image.NewRGBA(image.Rect(0, 0, W, H))
					needed := member.GetNeededXP()
					total := needed + member.XP
					progress := member.XP / total

					barWidth := int(float64(W) * progress)
					gray := color.RGBA{R: 35, G: 38, B: 42, A: 255}
					draw.Draw(img, image.Rect(0, 0, W, H), &image.Uniform{C: gray}, image.Point{}, draw.Src)
					green := color.RGBA{R: 16, G: 185, B: 129, A: 255}
					draw.Draw(img, image.Rect(0, 0, barWidth, H), &image.Uniform{C: green}, image.Point{}, draw.Src)

					label := fmt.Sprintf("%d/%d", int(member.XP), int(total))
					labelWidth := font.MeasureString(p.face, label).Round()
					_, labelHeight, _ := p.face.GlyphBounds('M')
					point := fixed.Point26_6{
						X: (fixed.I(W) - fixed.I(labelWidth)) / 2,
						Y: ((fixed.I(H) + labelHeight) / 2) - fixed.I(1),
					}
					d := &font.Drawer{
						Dst:  img,
						Src:  image.White,
						Face: p.face,
						Dot:  point,
					}
					d.DrawString(label)

					buf := new(bytes.Buffer)
					err := png.Encode(buf, img)
					if err != nil {
						log.Println("failed to create image:", err)
					}
					return bytes.NewReader(buf.Bytes())
				}(),
			},
		},
	})
	if err != nil {
		return err
	}

	imageURL := msg.Attachments[0].URL
	err = s.ChannelMessageDelete(i.ChannelID, msg.ID)
	if err != nil {
		return err
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedProfile(dcMember, member, imageURL),
			},
		},
	})
}

func (p *Profile) Content() *discordgo.ApplicationCommand {
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
