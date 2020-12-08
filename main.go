package main

//import "fmt"
import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database()
	store := cookie.NewStore([]byte("username"))
	r.Use(sessions.Sessions("session", store))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "message board")
	})
	r.POST("/register", func(c *gin.Context) {
		register(c)
	})
	r.POST("/login", func(c *gin.Context) {
		login(c)
	})
	r.GET("/logout", func(c *gin.Context) {
		logout(c)
	})
	r.POST("/addMsg", func(c *gin.Context) {
		if getSession(c) {
			addMsg(c)
		} else {
			c.JSON(400, gin.H{"code": "400", "Msg": "请先登录！"})
		}
	})
	r.GET("/delMsg", func(c *gin.Context) {
		if getSession(c) {
			delMsg(c)
		} else {
			c.JSON(400, gin.H{"code": "400", "Msg": "请先登录！"})
		}
	})

	r.Run("4444")
}
