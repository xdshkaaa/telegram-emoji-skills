// Package emoji: generated from Telegram pack "Wizard Menu Emoji Pack" (33 emoji).
// Copy into your project; rename the package if needed. Do not edit by hand.
//
// Rendering requires parse_mode=HTML. Custom emoji display only when sent
// from a Telegram Premium account (or an eligible channel), not via a
// regular bot through the Bot API.
package emoji

import "html"

// Emoji maps a fallback character to the custom emoji IDs that use it.
var Emoji = map[string][]string{
	"⭐️": {"5465175670311131821"},
	"📶": {"5465207538968468071"},
	"👤": {"5465217249889527344"},
	"👑": {"5465246962473282754"},
	"❗️": {"5465251768541685385"},
	"💰": {"5465419517079365890", "5467691129577184774", "5469806452510075469", "5469973475198283469", "5470070923711260165", "5470150037008849880"},
	"ℹ️": {"5465435292494245199"},
	"📊": {"5465573925448623342"},
	"⬅️": {"5465603543543097854"},
	"💵": {"5465665262223141614"},
	"💸": {"5467377188942688107"},
	"🔄": {"5467401146270261486"},
	"⏳": {"5467414039762084308"},
	"💳": {"5467417346886903203"},
	"🔑": {"5467560060060214180"},
	"🔗": {"5467577596411684383"},
	"❓": {"5467677372796937209"},
	"🌐": {"5467723925947458846"},
	"📔": {"5467735423574911159"},
	"💎": {"5469710679034341813"},
	"💬": {"5469781271116817074"},
	"💡": {"5469828262354002935"},
	"❌": {"5469837230245718750"},
	"⚠️": {"5469890071228360253"},
	"📱": {"5469912065755880065"},
	"➡️": {"5469943053944920520"},
	"🚀": {"5470016527950455654"},
	"🎁": {"5474546335468267391"},
}

// TgEmoji returns the <tg-emoji> HTML tag for parse_mode=HTML messages.
func TgEmoji(id, fallback string) string {
	return `<tg-emoji emoji-id="` + id + `">` + html.EscapeString(fallback) + `</tg-emoji>`
}
