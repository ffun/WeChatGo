package wx

import (
	"encoding/xml"
	"time"
)

type wxRequest struct {
	XMLName      string `xml:"xml"`
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        string `xml:"MsgId"`
}

type CDATAText struct {
	Text string `xml:",innerxml"`
}

type wxResponse struct {
	XMLName      string        `xml:"xml"`
	ToUserName   CDATAText     `xml:"ToUserName"`
	FromUserName CDATAText     `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      CDATAText     `xml:"MsgType"`
	Content      CDATAText     `xml:"Content"`
}

func v2CDATAText(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

func (req *wxRequest) Parse(data []byte) error {
	return xml.Unmarshal(data, req)
}

func NewResponse(fromUserName, toUserName, content, msgType string) ([]byte, error) {
	resp := wxResponse{}
	resp.FromUserName = v2CDATAText(fromUserName)
	resp.ToUserName = v2CDATAText(toUserName)
	resp.Content = v2CDATAText(content)
	resp.MsgType = v2CDATAText(msgType)
	resp.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(&resp, " ", " ")
}
