package event

import (
  "context"

  coretypes "github.com/cometbft/cometbft/rpc/core/types"
  "github.com/mingi3442/go-grpc/log"
  "github.com/mingi3442/tendermint-ws/utils"
)

func HandleEvents(ctx context.Context, events <-chan coretypes.ResultEvent) {
  for {
    select {
    case event := <-events:
      // msg := fmt.Sprint("Received event:", event)
      log.Log(log.INFO, "Received event")
      // utils.ParseJson(event)
      s, _ := utils.ParseJson(event)
      log.Log(log.DEBUG, s)
    case <-ctx.Done():
      log.Log(log.DEBUG, "Event processing stopped due to timeout or cancellation")
      return
    }
  }
}
