package event

import (
  "context"
  "fmt"

  coretypes "github.com/cometbft/cometbft/rpc/core/types"
  "github.com/mingi3442/tendermint-ws/utils"
)

func HandleEvents(ctx context.Context, events <-chan coretypes.ResultEvent) {
  for {
    select {
    case event := <-events:
      fmt.Println("Received event:", event)
      utils.ParseJson(event)
    case <-ctx.Done():
      fmt.Println("Event processing stopped due to timeout or cancellation")
      return
    }
  }
}
