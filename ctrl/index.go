package ctrl

import (
	"github.com/cocotyty/summer-web-application-example/dao"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/cocotyty/summer"
	"time"
)

func init() {
	summer.Put(&Index{})
}

type Index struct {
	User *dao.User `sm:"*"`
}

func (i *Index) Page(ctx *gin.Context) {
	users, err := i.User.List()
	if err != nil {
		ctx.JSON(500, err)
	} else {
		ctx.JSON(200, users)
	}
}

func (i *Index) Time(ctx *gin.Context) {
	ctx.JSON(200, time.Now())
}
