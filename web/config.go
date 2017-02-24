package web

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/cocotyty/summer-web-application-example/ctrl"
	"github.com/cocotyty/summer"
)

func init() {
	summer.Add("WebConf", &WebConf{})
}

type WebConf struct {
	// inject index controller from summer
	Index *ctrl.Index `sm:"*"`
	// inject listen address from config file
	ListenAddr string `sm:"#.web.listen"`

	gin *gin.Engine
}

func (conf *WebConf) Init() {
	// create engine
	conf.gin = gin.New()

	conf.routes()
}

// config routes
func (conf *WebConf) routes() {
	g := conf.gin
	// config web
	g.GET("/", conf.Index.ListUser)
	g.GET("/time", conf.Index.Time)
}


func (conf *WebConf) Listen() (err error) {
	return conf.gin.Run(conf.ListenAddr)
}
