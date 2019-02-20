package config

import (
	"github.com/gitstliu/log4go"
	"github.com/pelletier/go-toml"
)

type Configure struct {

	//System config
	ServicePort  int
	WorkerId     int64
	DatacenterId int64
	IDBufferSize int64
}

var configure *Configure

func GetConfigure() *Configure {

	return configure
}

func LoadConfigure(fileName string) error {

	config, err := toml.LoadFile(fileName)

	if err != nil {
		return err
	}

	conf := Configure{}
	conf.ServicePort = int(config.GetDefault("sysconf.ServicePort", "60000").(int64))
	conf.WorkerId = config.Get("sysconf.WorkerId").(int64)
	conf.DatacenterId = config.Get("sysconf.DatacenterId").(int64)
	conf.IDBufferSize = config.Get("sysconf.IDBufferSize").(int64)

	configure = &conf

	log4go.Debug("Config : %v", conf)

	return nil
}
