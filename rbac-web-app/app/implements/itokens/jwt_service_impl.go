package itokens

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
	"github.com/starter-go/v0/rbac-web-app/app/classes/webcontexts"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/vlog"
)

const (
	theJWTGetParamName = "x-jwt"
	theJWTSetParamName = "x-set-jwt"
)

type JWTokenServiceAPI interface {
	tokens.JWTokenService
}

type JWTokenServiceImpl struct {

	//starter:component

	_as func(tokens.JWTokenService) //starter:as("#")

	ConfigUseCookie        bool   //starter:inject("${server.www.jwt.use-cookie}")
	ConfigUseHeader        bool   //starter:inject("${server.www.jwt.use-header}")
	ConfigKeySecret        string //starter:inject("${server.www.jwt.key-secret}")
	ConfigJWTokenMaxAgeSec int    //starter:inject("${server.www.jwt.max-age-in-sec}")
	ConfigDoVerify         bool   //starter:inject("${server.www.jwt.do-verify}")

	cache *innerJWTokenConfig
}

func (inst *JWTokenServiceImpl) _impl() JWTokenServiceAPI {
	return inst
}

func (inst *JWTokenServiceImpl) DecodeJWT(jstr tokens.JWT) (*dto.Token, error) {
	decoder := new(innerJWTokenDecoder)
	decoder.ser = inst
	return decoder.decode(jstr)
}

func (inst *JWTokenServiceImpl) EncodeJWT(token *dto.Token) (tokens.JWT, error) {
	encoder := new(innerJWTokenEncoder)
	encoder.ser = inst
	return encoder.encode(token)
}

func (inst *JWTokenServiceImpl) ReadJWT(cc context.Context) (tokens.JWT, error) {
	reader := new(innerJWTokenReader)
	reader.ser = inst
	return reader.read(cc)
}

func (inst *JWTokenServiceImpl) WriteJWT(cc context.Context, jstr tokens.JWT) error {
	writer := new(innerJWTokenWriter)
	writer.ser = inst
	return writer.write(cc, jstr)
}

func (inst *JWTokenServiceImpl) innerLoadConfig() *innerJWTokenConfig {

	secret := inst.ConfigKeySecret
	sec := inst.ConfigJWTokenMaxAgeSec

	cfg := new(innerJWTokenConfig)

	cfg.keySecret = []byte(secret)
	cfg.maxAge = time.Duration(sec) * time.Second
	cfg.useCookie = inst.ConfigUseCookie
	cfg.useHeader = inst.ConfigUseHeader
	cfg.needVerify = inst.ConfigDoVerify

	return cfg
}

func (inst *JWTokenServiceImpl) innerGetConfig() *innerJWTokenConfig {
	cfg := inst.cache
	if cfg == nil {
		cfg = inst.innerLoadConfig()
		inst.cache = cfg
	}
	return cfg
}

////////////////////////////////////////////////////////////////////////////////
// cache

type innerJWTokenConfig struct {
	keySecret []byte

	maxAge time.Duration // the default value

	needVerify bool

	useHeader bool

	useCookie bool
}

////////////////////////////////////////////////////////////////////////////////
// reader

type innerJWTokenReader struct {
	ser *JWTokenServiceImpl
}

func (inst *innerJWTokenReader) read(cc context.Context) (tokens.JWT, error) {

	h, err := webcontexts.GetHolder(cc)
	if err != nil {
		return "", err
	}
	gc := h.Get()

	token, err := inst.innerReadFromHeader(gc)
	if err == nil {
		return token, nil
	}

	token, err = inst.innerReadFromCookie(gc)
	if err == nil {
		return token, nil
	}

	return "", fmt.Errorf("no token in the HTTP request")
}

func (inst *innerJWTokenReader) innerReadFromHeader(c *gin.Context) (tokens.JWT, error) {
	const name = theJWTGetParamName
	value := c.Request.Header.Get(name)
	value = strings.TrimSpace(value)
	if value == "" {
		return "", fmt.Errorf("no token")
	}
	return tokens.JWT(value), nil
}

func (inst *innerJWTokenReader) innerReadFromCookie(c *gin.Context) (tokens.JWT, error) {
	const name = theJWTGetParamName
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	value := strings.TrimSpace(cookie.Value)
	if value == "" {
		return "", fmt.Errorf("no token")
	}
	return tokens.JWT(value), nil
}

////////////////////////////////////////////////////////////////////////////////
// writer

type innerJWTokenWriter struct {
	ser *JWTokenServiceImpl
}

func (inst *innerJWTokenWriter) write(cc context.Context, jstr tokens.JWT) error {

	h, err := webcontexts.GetHolder(cc)
	if err != nil {
		return err
	}

	gc := h.Get()
	ser := inst.ser
	cfg := ser.innerGetConfig()
	count := 0

	if cfg.useHeader {
		err := inst.innerWriteToHeader(gc, jstr)
		if err != nil {
			vlog.Warn("innerJWTokenWriter.innerWriteToHeader(): %s", err.Error())
		}
		count++
	}

	if cfg.useCookie {
		err = inst.innerWriteToCookie(gc, jstr)
		if err != nil {
			vlog.Warn("innerJWTokenWriter.innerWriteToCookie(): %s", err.Error())
		}
		count++
	}

	if count == 0 {
		vlog.Warn("innerJWTokenWriter.write(): no token write to HTTP response")
	}

	return nil
}

func (inst *innerJWTokenWriter) innerWriteToHeader(c *gin.Context, jstr tokens.JWT) error {
	const name = theJWTSetParamName
	value := jstr.String()
	c.Header(name, value)
	return nil
}

func (inst *innerJWTokenWriter) innerWriteToCookie(c *gin.Context, jstr tokens.JWT) error {

	const name = theJWTSetParamName

	ser := inst.ser
	cfg := ser.innerGetConfig()

	value := jstr.String()
	ma := int(cfg.maxAge / time.Second)

	c.SetCookie(name, value, ma, "", "", false, false)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// encoder

type innerJWTokenEncoder struct {
	ser *JWTokenServiceImpl
}

func (inst *innerJWTokenEncoder) encode(token *dto.Token) (tokens.JWT, error) {

	config := inst.ser.innerGetConfig()
	codec := new(innerJWTCodec)
	attrs := make(map[string]string)

	codec.key = config.keySecret

	if (token.AliveFrom == 0) && (token.AliveTo == 0) {
		// use default (t1,t2)
		inst.applyDefaultAliveFromTo(token, config)
	}

	attrs["t1"] = inst.formatTime(token.AliveFrom)
	attrs["t2"] = inst.formatTime(token.AliveTo)
	attrs["se-id"] = inst.formatSessionID(token.SessionID)
	attrs["se-uuid"] = inst.formatUUID(token.SessionUUID)

	return codec.encode(attrs)
}

func (inst *innerJWTokenEncoder) applyDefaultAliveFromTo(token *dto.Token, cfg *innerJWTokenConfig) {

	now := lang.Now()
	ms := int64(cfg.maxAge / time.Millisecond)

	token.AliveFrom = now - (10 * 1000)
	token.AliveTo = now + lang.Time(ms)
}

func (inst *innerJWTokenEncoder) formatTime(value lang.Time) string {
	num := value.Int()
	return strconv.FormatInt(num, 10)
}

func (inst *innerJWTokenEncoder) formatUUID(value lang.UUID) string {
	return value.String()
}

func (inst *innerJWTokenEncoder) formatSessionID(value dxo.SessionID) string {
	num := int64(value)
	return strconv.FormatInt(num, 10)
}

////////////////////////////////////////////////////////////////////////////////
// decoder

type innerJWTokenDecoder struct {
	ser *JWTokenServiceImpl

	err error
}

func (inst *innerJWTokenDecoder) decode(jstr tokens.JWT) (*dto.Token, error) {

	config := inst.ser.innerGetConfig()
	codec := new(innerJWTCodec)
	out := new(dto.Token)

	codec.key = config.keySecret
	attrs, err := codec.decode(jstr)
	if err != nil {
		return nil, err
	}

	t1 := attrs["t1"]
	t2 := attrs["t2"]
	sesID := attrs["se-id"]
	sesUUID := attrs["se-uuid"]

	out.AliveFrom = inst.parseTime(t1)
	out.AliveTo = inst.parseTime(t2)
	out.SessionID = inst.parseSessionID(sesID)
	out.SessionUUID = inst.parseUUID(sesUUID)

	err = inst.err
	if err != nil {
		return nil, err
	}

	if config.needVerify {
		err = inst.verify(out)
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (inst *innerJWTokenDecoder) verify(tk *dto.Token) error {

	if tk == nil {
		return fmt.Errorf("token is nil")
	}

	if tk.SessionID < 1 {
		return fmt.Errorf("bad session-id")
	}
	if len(tk.SessionUUID) < 10 {
		return fmt.Errorf("bad session-uuid")
	}

	t1 := tk.AliveFrom
	t2 := tk.AliveTo
	now := lang.Now()

	if (0 < t1) && (t1 < t2) {
		if (t1 <= now) && (now <= t2) {
			return nil // ok
		}
	}
	return fmt.Errorf("token is NOT alive")
}

func (inst *innerJWTokenDecoder) parseInt64(str string) int64 {
	str = strings.TrimSpace(str)
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		inst.err = err
	}
	return num
}

func (inst *innerJWTokenDecoder) parseTime(str string) lang.Time {
	num := inst.parseInt64(str)
	return lang.Time(num)
}

func (inst *innerJWTokenDecoder) parseSessionID(str string) dxo.SessionID {
	num := inst.parseInt64(str)
	return dxo.SessionID(num)
}

func (inst *innerJWTokenDecoder) parseUUID(str string) lang.UUID {
	str = strings.TrimSpace(str)
	return lang.UUID(str)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
