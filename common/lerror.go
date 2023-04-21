package common

type LError struct {
	HttpCode int    `json:"status"`
	Msg      string `json:"msg"`
	Err      error  `json:"error"`
}
