package bus

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/insolar/insolar/instrumentation/inslogger"
)

func StartRouter(ctx context.Context, router *message.Router) {
	go func() {
		if err := router.Run(ctx); err != nil {
			inslogger.FromContext(ctx).Error("Error while running router", err)
		}
	}()
	<-router.Running()
}
