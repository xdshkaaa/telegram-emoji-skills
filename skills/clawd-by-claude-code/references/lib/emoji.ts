/**
 * Generated from Telegram pack "Clawd by Claude Code" (16 emoji). Do not edit by hand.
 *
 * Rendering requires parse_mode=HTML. Custom emoji display only when sent
 * from a Telegram Premium account (or an eligible channel), not via a
 * regular bot through the Bot API.
 */

/** Maps a fallback character to the custom emoji IDs that use it. */
export const EMOJI: Record<string, readonly string[]> = {
  "🛹": ["5055310853969021493"],
  "⚽": ["5055315385159518663"],
  "👋": ["5055549589726168985"],
  "🏀": ["5055643774064002784"],
  "💃": ["5057549442463303394"],
  "⬆️": ["5057574022561138923"],
  "🏎": ["5057580404882541671"],
  "🚶": ["5057604164641621832"],
  "🔍": ["5057657362106550145"],
  "📚": ["5057700728391338135"],
  "🦀": ["5057764177943201734"],
  "🐾": ["5057784463073740822"],
  "🎺": ["5057812758318287275"],
  "🤩": ["5057887598123419740"],
  "🤔": ["5058038712252762170"],
  "👉": ["5058052902824707746"],
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
