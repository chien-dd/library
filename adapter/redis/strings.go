package redis

import (
	"encoding/json"
	r "github.com/go-redis/redis"
	"time"
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

func (con *StringsHelper) PGet(key string, pointer interface{}) error {
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
