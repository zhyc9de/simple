package sqls

import (
	"database/sql"

	"github.com/mlogclub/simple/common/strs"
)

func SqlNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  len(value) > 0,
	}
}

var Wrapper = "`"

func KeywordWrap(keyword string) string {
	if strs.IsBlank(keyword) {
		return keyword
	}
	return Wrapper + keyword + Wrapper
}
