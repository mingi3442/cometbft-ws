package client

import (
  "fmt"

  rpchttp "github.com/cometbft/cometbft/rpc/client/http"
  "github.com/mingi3442/go-grpc/log"
)

type WsClient struct {
  RpcClient *rpchttp.HTTP
}

func Connect(url string) (*WsClient, error) {

  rpcWsClient, err := rpchttp.New(url, "/websocket")
  if err != nil {
    return nil, err
  }

  if err := rpcWsClient.Start(); err != nil {
    msg := fmt.Sprintf("Failed to start RPC client: %v", err)
    log.Log(log.ERROR, msg)
    return nil, err
  }
  return &WsClient{
    RpcClient: rpcWsClient,
  }, nil
}

func (c *WsClient) DisConnect() error {
  if c.RpcClient != nil {
    return c.RpcClient.Stop()
  }
  return nil
}
