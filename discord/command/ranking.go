package command

import (
	"github.com/bwmarrin/discordgo"
)

var Ranking = discordgo.ApplicationCommand{
	Name:        "ranking",
	Description: "Show stats ranking of this guild",
}
