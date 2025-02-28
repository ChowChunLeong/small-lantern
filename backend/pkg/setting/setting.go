package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type App struct {
	Name   string
	Origin []string
	Port   string
}

var AppSetting = &App{}

var cfg *ini.File

// Setup initialize the configuration instance

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
