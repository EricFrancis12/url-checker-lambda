package main

type LambdaResp struct {
	TargetHostname string `json:"targetHostname"`
}

func NewLambdaResp(thn string) LambdaResp {
	return LambdaResp{
		TargetHostname: thn,
	}
}
