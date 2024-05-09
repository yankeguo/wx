package wxreceive

import "encoding/xml"

type Message struct {
	XMLName xml.Name `xml:"xml"`

	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`

	// Event
	Event     string  `xml:"Event"`
	EventKey  string  `xml:"EventKey"`
	Ticket    string  `xml:"Ticket"`
	Latitude  float64 `xml:"Latitude"`
	Longitude float64 `xml:"Longitude"`
	Precision float64 `xml:"Precision"`

	// Text
	Content   string `xml:"Content"`
	MsgID     int64  `xml:"MsgId"`
	MsgDataId int64  `xml:"MsgDataId"`
	Idx       int    `xml:"Idx"`
}

func ParseMessage(buf []byte) (m Message, err error) {
	err = xml.Unmarshal(buf, &m)
	return
}
