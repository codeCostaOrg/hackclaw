package commands

import (
	"codecosta.com/hackclaw/app/models"
	"codecosta.com/hackclaw/app/utils"
	"github.com/bwmarrin/discordgo"
)

var CommandList = []*discordgo.ApplicationCommand{
	{
		Name:        "spawns",
		Description: "Show spawns",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "map",
				Description: "Map Name",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  string(models.ZERO_DAM),
						Value: models.ZERO_DAM,
					},
					{
						Name:  string(models.LAYALI_GROVE),
						Value: models.LAYALI_GROVE,
					},
					{
						Name:  string(models.BRAKKESH),
						Value: models.BRAKKESH,
					},
					{
						Name:  string(models.SPACE_CITY),
						Value: models.SPACE_CITY,
					},
				},
				Required: true,
			},
		},
	},
	{
		Name:        "test",
		Description: "Test command",
	},
}

var CommandHandlers = map[string]func(discord *discordgo.Session, interaction *discordgo.InteractionCreate){
	"test":   test,
	"spawns": SendSpawns,
}

func test(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	utils.LogCommand(interaction.Member.User.Username, "test")

	content := "test response"

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}
