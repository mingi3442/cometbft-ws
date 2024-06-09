package client

import (
  "log"

  rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

type WsClient struct {
  conn *rpchttp.HTTP
}

func Connect(url string) (*WsClient, error) {

  rpcWsClient, err := rpchttp.New(url, "/websocket")
  if err != nil {
    return nil, err
  }

  if err := rpcWsClient.Start(); err != nil {
    log.Fatalf("Failed to start RPC client: %v", err)
    return nil, err
  }
  return &WsClient{
    conn: rpcWsClient,
  }, nil
}

func (c *WsClient) DisConnect() error {
  if c.conn != nil {
    return c.conn.Stop()
  }
  return nil
}
