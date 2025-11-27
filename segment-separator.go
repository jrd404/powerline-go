package main

import (
	pwl "github.com/justjanne/powerline-go/powerline"
)

func segmentSeparator(p *powerline) []pwl.Segment {
	return []pwl.Segment{{
		Name:       "separator",
		Content:    "",
		Foreground: p.theme.SeparatorFg,
		Background: 237, // dark gray
	}}
}
