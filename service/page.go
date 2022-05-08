package service

import "myproject/repository"

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}
