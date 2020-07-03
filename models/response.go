package models

type Response struct {
	Status string `json:"status"`
	Error  bool   `json:"error"`
	Msg    string `json:"msg,omitempty"`
	Token  string `json:"token,omitempty"`
}
