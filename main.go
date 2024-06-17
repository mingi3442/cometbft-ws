package main

import (
  "context"

  log "github.com/mingi3442/go-grpc/log"

  ws "github.com/mingi3442/tendermint-ws/client"
  "github.com/mingi3442/tendermint-ws/event"

  "time"
)

func main() {
  rpcUrl := "https://cosmos-rpc.polkachu.com"
  subscriber := "relayer"
  query := "tm.event='Tx'"
  // query := "tm.event='Tx' AND (message.module='ibc_channel' OR message.module='ibc_transfer' OR message.module='ibc_client' OR EXISTS(ibc_channel.*) OR EXISTS(ibc_transfer.*) OR EXISTS(ibc_client.*) OR EXISTS(send_packet.*) OR EXISTS(recv_packet.*) OR EXISTS(acknowledge_packet.*) OR EXISTS(timeout_packet.*))"

  wsClient, err := ws.Connect(rpcUrl)
  if err != nil {
    log.Error("Failed to connect to RPC server: %v", err)
  }

  defer wsClient.DisConnect()

  ctx, cancel := context.WithTimeout(context.Background(), time.Minute*100)
  defer cancel()

  events, err := wsClient.Subscribe(ctx, subscriber, query)
  if err != nil {
    log.Error("Failed to subscribe to events: %v", err)
  }

  go event.HandleEvents(ctx, events)

  <-ctx.Done()
  log.Warn("Main loop timed out")
}
