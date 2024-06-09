package main

import (
  "context"
  "fmt"

  log "github.com/mingi3442/go-grpc/log"

  ws "github.com/mingi3442/tendermint-ws/client"
  "github.com/mingi3442/tendermint-ws/event"

  "time"
)

func main() {
  rpcUrl := "https://cosmos-rpc.polkachu.com"
  subscriber := "relayer"
  query := "tm.event='Tx'"
  // query := "tm.event='Tx' AND (message.action='send_packet' OR message.action='recv_packet' OR message.action='acknowledge_packet' OR message.action='timeout_packet')"

  wsClient, err := ws.Connect(rpcUrl)
  if err != nil {
    msg := fmt.Sprintf("Failed to connect to RPC server: %v", err)
    log.Log(log.ERROR, msg)
  }

  defer wsClient.DisConnect()

  ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
  defer cancel()

  events, err := wsClient.Subscribe(ctx, subscriber, query)
  if err != nil {
    msg := fmt.Sprintf("Failed to subscribe to events: %v", err)
    log.Log(log.ERROR, msg)
  }

  go event.HandleEvents(ctx, events)

  select {
  case <-ctx.Done():
    fmt.Println("Main loop timed out")
  }
}
