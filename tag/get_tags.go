package tag

import (
	util "github.com/sipt/go-getui-api/util"
	"git.yuapt.com/sipt/security/pkg/dep/sources/https---github.com-nats--io-gnatsd/conf"
	"github.com/sipt/go-getui-api/entity"
)

type GetTagsResult struct {
	Result string `json:"result"`
	Desc   string `json:"desc"`
}

func GetTags(conf entity.AppConfig, cid string) (*GetTagsResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/get_tags/" + cid

	reply := new(GetTagsResult)
	err := util.Get(url, conf.GetToken(), reply)

	return reply, err
}
