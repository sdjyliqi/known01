package utils

import "strings"

const sepChar = ","

func ChannelList(channels string) []string {
	return strings.Split(channels, sepChar)
}
