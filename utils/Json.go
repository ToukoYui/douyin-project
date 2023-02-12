package utils

import "encoding/json"

type TokenInfo struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
}

func JsontoStruct(jsonStr string) TokenInfo {
	tokenInfo := TokenInfo{}
	json.Unmarshal([]byte(jsonStr), &tokenInfo)
	return tokenInfo
}
