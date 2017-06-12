package lxd

import "cointhink/proto"

type AccountOperation struct {
	Account   proto.Account
	Operation OperationResponse
}

type OperationResponse struct {
	Type       string `json:"type"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Operation  string `json:"operation"`
	ErrorCode  int    `json:"error_code"`
	Error      string `json:"error"`
	Metadata   struct {
		ID         string `json:"id"`
		Class      string `json:"class"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
		Status     string `json:"status"`
		StatusCode int    `json:"status_code"`
		Resources  struct {
			Containers []string `json:"containers"`
		} `json:"resources"`
		Metadata  interface{} `json:"metadata"`
		MayCancel bool        `json:"may_cancel"`
		Err       string      `json:"err"`
	} `json:"metadata"`
}
