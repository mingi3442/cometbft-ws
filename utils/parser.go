package utils

import (
  "encoding/json"
  "fmt"
  coretypes "github.com/cometbft/cometbft/rpc/core/types"
  log "github.com/mingi3442/go-grpc/log"
)

func ParseJson(event coretypes.ResultEvent) (string, error) {

  jsonData, err := json.MarshalIndent(event, "", "  ")
  if err != nil {
    msg := fmt.Sprintf("Failed to marshal event to JSON: %v", err)
    log.Log(log.ERROR, msg)
    return "", err
  }

  fmt.Println("New IBC Event Received:")
  fmt.Println(string(jsonData))

  return string(jsonData), nil
}
