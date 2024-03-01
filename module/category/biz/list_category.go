package categorybiz

import (
	"context"
	"food-delivery/common"
	categorymodel "food-delivery/module/category/model"
)

type ListCategoryStore interface {
	ListCategoryWithCondition(ctx context.Context,
		filter *categorymodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]categorymodel.Category, error)
}

type ListCategoryBiz struct {
	store ListCategoryStore
}

func NewListCategoryBiz(store ListCategoryStore) *ListCategoryBiz {
	return &ListCategoryBiz{store: store}
}

func (biz *ListCategoryBiz) ListCategory(ctx context.Context, filter *categorymodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]categorymodel.Category, error) {
	result, err := biz.store.ListCategoryWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(categorymodel.EntityName, err)
	}
	return result, nil
}
