package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
	"time"
)

type Topic struct {
	Id         int64     `json:"Id,omitempty"`
	Title      string    `json:"Title,omitempty"`
	Content    string    `json:"Content,omitempty"`
	CreateTime time.Time `json:"CreateTime"`
}
type TopicDao struct {
}

var (
	topicDao      *TopicDao
	topicOnce     sync.Once
	TopicIndexMap = make(map[int64]*Topic)
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = new(TopicDao)
		})
	return topicDao
}
func (*TopicDao) QueryTopicById(id int64) *Topic {
	return TopicIndexMap[id]
}

func InitTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	TopicIndexMap = topicTmpMap
	return nil
}
