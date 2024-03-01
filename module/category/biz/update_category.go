package categorybiz

import (
	"context"
	"food-delivery/common"
	categorymodel "food-delivery/module/category/model"
)

type UpdateCategoryStore interface {
	Update(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error
	FindCategoryWithConditions(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*categorymodel.Category, error)
}

type updateCategoryBiz struct {
	store UpdateCategoryStore
}

func NewUpdateCategoryBiz(store UpdateCategoryStore) *updateCategoryBiz {
	return &updateCategoryBiz{store: store}
}

func (biz *updateCategoryBiz) Update(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	_, err := biz.store.FindCategoryWithConditions(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return common.RecordNotFound
		}
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}

	return nil
}
