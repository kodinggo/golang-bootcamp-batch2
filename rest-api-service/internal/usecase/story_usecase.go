package usecase

import (
	"context"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
)

type storyUsecase struct {
	storyRepository model.StoryRepository
}

func NewStoryUsecase(storyRepository model.StoryRepository) model.StoryRepository {
	return &storyUsecase{storyRepository: storyRepository}
}

func (u *storyUsecase) Create(ctx context.Context, data model.Story) (*model.Story, error) {
	return u.storyRepository.Create(ctx, data)
}

func (u *storyUsecase) FindAll(ctx context.Context, opt model.StoryOpt) ([]*model.Story, int64, error) {
	return u.storyRepository.FindAll(ctx, opt)
}

func (u *storyUsecase) FindByID(ctx context.Context, id int64) (*model.Story, error) {
	// TODO
	panic("implement me")
}

func (u *storyUsecase) Update(ctx context.Context, id int64) error {
	// TODO
	panic("implement me")
}

func (u *storyUsecase) Delete(ctx context.Context, id int64) error {
	// TODO
	panic("implement me")
}
