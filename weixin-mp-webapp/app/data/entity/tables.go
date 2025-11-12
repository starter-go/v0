package entity

import "strings"

func ListAllTables(prefix string) []any {

	list := make([]any, 0)

	list = append(list, new(Example))

	innerUpdateTableNamePrefix(prefix)
	return list
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamePrefix string = "weixin_mp_"

func innerUpdateTableNamePrefix(prefix string) {
	prefix = strings.TrimSpace(prefix)
	if prefix == "" {
		return
	}
	theTableNamePrefix = prefix
}

////////////////////////////////////////////////////////////////////////////////

func (Example) TableName() string {
	return theTableNamePrefix + "examples"
}
