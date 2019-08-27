package exchange

import "fmt"

type ErrorCode int32

type ExchangeError interface {
	error
	Code() ErrorCode
}

func (c ErrorCode) String() string {
	return errorCodeMap[c]
}

func (c ErrorCode) Error() string {
	return fmt.Sprintf("code:%d,desc:%s", c, c.String())
}

func (c ErrorCode) Code() ErrorCode {
	return c
}

var errorCodeMap = make(map[ErrorCode]string)

func MergeErrorCodeMap(codeMap map[ErrorCode]string) {
	for errorCode, errorDesc := range codeMap {
		errorCodeMap[errorCode] = errorDesc
	}
}

const (
	ErrorCodeExchange ErrorCode = 50000 * iota
)
