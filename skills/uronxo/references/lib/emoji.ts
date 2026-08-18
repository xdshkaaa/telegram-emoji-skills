/**
 * Generated from Telegram pack "@uronxo" (15 emoji). Do not edit by hand.
 *
 * Rendering requires parse_mode=HTML. Custom emoji display only when sent
 * from a Telegram Premium account (or an eligible channel), not via a
 * regular bot through the Bot API.
 */

/** Maps a fallback character to the custom emoji IDs that use it. */
export const EMOJI: Record<string, readonly string[]> = {
  "🔍": ["5408847075439845717", "5413729834780108071"],
  "🧴": ["5408891008660316195"],
  "📞": ["5409054367741419594", "5409119986251765622"],
  "🏷️": ["5409098713278750910"],
  "⛅": ["5409236139347320611"],
  "👑": ["5411157230974184252"],
  "💸": ["5411336571628592542"],
  "🪧": ["5411534045634929067"],
  "😈": ["5413368086864636617"],
  "🤩": ["5413403004948751950"],
  "🖤": ["5413446628931576835"],
  "⌨️": ["5413451224546584209"],
  "🛡": ["5413452500151870496"],
};

function escapeHtml(s: string): string {
  return s
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;");
}

/** Returns the <tg-emoji> HTML tag for parse_mode=HTML messages. */
export function tgEmoji(id: string, fallback: string): string {
  return `<tg-emoji emoji-id="${id}">${escapeHtml(fallback)}</tg-emoji>`;
}
