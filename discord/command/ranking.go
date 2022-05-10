package command

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/das08/utils/pkg/settings"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"time"
)

var Ranking = discordgo.ApplicationCommand{
	Name:        "ranking",
	Description: "Show stats ranking of this guild",
}

func RankingResponse(buf *bytes.Buffer, sett *settings.GuildSettings) *discordgo.InteractionResponse {
	embed := discordgo.MessageEmbed{
		URL:  "",
		Type: "",
		Title: sett.LocalizeMessage(&i18n.Message{
			ID:    "commands.ranking.title",
			Other: "Bot Info",
		}),
		Description: "",
		Timestamp:   time.Now().Format(ISO8601),
		Color:       2067276, // DARK GREEN
		Image:       nil,
		Thumbnail:   nil,
		Video:       nil,
		Provider:    nil,
		Author:      nil,
	}
	fields := make([]*discordgo.MessageEmbedField, 1)
	fmt.Println("====", buf.String())
	fields[0] = &discordgo.MessageEmbedField{
		Name: sett.LocalizeMessage(&i18n.Message{
			ID:    "commands.ranking.win",
			Other: "Win Rate Ranking",
		}),
		Value:  buf.String(),
		Inline: true,
	}

	embed.Fields = fields
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&embed},
		},
	}
}
