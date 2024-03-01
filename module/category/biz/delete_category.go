package categorybiz

import (
	"context"
	"food-delivery/common"
	categorymodel "food-delivery/module/category/model"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type DeleteCategoryStore interface {
	Delete(ctx context.Context, id int) error
	FindCategoryWithConditions(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*categorymodel.Category, error)
}

type deleteCategoryBiz struct {
	store DeleteCategoryStore
}

func NewDeleteCategoryBiz(store DeleteCategoryStore) *deleteCategoryBiz {
	return &deleteCategoryBiz{store}
}

func (biz *deleteCategoryBiz) DeleteCategory(ctx context.Context, id int) error {
	oldData, err := biz.store.FindCategoryWithConditions(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(categorymodel.EntityName, err)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	return nil
}
