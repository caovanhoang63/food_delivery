package categorybiz

import (
	"context"
	"food-delivery/common"
	categorymodel "food-delivery/module/category/model"
)

type FindCategoryStore interface {
	FindCategoryWithConditions(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*categorymodel.Category, error)
}

type findCategoryBiz struct {
	store FindCategoryStore
}

func NewFindCategoryBiz(store FindCategoryStore) *findCategoryBiz {
	return &findCategoryBiz{store: store}
}

func (biz *findCategoryBiz) FindById(ctx context.Context,
	id int,
	moreKeys ...string) (*categorymodel.Category, error) {
	result, err := biz.store.FindCategoryWithConditions(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrEntityNotFound(categorymodel.EntityName, err)
	}
	return result, nil
}
