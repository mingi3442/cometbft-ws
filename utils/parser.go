package utils

import (
  "encoding/json"
  "fmt"
  "log"

  coretypes "github.com/cometbft/cometbft/rpc/core/types"
)

// 이벤트 데이터를 JSON으로 변환
func ParseJson(event coretypes.ResultEvent) (string, error) {

  jsonData, err := json.MarshalIndent(event, "", "  ")
  if err != nil {
    log.Printf("Failed to marshal event to JSON: %v", err)
    return "", err
  }

  fmt.Println("New IBC Event Received:")
  fmt.Println(string(jsonData))

  return string(jsonData), nil
}
