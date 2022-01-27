package gh

import (
	_ "embed"
	"encoding/json"
)

//go:embed emojis.json
var emojiJson []byte

// Emojis Github emoji shortname to Unicode
var Emojis emojis

func init() {
	_ = json.Unmarshal(emojiJson, &Emojis)
}

type emojis []Emoji

// Emoji is a emoji info struct
type Emoji struct {
	Emoji     string `json:"emoji"`
	Name      string `json:"name"`
	Shortname string `json:"shortname"`
	Unicode   string `json:"unicode"`
	Html      string `json:"html"`
	Category  string `json:"category"`
	Order     string `json:"order"`
}

// Emoji2Shortname convert emoji to shortname
func (e emojis) Emoji2Shortname(emoji string) string {
	for _, e := range e {
		if e.Emoji == emoji {
			return e.Shortname
		}
	}
	return ""
}

// Shortname2Emoji convert shortname to emoji
func (e emojis) Shortname2Emoji(shortname string) string {
	for _, e := range e {
		if e.Shortname == shortname {
			return e.Emoji
		}
	}
	return ""
}

// GetEmoji get emoji info
func (e emojis) GetEmoji(ShortnameOrEmoji string) string {
	for _, e := range e {
		if e.Shortname == ShortnameOrEmoji || e.Emoji == ShortnameOrEmoji {
			return e.Emoji
		}
	}
	return ""
}
