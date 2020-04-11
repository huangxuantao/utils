package liveqing_util

import (
	"encoding/json"
	"fmt"
	"gitea.com/huangxuantao89/utils/http_util"
)

const (
	uriLogin             = "/api/v1/login"
	uriGetChannelsConfig = "/api/v1/getchannelsconfig"
	uriSetChannelConfig  = "/api/v1/setchannelconfig"
)

type liveQingService struct {
	Conf http_util.ServerConfig
}

func GetLiveQingService() *liveQingService {
	return &liveQingService{}
}

func (d *liveQingService) Login(reqVO LoginReqVO) (string, error) {
	client := http_util.Client(&d.Conf)
	defer client.Close()
	client.SetRequestURI(fmt.Sprintf("%s?username=%s&password=%s", uriLogin, reqVO.Username, reqVO.Password))

	body, err := client.GetWithoutFormat()
	if err != nil {
		return "", err
	}

	var respVO LoginRespVO
	if err := json.Unmarshal(body, &respVO); err != nil {
		return "", err
	}

	return respVO.LiveQing.Body.Token, nil
}

func (d *liveQingService) GetChannelsConfig(token string) (*GetChannelsConfigRespVO, error) {
	client := http_util.Client(&d.Conf)
	defer client.Close()
	client.Request.Header.SetCookie("token", token)
	client.SetRequestURI(uriGetChannelsConfig)

	body, err := client.GetWithoutFormat()
	if err != nil {
		return nil, err
	}

	var respVO GetChannelsConfigRespVO
	if err := json.Unmarshal(body, &respVO); err != nil {
		return nil, err
	}
	return &respVO, nil
}

func (d *liveQingService) SetChannelConfig(reqVO SetChannelConfigReqVO, token string) error {
	client := http_util.Client(&d.Conf)
	defer client.Close()
	client.Request.Header.SetCookie("token", token)
	client.SetRequestURI(fmt.Sprintf("%s?Channel=%d&Enable=%d&OnDemand=%d&Name=%s&Rtsp=%s&Audio=%d&Record=%d&Transport=%s&Protocol=%s",
		uriSetChannelConfig, reqVO.Channel, reqVO.Enable, reqVO.OnDemand, reqVO.Name, reqVO.RTSP, reqVO.Audio, reqVO.Record, reqVO.Transport, reqVO.Protocol))

	body, err := client.GetWithoutFormat()
	if err != nil {
		return err
	}

	var respVO SetChannelConfigRespVO
	if err := json.Unmarshal(body, &respVO); err != nil {
		return err
	}

	return nil
}
