package dto

import (
	"encoding/xml"

	"github.com/starter-go/v0/weixin-mp-webapp/app/data/dxo"
)

////////////////////////////////////////////////////////////////////////////////

type WeixinPlatformMsgResponseBase struct {
	ToUserName   dxo.WxMsgUserName `xml:"ToUserName"`
	FromUserName dxo.WxMsgUserName `xml:"FromUserName"`
	CreateTime   dxo.WxMsgTime     `xml:"CreateTime"`
	MsgType      WeixinMessageType `xml:"MsgType"`
}

////////////////////////////////////////////////////////////////////////////////

type WeixinPlatformMsgResponseText struct {
	WeixinPlatformMsgResponseBase

	XMLName xml.Name `xml:"xml"`

	Content string `xml:"Content"`
}

//////////////////////////////////////////////////////////////

type WeixinPlatformMsgResponseImage struct {
	WeixinPlatformMsgResponseBase

	XMLName xml.Name `xml:"xml"`

	Image WeixinPlatformMsgResponseImageInner `xml:"Image"`
}

type WeixinPlatformMsgResponseImageInner struct {
	MediaID string `xml:"MediaId"`
}

//////////////////////////////////////////////////////////////

type WeixinPlatformMsgResponseVoice struct {
	WeixinPlatformMsgResponseBase

	XMLName xml.Name `xml:"xml"`

	Voice WeixinPlatformMsgResponseVoiceInner `xml:"Voice"`
}

type WeixinPlatformMsgResponseVoiceInner struct {
	MediaID string `xml:"MediaId"`
}

//////////////////////////////////////////////////////////////

type WeixinPlatformMsgResponseVideo struct {
	WeixinPlatformMsgResponseBase

	XMLName xml.Name `xml:"xml"`

	Video WeixinPlatformMsgResponseVideoInner `xml:"Video"`
}

type WeixinPlatformMsgResponseVideoInner struct {
	MediaID     string `xml:"MediaId"`
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
}

//////////////////////////////////////////////////////////////

type WeixinPlatformMsgResponseMusic struct {
	WeixinPlatformMsgResponseBase

	XMLName xml.Name `xml:"xml"`

	Music WeixinPlatformMsgResponseMusicInfo `xml:"Music"`
}

type WeixinPlatformMsgResponseMusicInfo struct {
	Title        string `xml:"Title"`
	Description  string `xml:"Description"`
	MusicURL     string `xml:"MusicUrl"`
	HQMusicURL   string `xml:"HQMusicUrl"`
	ThumbMediaID string `xml:"ThumbMediaId"`
}

//////////////////////////////////////////////////////////////

type WeixinPlatformMsgResponseNews struct {
	WeixinPlatformMsgResponseBase

	XMLName xml.Name `xml:"xml"`

	ArticleCount int                                   `xml:"ArticleCount"`
	Articles     []WeixinPlatformMsgResponseNewsInner1 `xml:"Articles"`
}

type WeixinPlatformMsgResponseNewsInner1 struct {
	Item WeixinPlatformMsgResponseNewsInner2 `xml:"item"`
}

type WeixinPlatformMsgResponseNewsInner2 struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PictureURL  string `xml:"PicUrl"`
	URL         string `xml:"Url"`
}

////////////////////////////////////////////////////////////////////////////////

type WeixinPlatformMsgResponseHolder struct {
	Text *WeixinPlatformMsgResponseText

	Image *WeixinPlatformMsgResponseImage

	Voice *WeixinPlatformMsgResponseVoice

	Video *WeixinPlatformMsgResponseVideo

	Music *WeixinPlatformMsgResponseMusic

	News *WeixinPlatformMsgResponseNews
}
