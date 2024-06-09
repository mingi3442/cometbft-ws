package main

import (
  "context"
  "fmt"
  log "github.com/mingi3442/go-grpc/log"

  ws "github.com/mingi3442/tendermint-ws/client"
  utils "github.com/mingi3442/tendermint-ws/utils"
  "time"
)

func main() {
  rpcUrl := "https://cosmos-rpc.polkachu.com"

  wsClient, err := ws.Connect(rpcUrl)
  if err != nil {
    msg := fmt.Sprintf("Failed to connect to RPC server: %v", err)
    log.Log(log.ERROR, msg)
  }

  defer wsClient.DisConnect()

  // query := "tm.event='Tx' AND (message.action='send_packet' OR message.action='recv_packet' OR message.action='acknowledge_packet' OR message.action='timeout_packet')"
  query := "tm.event='Tx'"
  ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
  defer cancel()

  subscriber := "relayer"
  events, err := wsClient.RpcClient.Subscribe(ctx, subscriber, query)
  if err != nil {
    msg := fmt.Sprintf("Failed to subscribe to events: %v", err)
    log.Log(log.ERROR, msg)
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
