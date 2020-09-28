package http_util

import (
	"fmt"
	"github.com/huangxuantao/utils/errno_util"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

func (c *HttpClient) requestWithoutFormat() ([]byte, error) {
	if c.Port == "" {
		c.Request.SetHost(c.Host)
	} else {
		c.Request.SetHost(c.Host + ":" + c.Port)
	}
	c.Request.Header.SetContentType("application/json")
	if err := fasthttp.DoTimeout(c.Request, c.Response, time.Duration(c.Timeout)*time.Second); err != nil {
		return nil, err
	} else {
		if len(c.Response.Body()) == 0 {
			return nil, errno_util.HttpNoResponseContentReturn
		}
		if c.Response.StatusCode() != http.StatusOK {
			return nil, errors.New(fmt.Sprintf("response status code:%d", c.Response.StatusCode()))
		}
		return c.Response.Body(), nil
	}
}


func (c *HttpClient) requestWithMethodWithoutFormat(method string, body []byte) ([]byte, error) {
	c.Request.Header.SetMethod(method)
	if body != nil {
		c.Request.SetBody(body)
	}
	return c.requestWithoutFormat()
}

func (c *HttpClient) GetWithoutFormat() ([]byte, error) {
	return c.requestWithMethodWithoutFormat("GET", nil)
}

func (c *HttpClient) PostWithoutFormat(body []byte) ([]byte, error) {
	return c.requestWithMethodWithoutFormat("POST", body)
}

func (c *HttpClient) PutWithoutFormat(body []byte) ([]byte, error) {
	return c.requestWithMethodWithoutFormat("PUT", body)
}
