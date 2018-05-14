package tag

import (
	util "github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type SetTagsRequest struct {
	TagList []string `json:"tag_list"` //用户需要设置的标签列表
	Cid     string   `json:"cid"`      //指定需要设置标签的用户id
}

type SetTagsResult struct {
	Result string `json:"result"`
	Desc   string `json:"desc"`
}

func SetTags(conf entity.IAppConfig, req *SetTagsRequest) (*SetTagsResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/set_tags"

	reply := new(SetTagsResult)
	err := util.Post(url, conf.GetToken(), req, reply)

	return reply, err
}
