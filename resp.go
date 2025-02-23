package main

import "encoding/json"

type LambdaResp struct {
	Data  []byte `json:"data"`
	Error error  `json:"error"`
}

func NewLambdaResp(data []byte, err error) LambdaResp {
	return LambdaResp{
		Data:  data,
		Error: err,
	}
}

func NewErrResp(err error) LambdaResp {
	return NewLambdaResp([]byte{}, err)
}

func (lr LambdaResp) ToBytes() []byte {
	jsonData, err := json.Marshal(lr)
	if err != nil {
		return []byte(`{"data":"","error":"error marshalling to json"}`)
	}
	return jsonData
}
