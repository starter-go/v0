package api

import (
	"context"
	"crypto/sha1"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/i18n"
)

////////////////////////////////////////////////////////////////////////////////

// HyperError 表示一个适合工作于 http 的错误信息
type HyperError interface {
	error

	GetInfo(dst *HyperErrorInfo)
}

type HyperErrorI18n interface {
	HyperError

	Translate(c context.Context, lang i18n.Language) HyperErrorI18n
}

type HyperErrorInfo struct {
	Code       Code          // 错误码 , hash 头部的 int32
	Message    string        // 经过格式化的错误文本
	Hash       lang.Hex      // sha1sum(uri)
	Language   i18n.Language // 输出文本的本地化语言
	Name       Name          // 错误的名称, 在特定命名空间中是唯一的
	Namespace  Namespace     // 错误的命名空间, 通常是发生错误的包名
	Parent     error         // 堆栈上一层的错误, 可以为 nil
	StatusCode int           // http.status.code
	StatusText string        // http.status.message
	URI        lang.URI      // uri = 'uri:'+ ns + '#' + name
}

////////////////////////////////////////////////////////////////////////////////

func NewHyperError(info *HyperErrorInfo) HyperError {
	if info == nil {
		return nil
	}
	e1 := new(innerHyperError)
	e1.info = *info
	return e1
}

func ComputeInfoFields(info *HyperErrorInfo) {

	if info == nil {
		return
	}

	// uri
	ns := string(info.Namespace)
	name := string(info.Name)
	uri := "uri:" + ns + "#" + name

	// hash
	sum := sha1.Sum([]byte(uri))
	hex := lang.HexFromBytes(sum[:])

	// code
	code := 0
	bin := hex.Bytes()
	for i, b := range bin {
		if i < 4 {
			code = (code << 8) | (0xff & int(b))
		} else {
			break
		}
	}

	// write back
	info.URI = lang.URI(uri)
	info.Hash = hex
	info.Code = Code(code)
}

////////////////////////////////////////////////////////////////////////////////

type innerHyperError struct {
	info HyperErrorInfo
}

// Error implements [HyperError].
func (inst *innerHyperError) Error() string {
	return inst.info.Message
}

// GetInfo implements [HyperError].
func (inst *innerHyperError) GetInfo(dst *HyperErrorInfo) {
	if dst == nil {
		return
	}
	*dst = inst.info
}

func (inst *innerHyperError) _impl() HyperError {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
