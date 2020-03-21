package http_util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
	"time"
	"veryon/utils/errno_util"
)

const (
	userInfoHeaderKey = "username"
)

func GetUserName(c *gin.Context) (string, error) {
	var username = "unknown"
	if c == nil {
		return username, errno_util.NotFoundContext
	}
	if username = c.Request.Header.Get(userInfoHeaderKey); username != "" {
		return username, nil
	}
	if username = c.Query(userInfoHeaderKey); username != "" {
		return username, nil
	}
	return username, errno_util.NotFoundUserName
}

func GetPageInfo(c *gin.Context) (int, int) {
	rawPageIndex := c.DefaultQuery("page", "1")
	rawPageSize := c.DefaultQuery("limit", "20")
	pageIndex, err1 := strconv.Atoi(rawPageIndex)
	pageSize, err2 := strconv.Atoi(rawPageSize)
	if err1 != nil || err2 != nil {
		return 1, 20
	}
	return pageIndex, pageSize
}

// 服务后端向其他http服务发送请求时的http客户端对象结构体
type HttpClient struct {
	Request  *fasthttp.Request
	Response *fasthttp.Response
	Host     string
	Port     string
	Timeout  int
}

type ServerConfig struct {
	Host    string
	Port    string
	Timeout int
}

func Client(config *ServerConfig) *HttpClient {
	return &HttpClient{
		Request:  fasthttp.AcquireRequest(),
		Response: fasthttp.AcquireResponse(),
		Host:     config.Host,
		Port:     config.Port,
		Timeout:  config.Timeout,
	}
}

func (c *HttpClient) Close() {
	defer c.Request.ConnectionClose()
	defer c.Response.ConnectionClose()
}

func (c *HttpClient) SetRequestURI(uri string) {
	c.Request.SetRequestURI(uri)
}

func (c *HttpClient) requestResult() (*ResponseJson, error) {
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
		result := ResponseJson{}
		var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary
		if err = jsonIter.Unmarshal(c.Response.Body(), &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
}

func (c *HttpClient) request() ([]byte, error) {
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
		result := ResponseJson{}
		var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary
		if err = jsonIter.Unmarshal(c.Response.Body(), &result); err != nil {
			return nil, err
		}
		if result.Code >= http.StatusBadRequest {
			return nil, errors.New(fmt.Sprintf("bad request with message:%s, code:%d, data:%v", result.Message, result.Code, result.Data))
		}
		if result.Data == nil {
			return nil, errno_util.HttpNoResultData
		}
		if body, err := jsonIter.Marshal(result.Data); err != nil {
			return nil, err
		} else {
			return body, nil
		}
	}
}

func (c *HttpClient) requestWithMethodResult(method string, body []byte) (*ResponseJson, error) {
	c.Request.Header.SetMethod(method)
	if body != nil {
		c.Request.SetBody(body)
	}
	return c.requestResult()
}

func (c *HttpClient) requestWithMethod(method string, body []byte) ([]byte, error) {
	c.Request.Header.SetMethod(method)
	if body != nil {
		c.Request.SetBody(body)
	}
	return c.request()
}

func (c *HttpClient) GetResult() (*ResponseJson, error) {
	return c.requestWithMethodResult("GET", nil)
}

func (c *HttpClient) PostResult(body []byte) (*ResponseJson, error) {
	return c.requestWithMethodResult("GET", body)
}

func (c *HttpClient) PutResult(body []byte) (*ResponseJson, error) {
	return c.requestWithMethodResult("PUT", body)
}

func (c *HttpClient) Get() ([]byte, error) {
	return c.requestWithMethod("GET", nil)
}

func (c *HttpClient) Post(body []byte) ([]byte, error) {
	return c.requestWithMethod("POST", body)
}

func (c *HttpClient) Put(body []byte) ([]byte, error) {
	return c.requestWithMethod("PUT", body)
}
