package utils

import (
  "encoding/json"
  "fmt"
  coretypes "github.com/cometbft/cometbft/rpc/core/types"

  reflect "reflect"

  log "github.com/mingi3442/go-grpc/log"
)

func ParseJson(event coretypes.ResultEvent) (string, error) {
  actions, found := event.Events["message.action"]

  if !found {
    log.Log(log.WARN, "No message.action.")
    return "", nil
  }

  log.Log(log.DEBUG, "Extracted message.action values")

  fmt.Printf("Type of event.Data: %s\n", reflect.TypeOf(event.Data))
  for _, action := range actions {
    msg := fmt.Sprintf(" - %s", action)
    log.Log(log.DEBUG, msg)
  }

  jsonData, err := json.MarshalIndent(event.Data, "", "  ")
  if err != nil {
    msg := fmt.Sprintf("Failed to marshal event to JSON: %v", err)
    log.Log(log.ERROR, msg)
    return "", err
  }

  return string(jsonData), nil
}
