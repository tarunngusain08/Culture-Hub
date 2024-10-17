package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/tarunngusain08/culturehub/pkg/log"
)

var (
	logger = log.New("http/rest/session")
	store  = cookie.NewStore([]byte("secret"))

	Middleware = sessions.Sessions("mysession", store)
)

func SaveUserPass(userName, password string, playerID uint, c *gin.Context) error {
	sess := sessions.Default(c)
	sess.Set("username", userName)
	sess.Set("password", password)
	return sess.Save()
}

func GetUserPass(c *gin.Context) (string, string, error) {
	sess := sessions.Default(c)
	usr, _ := sess.Get("username").(string)
	pss, _ := sess.Get("password").(string)
	return usr, pss, nil
}

func GetUsername(c *gin.Context) (string, error) {
	sess := sessions.Default(c)
	usr, _ := sess.Get("username").(string)
	return usr, nil
}
