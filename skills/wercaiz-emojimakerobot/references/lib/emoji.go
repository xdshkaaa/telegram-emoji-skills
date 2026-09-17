// Package emoji: generated from Telegram pack "@wercaiz | @emojimakerobot" (3 emoji).
// Copy into your project; rename the package if needed. Do not edit by hand.
//
// Rendering requires parse_mode=HTML. Custom emoji display only when sent
// from a Telegram Premium account (or an eligible channel), not via a
// regular bot through the Bot API.
package emoji

import "html"

// Emoji maps a fallback character to the custom emoji IDs that use it.
var Emoji = map[string][]string{
	"✨": {"5210728411693102772", "5210800275085896375", "5213261158202515151"},
}

// TgEmoji returns the <tg-emoji> HTML tag for parse_mode=HTML messages.
func TgEmoji(id, fallback string) string {
	return `<tg-emoji emoji-id="` + id + `">` + html.EscapeString(fallback) + `</tg-emoji>`
}
