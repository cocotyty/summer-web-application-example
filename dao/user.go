package dao

import (
	"github.com/cocotyty/summer-web-application-example/model"
	"github.com/jmoiron/sqlx"
	"github.com/cocotyty/summer"
	"github.com/cocotyty/summer-web-application-example/provider"
)

func init() {
	summer.Put(&User{}, provider.DBMain)
}

type User struct {
	DB *sqlx.DB `sm:"@.(.)"`
}

func (u *User) List() (users []*model.User, err error) {
	users = []*model.User{}
	err = u.DB.Select(&users, "select * from user  order by id limit 10")
	return
}
