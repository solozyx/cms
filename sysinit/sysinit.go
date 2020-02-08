package sysinit

import (
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/solozyx/cms/utils"
)

func init() {
	// init cache
	utils.Cache = cache.New(60*time.Minute, 120*time.Minute)
	// init db
	initDB()
}
