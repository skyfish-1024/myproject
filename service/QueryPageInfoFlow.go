package service

import (
	"errors"
	"myproject/repository"
	"sync"
)

type QueryPageInfoFlow struct {
	PageInfo PageInfo
}

func (f *QueryPageInfoFlow) Do(topicId int64) (*PageInfo, error) {
	if err := f.checkParam(topicId); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(topicId); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return &f.PageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam(topicId int64) error {
	if _, ok := repository.TopicIndexMap[topicId]; !ok {
		return errors.New("topic is not exist")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo(topicId int64) error {
	var wg sync.WaitGroup
	wg.Add(2)
	//获取topic信息
	go func() {
		topicDao := repository.NewTopicDaoInstance()
		f.PageInfo.Topic = topicDao.QueryTopicById(topicId)
		defer wg.Done()
	}()
	//获取post列表
	go func() {
		defer wg.Done()
		postDao := repository.NewPostDaoInstance()
		f.PageInfo.PostList = postDao.QueryPostByTopicId(topicId)
	}()
	wg.Wait()
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	return nil
}
