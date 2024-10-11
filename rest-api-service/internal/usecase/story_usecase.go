package usecase

import (
	"context"
	"errors"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
)

type storyUsecase struct {
	storyRepository model.StoryRepository
	userRepository  model.UserRepository
}

func NewStoryUsecase(storyRepository model.StoryRepository, userRepository model.UserRepository) model.StoryRepository {
	return &storyUsecase{storyRepository: storyRepository, userRepository: userRepository}
}

func (u *storyUsecase) Create(ctx context.Context, data model.Story) (story *model.Story, err error) {
	user, err := u.userRepository.FindByID(ctx, data.AuthorID)
	if err != nil {
		return
	}

	if !model.AllowedRolesToManageStory[user.Role] {
		err = errors.New("user does not have access")
		return
	}

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
