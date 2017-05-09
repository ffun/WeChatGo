package robot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var _ IRobot = &QinyunRobot{} //确保QinyunRobot实现了IRobot接口

type QinyunRobot struct {
}

const (
	Msg_key   = "msg"
	Robot_URL = "http://api.qingyunke.com/api.php?key=free&appid=0&"
)

//Robt 调用rpc
func (robt *QinyunRobot) Call(method, params string) ([]byte, error) {
	//对查询字符串进行URL编码
	query := url.Values{}
	query.Add(Msg_key, params)
	queryURL := Robot_URL + query.Encode()
	method = strings.ToUpper(method)
	switch method {
	case "GET":
		//Get Api请求聊天机器人
		resp, err := http.Get(queryURL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		return body, nil
	case "POST":
		fallthrough
	default:
		return nil, errors.New("has no this method")
	}

}

type QinyunResp struct {
	Result  int    `json:"result"`
	Content string `json:"content"`
}

//robot对返回结果的处理
func (robt *QinyunRobot) Response(w io.Writer, resp []byte) error {
	result := QinyunResp{}
	if err := json.Unmarshal(resp, &result); err != nil {
		return err
	}
	fmt.Fprintln(w, result.Content)
	return nil
}

func (robot *QinyunRobot) Name() string {
	return "qinyun"
}
