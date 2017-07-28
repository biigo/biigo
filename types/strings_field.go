package types

import (
	"database/sql/driver"
	"encoding/json"
)

// StringsField 描述字符串数组字段类型
type StringsField []string

// Append items
func (sc *StringsField) Append(items ...string) {
	*sc = append(*sc, items...)
}

// Slices 转换成切片数据类型
func (sc StringsField) Slices() []string {
	return sc
}

// Count 返回集合数量
func (sc StringsField) Count() uint {
	return uint(len(sc))
}

// Value []String to json array
func (sc StringsField) Value() (driver.Value, error) {
	return json.Marshal(sc)
}

// Scan json array to []String
func (sc *StringsField) Scan(src interface{}) error {
	srcBytes := src.([]byte)
	return json.Unmarshal(srcBytes, sc)
}
