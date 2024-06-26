package officialaccount

import (
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/os/gcache"
)

type OfficialAccount struct {
	config      Config
	accessToken auth.AccessToken
	Logger      *log.Logger
	Cache       *gcache.Cache
}

func New(config Config) *OfficialAccount {
	if config.Cache == nil {
		config.Cache = cache.New("ewa.wechat.officialaccount")
	}
	if config.Logger == nil {
		config.Logger = log.New()
		if config.Logger.LogPath != "" {
			if err := config.Logger.SetPath(config.Logger.LogPath); err != nil {
				panic(fmt.Sprintf("[officialaccount] set log path '%s' error: %v", config.Logger.LogPath, err))
			}
		}

		// default set close debug / close stdout print
		config.Logger.LogStdout = false

		// config.Logger.SetDebug(false)
		// config.Logger.SetStdoutPrint(false)
	}
	var oa = &OfficialAccount{
		config: config,
		Logger: config.Logger,
		Cache:  config.Cache,
	}
	oa.accessToken = oa.getDefaultAccessToken()
	return oa
}
