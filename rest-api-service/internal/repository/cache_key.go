package repository

import (
	"fmt"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
)

var storyListBucketKey = "storyList"

func newStoryDetailKey(id int64) string {
	return fmt.Sprintf("storyDetail:id:%d", id)
}

func newStoryListCacheKey(opt model.StoryOpt) string {
	return fmt.Sprintf("storyList:sort_by:%s:keyword:%s:cursor:%s:limit:%d", opt.SortBy, opt.Keyword, opt.Cursor, opt.Limit)
}
