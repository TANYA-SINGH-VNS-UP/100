/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025 Ashok Shau
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/TANYA-SINGH-VNS-UP/100
 */

package handlers

import (
	"fmt"
	"strings"

	"github.com/TANYA-SINGH-VNS-UP/100/internal/core"
	"github.com/TANYA-SINGH-VNS-UP/100/internal/core/cache"
	"github.com/TANYA-SINGH-VNS-UP/100/internal/core/db"
	"github.com/TANYA-SINGH-VNS-UP/100/internal/lang"

	"github.com/Laky-64/gologging"
	"github.com/amarnathcjd/gogram/telegram"
)

func settingsHandler(m *telegram.NewMessage) error {
	if m.IsPrivate() {
		return nil
	}

	ctx, cancel := db.Ctx()
	defer cancel()

	chatID, _ := getPeerId(m.Client, m.ChatID())
	admins, err := cache.GetAdmins(m.Client, chatID, false)
	if err != nil {
		return err
	}

	// Check if user is admin
	var isAdmin bool
	for _, admin := range admins {
		if admin.User.ID == m.Sender.ID {
			isAdmin = true
			break
		}
	}
	if !isAdmin {
		return nil
	}
	langCode := db.Instance.GetLang(ctx, chatID)
	// Get current settings
	getPlayMode := db.Instance.GetPlayMode(ctx, chatID)
	getAdminMode := db.Instance.GetAdminMode(ctx, chatID)

	text := fmt.Sprintf(lang.GetString(langCode, "settings_header"),
		m.Chat.Title, getPlayMode, getAdminMode)

	_, err = m.Reply(text, telegram.SendOptions{
		ReplyMarkup: core.SettingsKeyboard(getPlayMode, getAdminMode),
	})
	return err
}
