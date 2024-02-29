package restaurantrepository

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
	"log"
)

type ListRestaurantStore interface {
	ListRestaurantWithCondition(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore LikeRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store, likeStore}
}

func (repo *listRestaurantRepo) List(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	restaurants, err := repo.store.ListRestaurantWithCondition(ctx, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	ids := make([]int, len(restaurants))

	for i := range restaurants {
		ids[i] = restaurants[i].Id
	}

	likes, err := repo.likeStore.GetRestaurantLikes(ctx, ids)
	if err != nil {
		log.Println(err)
		return restaurants, nil
	}

	for i := range restaurants {
		restaurants[i].LikeCount = likes[restaurants[i].Id]
	}

	return restaurants, nil
}
