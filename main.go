package main

import (
  "context"
  "fmt"
  "log"

  rpchttp "github.com/cometbft/cometbft/rpc/client/http"
  utils "github.com/mingi3442/tendermint-ws/utils/event_handler"
  "time"
)

func main() {
  url := "https://cosmos-rpc.polkachu.com"
  rpcClient, err := rpchttp.New(url, "/websocket")
  if err != nil {
    log.Fatalf("Failed to create RPC client: %v", err)
  }

  if err := rpcClient.Start(); err != nil {
    log.Fatalf("Failed to start RPC client: %v", err)
  }
  defer rpcClient.Stop()

  // query := "tm.event='Tx' AND (message.action='send_packet' OR message.action='recv_packet' OR message.action='acknowledge_packet' OR message.action='timeout_packet')"
  query := "tm.event='Tx'"
  ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
  defer cancel()

  subscriber := "my-relayer"
  events, err := rpcClient.Subscribe(ctx, subscriber, query)
  if err != nil {
    log.Fatalf("Failed to subscribe to events: %v", err)
  }

  fmt.Println("Subscribed to IBC events...")

  for {
    select {
    case event := <-events:
      utils.ParseJson(event)
    case <-ctx.Done():
      fmt.Println("Subscription timed out")
      return
    }
  }
}
