package main

import (
	"regexp"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// BanwordPatterns contains list of banword patterns
type BanwordPatterns []*regexp.Regexp

// BuildPatterns create list of banword patterns
func BuildPatterns(banword_patterns []string) BanwordPatterns {
	patterns := make([]*regexp.Regexp, 0)
	for _, pattern := range banword_patterns {
		compiled, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Println("Error compiling", pattern, err)
		} else {
			patterns = append(patterns, compiled)
		}
	}
	return patterns
}

// FilterString checks if string has any banwords
func FilterString(s string, patterns BanwordPatterns) bool {
	containsBanword := false
	for _, pattern := range patterns {
		if pattern.FindString(s) != "" {
			containsBanword = true
			break
		}
	}

	return containsBanword
}

// FilterMessage checks if message has any banwords, if so, deletes it
func FilterMessage(s *discordgo.Session, m *discordgo.MessageCreate, patterns BanwordPatterns) {
	if FilterString(m.Content, patterns) {
		s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	}
}