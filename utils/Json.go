package utils

import "encoding/json"

type TokenInfo struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
}

func JsonToStruct(jsonStr string) TokenInfo {
	tokenInfo := TokenInfo{}
	json.Unmarshal([]byte(jsonStr), &tokenInfo)
	return tokenInfo
}

func StructToJson(value TokenInfo) string {
	jsonString, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(jsonString)
}
