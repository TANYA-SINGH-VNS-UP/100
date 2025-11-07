/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025 Ashok Shau
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/TANYA-SINGH-VNS-UP/100
 */

package handlers

import (
	"github.com/TANYA-SINGH-VNS-UP/100/internal/core/cache"
	"github.com/TANYA-SINGH-VNS-UP/100/internal/core/db"
	"github.com/TANYA-SINGH-VNS-UP/100/internal/lang"
	"github.com/TANYA-SINGH-VNS-UP/100/internal/vc"

	"github.com/amarnathcjd/gogram/telegram"
)

// skipHandler handles the /skip command.
func skipHandler(m *telegram.NewMessage) error {
	chatID, _ := getPeerId(m.Client, m.ChatID())
	ctx, cancel := db.Ctx()
	defer cancel()
	langCode := db.Instance.GetLang(ctx, chatID)
	if !cache.ChatCache.IsActive(chatID) {
		_, _ = m.Reply(lang.GetString(langCode, "no_track_playing"))
		return nil
	}

	_ = vc.Calls.PlayNext(chatID)
	return nil
}
