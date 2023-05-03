// 使用 gin 生成 hello world
package main

import (
	"gin-gorm-mysql/models"
	"net/http"
	"strconv"

	"gin-gorm-mysql/servise"

	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化数据库
	models.Setup()
}

func main() {
	r := gin.Default()

	var wordServise servise.WordServise

	// 获取列表
	r.GET("/word/list", func(c *gin.Context) {
		pageNumStr, _ := c.GetQuery("pagenum")
		pageSizeStr, _ := c.GetQuery("pagesize")
		pageNum, _ := strconv.Atoi(pageNumStr)
		pageSize, _ := strconv.Atoi(pageSizeStr)

		if pageNum == 0 {
			pageNum = -1
		}

		if pageSize == 0 {
			pageSize = -1
		}

		offset := (pageNum - 1) * pageSize

		if pageSize == -1 && pageNum == -1 {
			offset = -1
		}

		words, err := wordServise.GetWords(pageSize, offset)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "data": words})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "获取失败"})

	})

	// 添加数据
	r.POST("/word", func(c *gin.Context) {
		var word models.Word
		err := c.ShouldBindJSON(&word)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
			return
		}

		if err = wordServise.AddWord(&word); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "添加失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "添加成功"})

	})

	// 通过 id 删除
	r.DELETE("/word/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
			return
		}
		err = wordServise.DeleteWordByID(id)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "删除失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "删除成功"})
	})

	// 更新
	r.PUT("/word", func(c *gin.Context) {
		var word models.Word
		err := c.ShouldBindJSON(&word)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
			return
		}

		if err = wordServise.UpdateWord(&word); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "更新失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "更新成功"})
	})

	// 监听8080端口的服务
	r.Run(":8080")
}
