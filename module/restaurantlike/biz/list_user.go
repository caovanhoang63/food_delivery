package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

type ListUserLikeRestaurantStore interface {
	GetUsersLikeRestaurant(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBiz(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{store: store}
}

func (biz *listUserLikeRestaurantBiz) ListUserLikeRestaurant(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUsersLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntiryName, err)
	}
	return users, nil
}
