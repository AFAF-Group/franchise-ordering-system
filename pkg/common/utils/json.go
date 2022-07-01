package utils

import jsoniter "github.com/json-iterator/go"

var (
	json          = jsoniter.ConfigCompatibleWithStandardLibrary
	JSONMarshal   = json.Marshal
	JSONUnmarshal = json.Unmarshal
)

func ToJSON(data interface{}) ([]byte, error) {
	return JSONMarshal(data)
}

func FromJSON(json []byte, result interface{}) error {
	return JSONUnmarshal(json, result)
}

func MarshalJSONToString(data interface{}) (string, error) {
	return json.MarshalToString(data)
}
