package models

type RequestBody struct {
	Key        string      `json:"key"`
	Value      interface{} `json:"value"`
	Expiration int         `json:"expiration"`
}
