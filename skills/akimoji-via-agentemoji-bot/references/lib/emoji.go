// Package emoji: generated from Telegram pack "Akimoji • 🌿 Лес (via @agentemoji_bot)" (24 emoji).
// Copy into your project; rename the package if needed. Do not edit by hand.
//
// Rendering requires parse_mode=HTML. Custom emoji display only when sent
// from a Telegram Premium account (or an eligible channel), not via a
// regular bot through the Bot API.
package emoji

import "html"

// Emoji maps a fallback character to the custom emoji IDs that use it.
var Emoji = map[string][]string{
	"✅": {"5388706152722769105", "5391015281759917747", "5391347823897777186"},
	"🧭": {"5388878801818133330"},
	"🎥": {"5388900053316315257"},
	"🏃‍♂️": {"5388991140982729878"},
	"💤": {"5389083568678936542"},
	"🖥": {"5390885732661370326", "5391111609286432918"},
	"🍏": {"5390917807477135334", "5391023863104575306", "5391210732836657337"},
	"📲": {"5390959077817885562"},
	"🏋️‍♂️": {"5390970188898277623"},
	"🖼️": {"5390974526815261753"},
	"✈️": {"5391068178577133793"},
	"⭐️": {"5391101700796881561"},
	"🍓": {"5391155813089846747"},
	"⚙️": {"5391156977025980026"},
	"🎚️": {"5391201910973832565"},
	"🍆": {"5391278232542682086"},
	"🔦": {"5391323475728176620"},
	"🍑": {"5391368255057207328"},
	"🔕": {"5391369663806479044"},
}

// TgEmoji returns the <tg-emoji> HTML tag for parse_mode=HTML messages.
func TgEmoji(id, fallback string) string {
	return `<tg-emoji emoji-id="` + id + `">` + html.EscapeString(fallback) + `</tg-emoji>`
}
