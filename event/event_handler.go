package event

import (
  "context"
  // "encoding/base64"
  "fmt"
  types "github.com/cometbft/cometbft/types"
  // "reflect"

  coretypes "github.com/cometbft/cometbft/rpc/core/types"
  "github.com/mingi3442/go-grpc/log"
  "github.com/mingi3442/tendermint-ws/utils"
)

func HandleEvents(ctx context.Context, txCh <-chan coretypes.ResultEvent) {
  for {
    select {
    case event := <-txCh:
      log.Log(log.INFO, "Received event")

      // fmt.Printf("Event Data Type: %s\n", reflect.TypeOf(event.Data))
      // fmt.Printf("Event Data: %+v\n", event.Data)

      // s, err := utils.ParseJson(event)
      // if err != nil {
      //   log.Log(log.ERROR, fmt.Sprintf("Failed to parse event to JSON: %v", err))
      // } else {
      //   log.Log(log.DEBUG, s)
      // }

      // 이벤트 데이터가 types.EventDataTx 타입인지 확인하고 처리
      if txEvent, ok := event.Data.(types.EventDataTx); ok {
        log.Log(log.DEBUG, "Received EventDataTx data")

        // 트랜잭션 데이터를 디코딩
        decodedTx, err := utils.DecodeTxData(txEvent.Tx)
        if err != nil {
          msg := fmt.Sprintf("Failed to decode transaction data: %v", err)
          log.Log(log.ERROR, msg)
        } else {
          msg := fmt.Sprintf("\n---Decoded Transaction---\n%s\n", decodedTx)
          log.Log(log.INFO, msg)
        }
      } else {
        log.Log(log.WARN, "Unknown event data type")
      }

    case <-ctx.Done():
      log.Log(log.DEBUG, "Event processing stopped due to timeout or cancellation")
      return
    }
  }
}
