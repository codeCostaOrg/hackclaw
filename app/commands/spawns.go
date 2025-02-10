package commands

import (
	"fmt"

	"codecosta.com/hackclaw/app/models"
	"codecosta.com/hackclaw/app/utils"
	"github.com/bwmarrin/discordgo"
)

func SendSpawns(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	utils.LogCommand(interaction.Member.User.Username, "/spawns")

	options := interaction.ApplicationCommandData().Options
	selectedMap := options[0].Value.(string)
	var mapURL string

	switch selectedMap {
	case string(models.ZERO_DAM):
		mapURL = "https://i.imgur.com/rZ6XPzk.jpeg"
	case string(models.LAYALI_GROVE):
		mapURL = "https://i.imgur.com/Ew95QDg.jpeg"
	case string(models.BRAKKESH):
		mapURL = "https://i.imgur.com/7bYkyki.png"
	case string(models.SPACE_CITY):
		mapURL = "https://i.imgur.com/Qm3cHCu.png"
	default:
		errMessage := fmt.Sprintf("Sorry %s is not supported", selectedMap)
		utils.LogDiscordError("SendSpawns.selectedMap.default", errMessage)
		err := utils.DiscordRespondWithError(discord, interaction, errMessage)
		if err != nil {
			utils.LogDiscordError("SendSpawns.selectedMap.default.InteractionRespond", err.Error())
		}

		return
	}

	// embed map
	embeds := []*discordgo.MessageEmbed{
		{
			Title: selectedMap + " Spawns",
			Image: &discordgo.MessageEmbedImage{
				URL: mapURL,
			},
		},
	}

	err := discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Title:  selectedMap + " Spawns",
			Embeds: embeds,
		},
	})
	if err != nil {
		utils.LogDiscordError("SendSpawns.InteractionRespond", err.Error())
	}
}
