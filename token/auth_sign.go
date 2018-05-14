package token

import (
	"github.com/sipt/go-getui-api/util"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Param struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	AppKey    string `json:"appkey"`
}

type TokenResult struct {
	Result    string `json:"result"`
	AuthToken string `json:"auth_token"`
}

func GetAuthSign(appId string, appKey string, masterSecret string) (*TokenResult, error) {

	param, err := GetPostBody(appKey, masterSecret)
	if err != nil {
		return nil, err
	}

	url := util.TOKEN_DOMAIN + appId + "/auth_sign"
	reply := new(TokenResult)
	err = util.Post(url, "", param, &reply)

	return reply, err
}

func GetPostBody(appKey string, masterSecret string) (*Param, error) {

	signStr, timestamp := HmacSha256(appKey, masterSecret)

	return &Param{
		Sign:      signStr,
		Timestamp: timestamp,
		AppKey:    appKey,
	}, nil
}

func HmacSha256(appKey string, masterSecret string) (string, string) {
	timestamp := strconv.FormatInt((time.Now().UnixNano() / 1000000), 10) //签名开始生成毫秒时间
	original := appKey + timestamp + masterSecret

	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum), timestamp
}
