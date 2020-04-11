package liveqing_util

import "gitea.com/huangxuantao89/utils/http_util"

type Config struct {
	http_util.ServerConfig
	MaxChannels int
	Username    string
	Password    string
}

type LiveQingHeader struct {
	CSeq        string `json:"CSeq"`
	Version     string `json:"Version"`
	MessageType string `json:"MessageType"`
	ErrorNum    string `json:"ErrorNum"`
	ErrorString string `json:"ErrorString"`
}

type LoginRespVO struct {
	LiveQing struct {
		Header LiveQingHeader `json:"Header"`
		Body   struct {
			Token        string `json:"Token"`
			TokenTimeout int    `json:"TokenTimeout"`
		} `json:"Body"`
	} `json:"LiveQing"`
}

type GetChannelsConfigRespVO struct {
	LiveQing struct {
		Header LiveQingHeader `json:"Header"`
		Body   struct {
			ChannelCount int `json:"ChannelCount"`
			Channels     []struct {
				Enable     int    `json:"Enable"`
				Audio      int    `json:"Audio"` // 开启音频 0-关 1-开
				CDN        string `json:"Cdn"`
				Channel    int    `json:"Channel"`
				GBID       string `json:"GBID"` // 国标编号
				IP         string `json:"IP"`
				Name       string `json:"Name"`
				OnDemand   int    `json:"OnDemand"` // 按需直播 0-关 1-开
				Onvif      string `json:"Onvif"`
				Username   string `json:"UserName"`
				Password   string `json:"Password"`
				Port       int    `json:"Port"`
				Protocol   string `json:"Protocol"` // 接入协议
				Record     int    `json:"Record"`   // 开启录像
				RecordPlan string `json:"RecordPlan"`
				RTSP       string `json:"Rtsp"`
				Transport  string `json:"Transport"` // 传输协议
			} `json:"Channels"`
		} `json:"Body"`
	} `json:"LiveQing"`
}

type SetChannelConfigReqVO struct {
	Channel   int    `json:"Channel"`
	Enable    int    `json:"Enable"`
	OnDemand  int    `json:"OnDemand"`
	Name      string `json:"Name"`
	Protocol  string `json:"Protocol"`
	RTSP      string `json:"Rtsp"`
	Audio     int    `json:"Audio"`
	Record    int    `json:"Record"`
	Transport string `json:"Transport"`
	CDN       string `json:"Cdn"`
}

type SetChannelConfigRespVO struct {
	LiveQing struct {
		Header LiveQingHeader `json:"Header"`
	} `json:"LiveQing"`
}
