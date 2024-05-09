package wxreceive

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResponseMarshal(t *testing.T) {
	r := NewResponse()
	r.SetFromUserName("aaa")
	r.SetToUserName("bbb")
	r.SetCreateTimeInt64(123)
	r.SetMsgType("text")

	buf, err := r.Marshal()
	require.NoError(t, err)
	require.Equal(t, "<xml><ToUserName><![CDATA[bbb]]></ToUserName><FromUserName><![CDATA[aaa]]></FromUserName><CreateTime>123</CreateTime><MsgType><![CDATA[text]]></MsgType></xml>", string(buf))
}
