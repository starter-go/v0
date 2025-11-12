package dxo

import "strings"

////////////////////////////////////////////////////////////////////////////////

// in sec
type WxMsgTime int64

////////////////////////////////////////////////////////////////////////////////

type WxMsgType string

const (
	WxMsgTypeText       WxMsgType = "text"
	WxMsgTypeImage      WxMsgType = "image"
	WxMsgTypeVoice      WxMsgType = "voice"
	WxMsgTypeVideo      WxMsgType = "video"
	WxMsgTypeShortVideo WxMsgType = "shortvideo"
	WxMsgTypeLocation   WxMsgType = "location"
	WxMsgTypeLink       WxMsgType = "link"
)

func (t WxMsgType) Normalize() WxMsgType {
	str := t.String()
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	return WxMsgType(str)
}

func (t WxMsgType) String() string {
	return string(t)
}

////////////////////////////////////////////////////////////////////////////////

type WxMsgUserName string

func (name WxMsgUserName) String() string {
	return string(name)
}

func (name WxMsgUserName) Normalize() WxMsgUserName {
	str := name.String()
	str = strings.TrimSpace(str)
	// str = strings.ToLower(str)
	return WxMsgUserName(str)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
