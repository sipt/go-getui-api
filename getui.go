package main

import (
	"github.com/sipt/go-getui-api/push"
	"github.com/sipt/go-getui-api/query"
	"github.com/sipt/go-getui-api/style"
	"github.com/sipt/go-getui-api/token"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
	"github.com/sipt/go-getui-api/entity"
)

var (
	appId        string = "XH93kDE2AZ6x3pCGwEQNn"
	appKey       string = "mL0IIpwukX53MGE4BjZjs1"
	appSecret    string = "KUp3G7LC6V98fZsUdeTGO5"
	masterSecret string = "tT1khrhlup8vskHi5iVpk4"
	cid          string = "45f8a382f93b018a4ba4b5cb6c497cc0"
)

func main() {

	conf := entity.NewAppConfig(appId, appSecret, appKey, masterSecret)

	saveListBodyParam := GetSaveListBodyParam(appKey)
	saveRes, err := SaveListBody(conf, saveListBodyParam)
	if err != nil {
		log.Error(fmt.Sprintf("save list body  err : ") + err.Error())
		return
	}

	Param := GetPushListParam(saveRes.TaskId, []string{cid})
	_, err = push.PushList(conf, Param)
	if err != nil {
		log.Error(fmt.Sprintf("save list body  err : ") + err.Error())
		return
	}

	pushSingleResult, err := pushSingle(conf)
	if err != nil {
		log.Error(fmt.Sprintf("get push single err : ", err))
		return
	}

	_, err = getPushResult(conf, pushSingleResult.TaskId)
	if err != nil {
		log.Error(fmt.Sprintf("query push result err : ", err))
		return
	}

}

func GetPushListParam(taskId string, cids []string) *push.PushListParam {

	pushListParam := &push.PushListParam{
		TaskId:     taskId,
		Cid:        cids,
		NeedDetail: true,
	}

	return pushListParam
}

func SaveListBody(conf entity.IAppConfig, Param *push.SaveListBodyParam) (*push.SaveListBodyResult, error) {

	saveListBodyResult, err := push.SaveListBody(conf, Param)
	if err != nil {
		log.Error(fmt.Sprintf("get push single err : ", err))
		return saveListBodyResult, err
	}
	log.Info("saveListBodyResult", log.Ctx{
		"saveListBodyResult": saveListBodyResult,
	})
	return saveListBodyResult, err
}

func GetSaveListBodyParam(appKey string) *push.SaveListBodyParam {

	message := entity.GetMessage()
	message.SetAppKey(appKey)
	message.SetMsgType("notification")

	notification := entity.GetNotification()
	notification.SetTransmissionContent("透传内容")

	unWindStyle := style.GetUnwindStyle("检测到可疑人员", "警告通知")
	unWindStyle.SetBigStyle("1")
	unWindStyle.SetBigImageUrl("http://s0.hao123img.com/res/r/image/2016-04-14/2a3b604cdc47bdc4e2ffa252d31179d1.jpg")

	notification.SetNotifyStyle(unWindStyle)

	saveListBodyParam := &push.SaveListBodyParam{
		Message:      message,
		Notification: notification,
		TaskName:     time.Now().Format("20160102150405"),
	}
	log.Info("saveListBodyParam", log.Ctx{
		"saveListBodyParam": saveListBodyParam,
	})
	return saveListBodyParam
}

func getPushResult(conf entity.IAppConfig, taskId string) (*query.PushRESResult, error) {
	pushRESParam := &query.PushRESParam{
		TaskIdList: []string{taskId},
	}

	PushRESResult, err := query.PushResult(conf, pushRESParam)
	if err != nil {
		log.Error(fmt.Sprintf("query push result err : ", err))
		return PushRESResult, err
	}
	return PushRESResult, nil
}

//单推
func pushSingle(conf entity.IAppConfig) (*push.PushSingleResult, error) {

	message := entity.GetMessage()
	message.SetAppKey(appKey)
	message.SetMsgType("notification")

	notification := entity.GetNotification()
	notification.SetTransmissionContent("透传内容")

	unWindStyle := style.GetUnwindStyle("检测到可疑人员", "警告通知")
	unWindStyle.SetBigStyle("1")
	unWindStyle.SetBigImageUrl("http://s0.hao123img.com/res/r/image/2016-04-14/2a3b604cdc47bdc4e2ffa252d31179d1.jpg")

	notification.SetNotifyStyle(unWindStyle)

	pushSingleParam := &push.PushSingleParam{
		Message:      message,
		Notification: notification,
		Cid:          cid,
		RequestId:    time.Now().Format("20160102150405"),
	}
	log.Info("pushSingleParam", log.Ctx{
		"pushSingleParam": pushSingleParam,
	})

	pushSingleResult, err := push.PushSingle(conf, pushSingleParam)
	if err != nil {
		log.Error(fmt.Sprintf("get push single err : ", err))
		return pushSingleResult, err
	}
	log.Info("push single", log.Ctx{
		"result": pushSingleResult.Result,
		"status": pushSingleResult.Status,
		"taskId": pushSingleResult.TaskId,
	})

	return pushSingleResult, nil
}

func getGeTuiToken() (string, error) {
	tokenResult, err := token.GetAuthSign(appId, appKey, masterSecret)
	if err != nil {
		log.Error(fmt.Sprintf("get getui sign token err : ", err))
		return "", err
	}
	log.Info("token", log.Ctx{
		"tokenResult": tokenResult,
	})
	return tokenResult.AuthToken, nil
}
