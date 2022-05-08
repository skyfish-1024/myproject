package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
	"time"
)

type Post struct {
	Id         int64     `json:"Id,omitempty"`
	TopicId    int64     `json:"TopicId,omitempty"`
	Content    string    `json:"Content,omitempty"`
	CreateTime time.Time `json:"CreateTime"`
}
type PostDao struct {
}

var (
	postDao      *PostDao
	postOnce     sync.Once
	PostIndexMap = make(map[int64][]*Post)
)

func NewPostDaoInstance() *PostDao {
	topicOnce.Do(
		func() {
			postDao = new(PostDao)
		})
	return postDao
}
func (*PostDao) QueryPostByTopicId(id int64) []*Post {
	return PostIndexMap[id]
}

func InitPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		postTmpMap[post.TopicId] = append(postTmpMap[post.TopicId], &post)
	}
	PostIndexMap = postTmpMap
	return nil
}
