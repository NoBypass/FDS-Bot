package components

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/model"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
)

func ImageProfile(member *model.MemberResponse) *discordgo.File {
	return &discordgo.File{
		Name:        "profile.png",
		ContentType: "image/png",
		Reader: func() io.Reader {
			const W = 200
			const H = 14
			face := FontInter()
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
			labelWidth := font.MeasureString(face, label).Round()
			_, labelHeight, _ := face.GlyphBounds('M')
			point := fixed.Point26_6{
				X: (fixed.I(W) - fixed.I(labelWidth)) / 2,
				Y: ((fixed.I(H) + labelHeight) / 2) - fixed.I(1),
			}
			d := &font.Drawer{
				Dst:  img,
				Src:  image.White,
				Face: face,
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
	}
}
