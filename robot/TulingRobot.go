package robot

import "io"
import "strings"

var _ IRobot = &TulingRobot{} //确保QinyunRobot实现了IRobot接口

type TulingRobot struct {
}

type TulingResquest struct {
	Key    string `json:"key"`
	Info   string `json:"info"`
	UserID string `json:"userid"`
}

type TulingResp struct {
	Code int    `json:"code"`
	Text string `json:"text"`
	Url  string `json:"url"`
	
}

func (robot *TulingRobot) Call(method, params string) ([]byte, error) {
	method = strings.ToUpper(method)

	return nil, nil
}

func (robot *TulingRobot) Response(w io.Writer, resp []byte) error {
	return nil
}

func (robot *TulingRobot) Name() string {
	return "tuling"
}
