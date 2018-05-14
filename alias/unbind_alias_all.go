package alias

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type UnbindAliasAllRequest struct {
	Alias string `json:"alias"`
}
type UnbindAliasAllResult struct {
	Result string `json:"result"`
	Desc   string `json:"desc"`
}

func UnbindAliasAll(conf entity.IAppConfig, req *UnbindAliasAllRequest) (*UnbindAliasAllResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/unbind_alias_all"

	reply := new(UnbindAliasAllResult)
	err := util.Post(url, conf.GetToken(), req, reply)

	return reply, err
}
