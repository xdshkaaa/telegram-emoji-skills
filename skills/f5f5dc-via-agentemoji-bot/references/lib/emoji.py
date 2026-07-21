"""Generated from Telegram pack "Свои эмодзи • #f5f5dc (via @agentemoji_bot)" (1 emoji). Do not edit by hand.

Rendering requires parse_mode=HTML. Custom emoji display only when sent
from a Telegram Premium account (or an eligible channel), not via a
regular bot through the Bot API.
"""

from html import escape

EMOJI: dict[str, list[str]] = {
    "😻": ["5402470230121291343"],
}


def tg_emoji(emoji_id: str, fallback: str) -> str:
    """Return the <tg-emoji> HTML tag for parse_mode=HTML messages."""
    return f'<tg-emoji emoji-id="{emoji_id}">{escape(fallback)}</tg-emoji>'
