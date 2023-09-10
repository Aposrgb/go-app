package response

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ToJsonBytes(object interface{}) ([]byte, error) {
	json, err := ToJson(object)
	if err != nil {
		return []byte{}, err
	}
	return []byte(json), nil
}

func ToJson(object interface{}) (string, error) {
	mapData := make(map[string]interface{})
	mapData["data"] = object
	json, err := json.Marshal(mapData)
	if err != nil {
		fmt.Println("Object not serialaziable")
		return "", err
	}
	return string(json), nil
}

func getHeadersJson() map[string]string {
	return map[string]string{
		"Content-Type":"application/json",
	}
}

func SetHeadersJson(c *gin.Context) {
	headers := getHeadersJson()
	for key, item := range headers {
		c.Header(key, item)
	}
}