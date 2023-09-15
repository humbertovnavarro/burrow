package youtubemuse

import (
	"github.com/bwmarrin/discordgo"
)

func Play(videoId string, token string, guildId string, channelId string) (chan bool, error) {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	err = discord.Open()
	if err != nil {
		return nil, err
	}
	dgv, err := discord.ChannelVoiceJoin(guildId, channelId, false, true)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	stop := make(chan bool)
	dgv.Close()
	discord.Close()
	return stop, err
}
