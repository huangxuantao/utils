package dal_util

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"utils/log_util"
)

type MemcacheConfig struct {
	Host string
	Port string
}

type Memcache struct {
	MC *memcache.Client
}

func connMC(cfg *MemcacheConfig) (*Memcache, error) {
	var MCInstance = Memcache{}
	MCInstance.MC = memcache.New(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	log_util.Logger.Debugf("MC/Connect: Connected to %s", cfg.Host)
	return &MCInstance, nil
}

func GetMC(cfg *MemcacheConfig) (*Memcache, error) {
	d, err := connMC(cfg)
	if err != nil {
		log_util.Logger.Errorf("MC/Connect:%s", err.Error())
		return nil, err
	}
	return d, nil
}
