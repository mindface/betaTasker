package model

import (
	"database/sql/driver"
	"fmt"
)

type LabelJSON []byte

// DBに保存する際に呼ばれる
func (j LabelJSON) Value() (driver.Value, error) {
	return string(j), nil
}

// DBから読み込む際に呼ばれる
func (j *LabelJSON) Scan(value interface{}) error {
	if value == nil {
		*j = LabelJSON("{}")
		return nil
	}

	switch v := value.(type) {
	case []byte:
		*j = LabelJSON(v)
	case string:
		*j = LabelJSON([]byte(v))
	default:
		return fmt.Errorf("unsupported type for LabelJSON: %T", v)
	}
	return nil
}
