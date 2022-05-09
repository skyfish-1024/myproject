package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

func insertdata() {
	filePath := "./data/"
	topic := Topic{
		Id:         1,
		Title:      "这是第一篇topic",
		Content:    "啦啦啦啦啦啦啦",
		CreateTime: time.Now(),
	}
	post := Post{
		Id:         1,
		TopicId:    1,
		Content:    "真不错",
		CreateTime: time.Now(),
	}
	SaveTopic(filePath, topic)
	SavePost(filePath, post)
}
func SaveTopic(filePath string, topic Topic) error {
	file, err := os.OpenFile(filePath+"topic", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//序列化topic
	buf, err := json.Marshal(topic)
	if err != nil {
		return err
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.Write(buf)
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	return nil
}

func SavePost(filePath string, post Post) error {
	var wg sync.RWMutex
	file, err := os.OpenFile(filePath+"post", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	wg.Lock()
	PostIndexMap[post.TopicId] = append(PostIndexMap[post.TopicId], &post)
	wg.Unlock()
	//序列化post
	buf, err := json.Marshal(post)
	if err != nil {
		return err
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString("\n")
	write.Write(buf)
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	return nil
}
