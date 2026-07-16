# telegram-emoji-skills

Auto-generated [Claude Code](https://claude.com/claude-code) agent skills for Telegram
premium (custom) emoji packs. Each skill teaches an agent how to insert a specific
pack's emoji into Telegram content using the `<tg-emoji>` HTML tag.

Generated and published by [@agentemoji_bot](https://t.me/agentemoji_bot) — send it a
`t.me/addemoji/...` link or a message containing premium emoji, and it publishes a
skill here automatically.

## Structure

Each skill lives under `skills/<slug>/`:

```
skills/<slug>/
├── SKILL.md                       # agent instructions + usage context
└── references/
    └── emoji-catalog.md           # fallback char, emoji id, ready <tg-emoji> snippet
```

`<slug>` is `<telegram-user-id>-<pack-name>`.

## Install a skill

```bash
npx skills add https://github.com/xdshkaaa/telegram-emoji-skills/tree/main/skills/<slug>
```

The bot sends you this exact command after publishing.

## How skills get here

1. Send the bot a pack link or a message with premium emoji.
2. It extracts every custom emoji (id + fallback character).
3. It builds `SKILL.md` and `references/emoji-catalog.md` and commits them here.
4. Sending the same pack again merges new emoji into the same skill — no duplicates.

No AI description or categorization step — the catalog is a flat list: fallback
character, emoji id, and a copy-paste-ready `<tg-emoji emoji-id="...">fallback</tg-emoji>`
snippet.

## Note on contents

Pack names and Telegram user ids are part of the file paths in this repo, which is
public. The bot discloses this before first publish.
