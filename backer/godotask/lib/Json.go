package lib

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
  "log"
)

type JSON map[string]interface{}



func (j JSON) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, j)
}

func JSONUnm(text string) (JSON, error) {
	var j JSON

	if err := json.Unmarshal([]byte(text), &j); err != nil {
		log.Fatal(err)
	}

	return j, nil
}
