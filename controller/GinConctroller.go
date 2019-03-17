package controller

import (
	"github.com/gin-gonic/gin"
	. "home/model"
	"log"
	"net/http"
	"strconv"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404 ,page not exists!",
	})
}

func Add(c *gin.Context) {
	user := new(User)
	err := c.BindJSON(user)
	if err != nil {
		log.Fatal(err)
	}
	insert, flag := Insert(user.Name, user.Balance)
	if flag {
		c.JSON(200, gin.H{
			"msg":      "添加成功",
			"affected": insert,
			"state":    1,
		})
	} else {
		c.JSON(200, gin.H{"msg": "添加失败", "state": 0})
	}
}

func Delete(c *gin.Context) {
	id := c.Query("id")
	uid, _ := strconv.ParseInt(id, 10, 64)
	Del(uid)
}

func Edit(c *gin.Context) {
	query := c.Query("id")
	id, _ := strconv.ParseInt(query, 10, 64)
	user := new(User)
	c.BindJSON(user)
	update := Update(id, user)
	c.JSON(200, gin.H{
		"result": update,
	})
}

func Find(c *gin.Context) {
	query := c.Query("id")
	id, _ := strconv.ParseInt(query, 10, 64)
	info := GetInfo(id)
	c.JSON(200, gin.H{
		"user": info,
	})
}
