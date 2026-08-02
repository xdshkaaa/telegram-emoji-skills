"""Generated from Telegram pack "Свои эмодзи • #9c27b0 (via @agentemoji_bot)" (5 emoji). Do not edit by hand.

Rendering requires parse_mode=HTML. Custom emoji display only when sent
from a Telegram Premium account (or an eligible channel), not via a
regular bot through the Bot API.
"""

from html import escape

EMOJI: dict[str, list[str]] = {
    "⭐️": ["5438130961327893724", "5438179477278465764", "5438467046813773202"],
    "✈️": ["5438285837848583985"],
    "💫": ["5440679122474934976"],
}


def tg_emoji(emoji_id: str, fallback: str) -> str:
    """Return the <tg-emoji> HTML tag for parse_mode=HTML messages."""
    return f'<tg-emoji emoji-id="{emoji_id}">{escape(fallback)}</tg-emoji>'
