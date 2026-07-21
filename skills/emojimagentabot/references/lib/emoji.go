// Package emoji: generated from Telegram pack "секс на пляже @emojimagentabot" (17 emoji).
// Copy into your project; rename the package if needed. Do not edit by hand.
//
// Rendering requires parse_mode=HTML. Custom emoji display only when sent
// from a Telegram Premium account (or an eligible channel), not via a
// regular bot through the Bot API.
package emoji

import "html"

// Emoji maps a fallback character to the custom emoji IDs that use it.
var Emoji = map[string][]string{
	"🙂": {"5208710077711687793", "5208745549846582946", "5208771577348400835", "5208891140647986664", "5211060940946054224", "5435890491637932045", "5436072748575139555", "5436099033774988363", "5436294987362900460", "5436323879607901053", "5436333573349087916", "5438293409875926513", "5438611739967005331", "5449454174452359566", "5453884506232627052", "5454008802586174421", "5456123734677036423"},
}

// TgEmoji returns the <tg-emoji> HTML tag for parse_mode=HTML messages.
func TgEmoji(id, fallback string) string {
	return `<tg-emoji emoji-id="` + id + `">` + html.EscapeString(fallback) + `</tg-emoji>`
}
