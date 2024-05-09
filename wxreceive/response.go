package wxreceive

import "encoding/xml"

const ResponseSuccess = "success"

type CDATA struct {
	Value string `xml:",cdata"`
}

type Response struct {
	XMLName xml.Name `xml:"xml"`

	ToUserName   CDATA  `xml:"ToUserName"`
	FromUserName CDATA  `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      CDATA  `xml:"MsgType"`
	Content      *CDATA `xml:"Content,omitempty"`
}

func (r *Response) SetToUserName(v string) *Response {
	r.ToUserName.Value = v
	return r
}

func (r *Response) SetFromUserName(v string) *Response {
	r.FromUserName.Value = v
	return r
}

func (r *Response) SetCreateTimeInt64(v int64) *Response {
	r.CreateTime = v
	return r
}

func (r *Response) SetMsgType(v string) *Response {
	r.MsgType.Value = v
	return r
}

func (r *Response) SetContent(v string) *Response {
	if r.Content == nil {
		r.Content = &CDATA{}
	}
	r.Content.Value = v
	return r
}

func (r *Response) Marshal() (buf []byte, err error) {
	buf, err = xml.Marshal(r)
	return
}

func NewResponse() (res *Response) {
	return &Response{}
}

func NewResponseFromMessage(m *Message) (res *Response) {
	res = NewResponse()
	res.SetToUserName(m.FromUserName)
	res.SetFromUserName(m.ToUserName)
	res.SetCreateTimeInt64(m.CreateTime)
	return
}
