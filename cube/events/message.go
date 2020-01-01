package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/cacher"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/wait_for_message"
	"time"
)

func init() {
	// Defines the message handler.
	Events = append(Events, func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Handle message waiters.
		go messagewaiter.MessageWaitHandler(m.Message)

		go func() {
			// Gets the start time.
			StartTime := time.Now()

			// Gets the channel.
			channel, err := cacher.GetChannel(m.ChannelID, s)
			if err != nil {
				// TODO: Error report here to Sentry.
				return
			}

			// If this is a DM, ignore it.
			if channel.Type == discordgo.ChannelTypeGroupDM || channel.Type == discordgo.ChannelTypeDM {
				return
			}

			// Fire up the commands processor in a new thread.
			go commandprocessor.Processor(m.Message, channel, s, &StartTime)
		}()
	})
}
