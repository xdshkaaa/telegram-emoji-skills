---
name: 339193247-topic-icons
description: Insert premium custom emoji from "Topic Icons" into Telegram content. Use when writing Telegram постов и сообщений в Telegram that should include premium emoji, or when the user mentions "Topic Icons" emoji.
---

# Topic Icons — Telegram premium emoji

## How to insert

- Requires `parse_mode=HTML` on the Telegram message.
- Syntax: `<tg-emoji emoji-id="CUSTOM_EMOJI_ID">FALLBACK</tg-emoji>`
- `FALLBACK` must be the plain emoji character shown to non-premium viewers or clients that don't support custom emoji.
- Custom emoji render only when sent from a Telegram Premium account (or an eligible channel). Regular bots generally cannot send them via the Bot API — this skill targets user-authored content, not bot replies.

## Where these are used

постов и сообщений в Telegram

## Catalog

160 emoji, fallback + id + ready snippet: [references/emoji-catalog.md](references/emoji-catalog.md)

Pick by fallback character, then copy the emoji-id exactly — IDs are long numeric strings and must not be altered.
