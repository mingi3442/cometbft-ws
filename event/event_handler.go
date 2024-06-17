package event

import (
  "context"
  // "encoding/base64"

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
      log.Info("Received event")
      // fmt.Printf("Event Data Type: %s\n", reflect.TypeOf(event.Data))
      // fmt.Printf("Event Data: %+v\n", event.Data)

      // s, err := utils.ParseJson(event)
      // if err != nil {
      //   log.Log(log.ERROR, fmt.Sprintf("Failed to parse event to JSON: %v", err))
      // } else {
      //   log.Log(log.DEBUG, s)
      // }

      if txEvent, ok := event.Data.(types.EventDataTx); ok {
        log.Debug("Received EventDataTx data")

        // 트랜잭션 데이터를 디코딩
        decodedTx, err := utils.DecodeTxData(txEvent.Tx)
        if err != nil {
          log.Error("Failed to decode transaction data: %v", err)
        } else {
          log.Info("\n---Decoded Transaction---\n%s\n", decodedTx)
        }

        err = utils.SaveTransactionToFile(txEvent, "./transactions")
        if err != nil {
          log.Error("Failed to save transaction to file: %v", err)
        } else {
          log.Info("Transaction saved to file successfully")
        }
      } else {
        log.Warn("Unknown event data type")
      }

    case <-ctx.Done():
      log.Debug("Event processing stopped due to timeout or cancellation")
      return
    }
  }
}
