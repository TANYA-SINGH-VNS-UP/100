/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025 Ashok Shau
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/TANYA-SINGH-VNS-UP/100
 */

package pkg

import (
	"github.com/TANYA-SINGH-VNS-UP/100/internal/config"
	"github.com/TANYA-SINGH-VNS-UP/100/internal/handlers"
	"github.com/TANYA-SINGH-VNS-UP/100/internal/vc"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func Init(client *tg.Client) error {
	for _, session := range config.Conf.SessionStrings {
		_, err := vc.Calls.StartClient(config.Conf.ApiId, config.Conf.ApiHash, session)
		if err != nil {
			return err
		}
	}

	vc.Calls.RegisterHandlers(client)
	handlers.LoadModules(client)
	return nil
}
