package token

import (
	"github.com/sipt/go-getui-api/util"
)

type CloseTokenResult struct {
	Result string `json:"result"`
}

func SetAuthClose(appID, authToken string) (*CloseTokenResult, error) {

	url := util.TOKEN_DOMAIN + appID + "/auth_close"

	reply := new(CloseTokenResult)
	err := util.Post(url, authToken, nil, reply)

	return reply, err
}
