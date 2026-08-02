/**
 * Generated from Telegram pack "Свои эмодзи • #9c27b0 (via @agentemoji_bot)" (5 emoji). Do not edit by hand.
 *
 * Rendering requires parse_mode=HTML. Custom emoji display only when sent
 * from a Telegram Premium account (or an eligible channel), not via a
 * regular bot through the Bot API.
 */

/** Maps a fallback character to the custom emoji IDs that use it. */
export const EMOJI: Record<string, readonly string[]> = {
  "⭐️": ["5438130961327893724", "5438179477278465764", "5438467046813773202"],
  "✈️": ["5438285837848583985"],
  "💫": ["5440679122474934976"],
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
