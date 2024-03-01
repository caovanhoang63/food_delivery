package categorybiz

import (
	"context"
	"food-delivery/common"
	categorymodel "food-delivery/module/category/model"
)

type CreateCategoryStorage interface {
	Create(ctx context.Context, data *categorymodel.CategoryCreate) error
}

type createCategoryBiz struct {
	store CreateCategoryStorage
}

func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error {

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}
	return nil
}
