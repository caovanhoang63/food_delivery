package subcriber

import (
	"context"
	"food-delivery/common"
	"food-delivery/component/appctx"
	"food-delivery/component/asyncjob"
	"food-delivery/pubsub"
	"log"
)

type consumerJob struct {
	Title   string
	Handler func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx appctx.AppContext
}

func NewEngine(appCtx appctx.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appCtx}
}

func (engine *consumerEngine) Start() error {
	if err := engine.startSubTopic(common.TopicIncreaseLikeCountWhenUserLikeRestaurant, true,
		IncreaseLikeCountWhenUserLikeRestaurant(engine.appCtx, context.Background())); err != nil {
		log.Println("Err:", err)
	}

	if err := engine.startSubTopic(common.TopicDecreaseLikeCountWhenUserDislikeRestaurant, true,
		DecreaseLikeCountWhenUserLikeRestaurant(engine.appCtx, context.Background())); err != nil {
		log.Println("Err:", err)
	}
	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Printf("Set up consumer for: %s", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Printf("running job for %s for. Value %s \n", job.Title, message.Data())
			return job.Handler(ctx, message)
		}
	}

	go func() {
		common.AppRecover()
		for {
			msg := <-c
			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}
			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println("Err:", err)
			}
		}
	}()

	return nil
}
