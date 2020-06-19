package command

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"

	assets "vcListBot/command/assets"
)

const textLength = 8

func List(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	if message.Content != "!list" {
		return
	}
	channel, guild, err := assets.GetGuildData(session, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	memberCount := guild.MemberCount
	voiceJoinNumber, voiceMuteNumber := GetVoiceStates(channel, guild)
	utterance := " ***限界リスト*** \n```asciidoc\n= 現在の状況 =\n"
	utterance += AllMember(memberCount) + InVoiceMembers(voiceJoinNumber) + MuteMembers(voiceMuteNumber)
	utterance += VoiceMemberRate(memberCount, voiceJoinNumber) + MuteRate(voiceJoinNumber, voiceMuteNumber)
	utterance += "```"
	session.ChannelMessageSend(message.ChannelID, utterance)
}

func GetVoiceStates(channel *discordgo.Channel, guild *discordgo.Guild) (voiceJoinNumber int, voiceMuteNumber int) {
	voiceJoinNumber = len(guild.VoiceStates)
	for _, vs := range guild.VoiceStates {
		if vs.SelfMute {
			voiceMuteNumber++
		}
	}
	return
}

func AllMember(members int) string {
	return assets.PaddingRight("鯖人数", textLength, "　") + ":: " + strconv.Itoa(members) + " 人\n"
}

func InVoiceMembers(members int) string {
	return assets.PaddingRight("通話人数", textLength, "　") + ":: " + strconv.Itoa(members) + " 人\n"
}

func MuteMembers(members int) string {
	return assets.PaddingRight("ミュート人数", textLength, "　") + ":: " + strconv.Itoa(members) + " 人\n"
}

func VoiceMemberRate(memberCount int, voiceJoinNumber int) string {
	rate := float64(voiceJoinNumber) / float64(memberCount)
	return assets.PaddingRight("通話率", textLength, "　") + ":: " + assets.FormatRateNum(rate) + " %\n"
}

func MuteRate(voiceJoinNumber int, voiceMuteNumber int) string {
	rate := float64(voiceMuteNumber) / float64(voiceJoinNumber)
	return assets.PaddingRight("ミュート率", textLength, "　") + ":: " + assets.FormatRateNum(rate) + " %\n"
}