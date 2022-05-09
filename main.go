package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myproject/controller"
	"myproject/repository"
	"os"
)

func main() {
	if err := Init("./data/"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//insertdata()
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	r.POST("/community/post/do", func(c *gin.Context) {
		topicId, _ := c.GetPostForm("topic_id")
		content, _ := c.GetPostForm("content")
		data := controller.PublishPost(topicId, content)
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
func Init(filePath string) error {
	if err := repository.InitTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := repository.InitPostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}
