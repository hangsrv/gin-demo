package util

import (
	"gin_demo/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"sync"
)

var store = sessions.NewCookieStore([]byte("session-secret"))
var lock = sync.Mutex{}

func Set(c *gin.Context, data model.Employee) error {
	lock.Lock()
	defer lock.Unlock()

	count := 0
	GetDb().Table((&model.Department{}).TableName()).Where("manager_name = ?", data.Username).Count(&count)
	if count > 0 {
		data.IsManger = 1
	}

	r := c.Request
	w := c.Writer
	session, err := store.Get(r, "user_info")
	if err != nil {
		return err
	}
	session.Values["e"] = data
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func Remove(c *gin.Context) error {
	lock.Lock()
	defer lock.Unlock()

	r := c.Request
	w := c.Writer

	session, err := store.Get(r, "user_info")
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func Get(c *gin.Context) (interface{}, error) {
	lock.Lock()
	defer lock.Unlock()

	r := c.Request
	session, err := store.Get(r, "user_info")
	if err != nil {
		return nil, err
	}
	return session.Values["e"], err
}
