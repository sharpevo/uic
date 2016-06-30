package models

const (
	ERROR_DATABASE  = 1
	ERROR_DUPLICATE = 2
	ERROR_NOT_FOUND = 3
	ERROR_INPUT     = 4
)

type CodeInfo struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

func ErrorJSON(info string) *CodeInfo {
	return &CodeInfo{-1, info}
}

func NormalJSON(info string) *CodeInfo {
	return &CodeInfo{0, info}
}
