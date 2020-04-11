package liveqing_util

import (
	"encoding/json"
	"fmt"
	"gitea.com/huangxuantao89/utils/http_util"
)

const (
	ChannelEnable        = 1 // 通道开
	ChannelDisable       = 0 // 通道关
	ChannelOnDemandOn    = 1 // 按需直播开
	ChannelOnDemandOff   = 0 // 按需直播关
	ChannelProtocolRTSP  = "RTSP"
	ChannelProtocolONVIF = "ONVIF"
	ChannelAudioOn       = 1
	ChannelAudioOff      = 0
	ChannelRecordOff     = 0
	ChannelTransportTCP  = "TCP"

	uriLogin             = "/api/v1/login"
	uriGetChannelsConfig = "/api/v1/getchannelsconfig"
	uriSetChannelConfig  = "/api/v1/setchannelconfig"
)

type liveQingService struct {
	Conf Config
}

func GetLiveQingService() *liveQingService {
	return &liveQingService{}
}

func (d *liveQingService) Login() (string, error) {
	client := http_util.Client(&d.Conf.ServerConfig)
	defer client.Close()
	client.SetRequestURI(fmt.Sprintf("%s?username=%s&password=%s", uriLogin, d.Conf.Username, d.Conf.Password))

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
	client := http_util.Client(&d.Conf.ServerConfig)
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
	client := http_util.Client(&d.Conf.ServerConfig)
	defer client.Close()
	client.Request.Header.SetCookie("token", token)
	client.SetRequestURI(fmt.Sprintf("%s?Channel=%d&Enable=%d&OnDemand=%d&Name=%s&Rtsp=%s&Audio=%d&Record=%d&Transport=%s&Protocol=%s&Cdn=%s",
		uriSetChannelConfig,
		reqVO.Channel,
		reqVO.Enable,
		reqVO.OnDemand,
		reqVO.Name,
		reqVO.RTSP,
		reqVO.Audio,
		reqVO.Record,
		reqVO.Transport,
		reqVO.Protocol,
		reqVO.CDN))
	
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
