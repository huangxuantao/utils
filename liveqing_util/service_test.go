package liveqing_util

import (
	"net/url"
	"testing"
)

func Test_1(t *testing.T) {
	t.Log(url.QueryEscape("rtsp://admin:Xhh20202020@192.168.0.202:554/cam/realmonitor?channel=1&subtype=1"))
}
