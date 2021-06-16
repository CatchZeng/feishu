package feishu

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/CatchZeng/feishu/internal/security"
	"github.com/go-resty/resty/v2"
)

const feishuAPI = "https://open.feishu.cn/open-apis/bot/v2/hook/"

// Client feishu client
type Client struct {
	AccessToken string
	Secret      string
}

// NewClient new dingtalk client
func NewClient(accessToken, secret string) *Client {
	return &Client{
		AccessToken: accessToken,
		Secret:      secret,
	}
}

// Response response struct
type Response struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

// Send send message
func (d *Client) Send(message Message) (*Response, error) {
	res := &Response{}

	if len(d.AccessToken) < 1 {
		return res, fmt.Errorf("accesstoken is empty")
	}

	timestamp := time.Now().Unix()
	sign, err := security.GenSign(d.Secret, timestamp)
	if err != nil {
		return res, err
	}

	body := message.Body()
	body["timestamp"] = strconv.FormatInt(timestamp, 10)
	body["sign"] = sign

	log.Print(body)

	client := resty.New()
	URL := fmt.Sprintf("%v%v", feishuAPI, d.AccessToken)
	resp, err := client.SetRetryCount(3).R().
		SetBody(body).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetResult(&Response{}).
		ForceContentType("application/json").
		Post(URL)

	if err != nil {
		return nil, err
	}

	result := resp.Result().(*Response)
	if result.Code != 0 {
		return res, fmt.Errorf("send message to feishu error = %s", result.Msg)
	}
	return result, nil
}
