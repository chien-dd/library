package redis

import (
	"encoding/json"
	"time"

	r "github.com/go-redis/redis"
)

type StringsHelper struct {
	instance *r.Client
}

func NewStringsHelper(con *Connector) Strings {
	// Success
	return &StringsHelper{instance: con.instance}
}

func (con *StringsHelper) Get(key string) (string, error) {
	// Success
	return con.instance.Get(key).Result()
}

func (con *StringsHelper) GetS(key string, pointer interface{}) error {
	bts, err := con.instance.Get(key).Bytes()
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bts, pointer); err != nil {
		return err
	}
	// Success
	return nil
}

func (con *StringsHelper) Set(key, value string, expires time.Duration) error {
	// Success
	return con.instance.Set(key, value, expires).Err()
}

func (con *StringsHelper) SetS(key string, data interface{}, expires time.Duration) error {
	bts, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Success
	return con.Set(key, string(bts), expires)
}
