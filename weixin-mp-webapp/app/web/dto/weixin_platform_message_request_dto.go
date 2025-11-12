package dto

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/weixin-mp-webapp/app/data/dxo"
)

type WeixinMessageType = dxo.WxMsgType

const (
	WxMsgTypeText       WeixinMessageType = dxo.WxMsgTypeText
	WxMsgTypeImage      WeixinMessageType = dxo.WxMsgTypeImage
	WxMsgTypeVoice      WeixinMessageType = dxo.WxMsgTypeVoice
	WxMsgTypeVideo      WeixinMessageType = dxo.WxMsgTypeVideo
	WxMsgTypeShortVideo WeixinMessageType = dxo.WxMsgTypeShortVideo
	WxMsgTypeLocation   WeixinMessageType = dxo.WxMsgTypeLocation
	WxMsgTypeLink       WeixinMessageType = dxo.WxMsgTypeLink
)

// func (t WeixinMessageType) Normalize() WeixinMessageType {
// 	str := t.String()
// 	str = strings.TrimSpace(str)
// 	str = strings.ToLower(str)
// 	return WeixinMessageType(str)
// }

// func (t WeixinMessageType) String() string {
// 	return string(t)
// }

////////////////////////////////////////////////////////////////////////////////

// WeixinPlatformMessageBase ...
type WeixinPlatformMessageBase struct {
	ToUserName   dxo.WxMsgUserName `xml:"ToUserName"`
	FromUserName dxo.WxMsgUserName `xml:"FromUserName"`
	CreateTime   dxo.WxMsgTime     `xml:"CreateTime"`
	MsgType      dxo.WxMsgType     `xml:"MsgType"`
	MsgId        string            `xml:"MsgId"`
	MsgDataId    string            `xml:"MsgDataId"`
	Idx          string            `xml:"Idx"`
}

type WeixinPlatformMessageHolder struct {
	Text       *WeixinPlatformMessageText       `json:"text"`
	Image      *WeixinPlatformMessageImage      `json:"image"`
	Voice      *WeixinPlatformMessageVoice      `json:"voice"`
	Video      *WeixinPlatformMessageVideo      `json:"video"`
	ShortVideo *WeixinPlatformMessageShortVideo `json:"svideo"`
	Location   *WeixinPlatformMessageLocation   `json:"location"`
	Link       *WeixinPlatformMessageLink       `json:"link"`
}

////////////////////////////////////////////////////////////////////////////////

type WeixinPlatformMessageText struct {
	WeixinPlatformMessageBase

	Content string `xml:"Content"`
}

type WeixinPlatformMessageImage struct {
	WeixinPlatformMessageBase

	PictureURL string `xml:"PicUrl"`
	MediaID    string `xml:"MediaId"`
}

type WeixinPlatformMessageVoice struct {
	WeixinPlatformMessageBase

	MediaID    string `xml:"MediaId"`
	Format     string `xml:"Format"`
	MediaID16K string `xml:"MediaId16K"`
}

type WeixinPlatformMessageVideo struct {
	WeixinPlatformMessageBase

	MediaID    string `xml:"MediaId"`
	Format     string `xml:"Format"`
	MediaID16K string `xml:"MediaId16K"`
}

type WeixinPlatformMessageShortVideo struct {
	WeixinPlatformMessageBase

	MediaID      string `xml:"MediaId"`
	ThumbMediaID string `xml:"ThumbMediaId"`
}

type WeixinPlatformMessageLocation struct {
	WeixinPlatformMessageBase

	LocationX string `xml:"Location_X"`
	LocationY string `xml:"Location_Y"`
	Scale     string `xml:"Scale"`
	Label     string `xml:"Label"`
}

type WeixinPlatformMessageLink struct {
	WeixinPlatformMessageBase

	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	URL         string `xml:"Url"`
}

func (inst *WeixinPlatformMessageBase) Normalize() *WeixinPlatformMessageBase {
	o2 := new(WeixinPlatformMessageBase)
	nor := new(innerWeixinPlatformMessageNormalizer)
	nor.normalizeBase(inst, o2)
	return o2
}

func (inst *WeixinPlatformMessageText) Normalize() *WeixinPlatformMessageText {
	o2 := new(WeixinPlatformMessageText)
	nor := new(innerWeixinPlatformMessageNormalizer)
	nor.normalizeText(inst, o2)
	return o2
}

func (inst *WeixinPlatformMessageImage) Normalize() *WeixinPlatformMessageImage {
	o2 := new(WeixinPlatformMessageImage)
	nor := new(innerWeixinPlatformMessageNormalizer)
	nor.normalizeImage(inst, o2)
	return o2
}

func (inst *WeixinPlatformMessageVideo) Normalize() *WeixinPlatformMessageVideo {
	o2 := new(WeixinPlatformMessageVideo)
	nor := new(innerWeixinPlatformMessageNormalizer)
	nor.normalizeVideo(inst, o2)
	return o2
}

func (inst *WeixinPlatformMessageShortVideo) Normalize() *WeixinPlatformMessageShortVideo {
	o2 := new(WeixinPlatformMessageShortVideo)
	nor := new(innerWeixinPlatformMessageNormalizer)
	nor.normalizeShortVideo(inst, o2)
	return o2
}

func (inst *WeixinPlatformMessageLocation) Normalize() *WeixinPlatformMessageLocation {
	o2 := new(WeixinPlatformMessageLocation)
	nor := new(innerWeixinPlatformMessageNormalizer)
	nor.normalizeLocation(inst, o2)
	return o2
}

func (inst *WeixinPlatformMessageLink) Normalize() *WeixinPlatformMessageLink {
	o2 := new(WeixinPlatformMessageLink)
	nor := new(innerWeixinPlatformMessageNormalizer)
	nor.normalizeLink(inst, o2)
	return o2
}

func (inst *WeixinPlatformMessageVoice) Normalize() *WeixinPlatformMessageVoice {
	o2 := new(WeixinPlatformMessageVoice)
	nor := new(innerWeixinPlatformMessageNormalizer)
	nor.normalizeVoice(inst, o2)
	return o2
}

////////////////////////////////////////////////////////////////////////////////

type innerWeixinPlatformMessageNormalizer struct {
}

func (inst *innerWeixinPlatformMessageNormalizer) normalizeString(src string) string {
	return strings.TrimSpace(src)
}

func (inst *innerWeixinPlatformMessageNormalizer) normalizeUserName(src dxo.WxMsgUserName) dxo.WxMsgUserName {
	str := string(src)
	str = strings.TrimSpace(str)
	return dxo.WxMsgUserName(str)
}

func (inst *innerWeixinPlatformMessageNormalizer) normalizeBase(src, dst *WeixinPlatformMessageBase) {

	dst.FromUserName = inst.normalizeUserName(src.FromUserName)
	dst.ToUserName = inst.normalizeUserName(src.ToUserName)

	dst.CreateTime = src.CreateTime

	dst.MsgId = inst.normalizeString(src.MsgId)
	dst.MsgDataId = inst.normalizeString(src.MsgDataId)
	dst.Idx = inst.normalizeString(src.Idx)

	dst.MsgType = src.MsgType.Normalize()
}

func (inst *innerWeixinPlatformMessageNormalizer) normalizeText(src, dst *WeixinPlatformMessageText) {

	inst.normalizeBase(&src.WeixinPlatformMessageBase, &dst.WeixinPlatformMessageBase)

	dst.Content = inst.normalizeString(src.Content)
}

func (inst *innerWeixinPlatformMessageNormalizer) normalizeImage(src, dst *WeixinPlatformMessageImage) {

	inst.normalizeBase(&src.WeixinPlatformMessageBase, &dst.WeixinPlatformMessageBase)

	dst.MediaID = inst.normalizeString(src.MediaID)
	dst.PictureURL = inst.normalizeString(src.PictureURL)
}

func (inst *innerWeixinPlatformMessageNormalizer) normalizeVideo(src, dst *WeixinPlatformMessageVideo) {

	inst.normalizeBase(&src.WeixinPlatformMessageBase, &dst.WeixinPlatformMessageBase)

	dst.Format = inst.normalizeString(src.Format)
	dst.MediaID = inst.normalizeString(src.MediaID)
	dst.MediaID16K = inst.normalizeString(src.MediaID16K)
}

func (inst *innerWeixinPlatformMessageNormalizer) normalizeShortVideo(src, dst *WeixinPlatformMessageShortVideo) {

	inst.normalizeBase(&src.WeixinPlatformMessageBase, &dst.WeixinPlatformMessageBase)

	dst.MediaID = inst.normalizeString(src.MediaID)
	dst.ThumbMediaID = inst.normalizeString(src.ThumbMediaID)
}

func (inst *innerWeixinPlatformMessageNormalizer) normalizeLocation(src, dst *WeixinPlatformMessageLocation) {

	inst.normalizeBase(&src.WeixinPlatformMessageBase, &dst.WeixinPlatformMessageBase)

	dst.LocationX = inst.normalizeString(src.LocationX)
	dst.LocationY = inst.normalizeString(src.LocationY)
	dst.Label = inst.normalizeString(src.Label)
	dst.Scale = inst.normalizeString(src.Scale)
}

func (inst *innerWeixinPlatformMessageNormalizer) normalizeLink(src, dst *WeixinPlatformMessageLink) {

	inst.normalizeBase(&src.WeixinPlatformMessageBase, &dst.WeixinPlatformMessageBase)

	dst.Description = inst.normalizeString(src.Description)
	dst.Title = inst.normalizeString(src.Title)
	dst.URL = inst.normalizeString(src.URL)
}

func (inst *innerWeixinPlatformMessageNormalizer) normalizeVoice(src, dst *WeixinPlatformMessageVoice) {

	inst.normalizeBase(&src.WeixinPlatformMessageBase, &dst.WeixinPlatformMessageBase)

	dst.MediaID = inst.normalizeString(src.MediaID)
	dst.MediaID16K = inst.normalizeString(src.MediaID16K)
	dst.Format = inst.normalizeString(src.Format)
}

////////////////////////////////////////////////////////////////////////////////

// WeixinPlatformQuery 是表示微信服务号消息请求的 query 参数
type WeixinPlatformQuery struct {
	Signature string `json:"signature"`
	TimeStamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	EchoStr   string `json:"echostr"`
}

func (inst *WeixinPlatformQuery) VerifyWithToken(token string) error {

	timestamp := inst.TimeStamp
	nonce := inst.Nonce
	builder := new(strings.Builder)

	list := []string{token, timestamp, nonce}
	sort.Strings(list)

	for _, item := range list {
		builder.WriteString(item)
	}

	data := builder.String()
	sum := sha1.Sum([]byte(data))
	hex := lang.HexFromBytes(sum[:])

	if hex.String() == inst.Signature {
		return nil
	}
	return fmt.Errorf("WeixinPlatformQuery: bad message signature")
}

////////////////////////////////////////////////////////////////////////////////
// EOF
