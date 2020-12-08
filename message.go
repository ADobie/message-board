package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Msg struct {
	gorm.Model
	Message  string `json:"message" gorm:"type:varchar(255)"`
	Username string `json:"username" gorm:"type:varchar(255)"`
}

func addMsg(c *gin.Context) {
	var addMessage Msg
	err := c.BindJSON(&addMessage)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "Msg": "请求出错！"})
		return
	}
	Db.Create(&addMessage)
	c.JSON(200, gin.H{"code": "200", "Msg": "留言成功！"})
}

func delMsg(c *gin.Context) {
	var delMessage Msg
	err := c.BindJSON(&delMessage)
	if err != nil {
		c.JSON(400, gin.H{"code": "400", "Msg": "请求出错！"})
		return
	}
	if getUsername(c) == delMessage.Username {
		Db.Table("msg").Delete(delMessage)
		c.JSON(200, gin.H{"code": "200", "Msg": "删除成功！"})
	} else {
		c.JSON(200, gin.H{"code": "400", "Msg": "不能删除他人留言！"})
	}
}
