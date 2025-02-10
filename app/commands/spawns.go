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
	var mapEmbed *discordgo.MessageEmbed

	switch selectedMap {
	case string(models.ZERO_DAM):
		mapEmbed = &discordgo.MessageEmbed{
			Title: selectedMap + " Spawns",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://i.imgur.com/hN7W003.jpeg",
			},
		}
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
	embeds := []*discordgo.MessageEmbed{}
	embeds = append(embeds, mapEmbed)

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
