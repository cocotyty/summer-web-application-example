package main

import (
	"github.com/cocotyty/summer"
	"github.com/cocotyty/summer-web-application-example/web"
)

const configPath = "./conf/app.toml"

func main() {
	summer.Strict()
	err := summer.TomlFile(configPath)
	if err != nil {
		panic(err)
	}
	// resolve dependencies
	summer.Start()

	web := summer.GetStoneWithName("WebConf").(*web.WebConf)

	web.Listen()
}
