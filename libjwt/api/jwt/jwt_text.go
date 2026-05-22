package jwt

// 表示字符串形式的 JWT
type Text string

func (txt Text) String() string {
	return string(txt)
}
