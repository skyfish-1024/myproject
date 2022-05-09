package controller

import (
	"myproject/repository"
	"myproject/service"
	"myproject/util"
	"strconv"
	"time"
)

type PageData struct {
	Code int64       `json:"Code,omitempty"`
	Msg  string      `json:"Msg,omitempty"`
	Data interface{} `json:"Data,omitempty"`
}

var f *service.QueryPageInfoFlow

func QueryPageInfo(topicIdstr string) *PageData {
	f = new(service.QueryPageInfoFlow)
	topicId, err := strconv.ParseInt(topicIdstr, 10, 64)
	if err != nil {
		return &PageData{-1, "failed", nil}
	}
	pageInfo, err := f.Do(topicId)
	if err != nil {
		return &PageData{-1, "failed", nil}
	}
	return &PageData{0, "success", pageInfo}
}
func PublishPost(topicIdstr string, content string) *PageData {
	topicId, err := strconv.ParseInt(topicIdstr, 10, 64)
	if err != nil {
		return &PageData{-1, "failed", nil}
	}
	post := repository.Post{
		Id:         util.CreateId(),
		TopicId:    topicId,
		Content:    content,
		CreateTime: time.Now(),
	}
	err = repository.SavePost("./data/", post)
	if err != nil {
		return &PageData{-1, "failed", nil}
	}
	pageInfo, err := f.Do(topicId)
	if err != nil {
		return &PageData{-1, "failed", nil}
	}
	return &PageData{0, "success", pageInfo}
}
