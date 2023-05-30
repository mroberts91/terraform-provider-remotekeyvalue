package remotekeyvalue

import "encoding/json"

func UnmarshalApiResponseItem(data []byte) (ApiResponseItem, error) {
	var r ApiResponseItem
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalApiResponseArray(data []byte) ([]ApiResponseItem, error) {
	var r []ApiResponseItem
	err := json.Unmarshal(data, &r)
	return r, err
}

func (a ApiResponseItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":        a.ID,
		"key":       a.Key,
		"value":     a.Value,
		"sensitive": a.Value,
	}
}

type ApiResponseItem struct {
	ID          int64  `json:"id" mapstructure:"id"`
	Key         string `json:"key" mapstructure:"key"`
	Value       string `json:"value" mapstructure:"value"`
	IsSensitive bool   `json:"isSensitive" mapstructure:"sensitive"`
}
