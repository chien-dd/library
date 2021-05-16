package redis

import "time"

type (
	Redis interface {
		// Common
		Exists(key string) bool
		Ping() error
		String() Strings
		Lists() Lists
		Sets() Sets
		Hashes() Hashes
		Streams() Streams
	}

	Strings interface {
		Get(key string) (string, error)
		GetS(key string, pointer interface{}) error
		Set(key, value string, expires time.Duration) error
		SetS(key string, data interface{}, expires time.Duration) error
	}

	Lists interface {
	}

	Sets interface {
	}

	Hashes interface {
	}

	Streams interface {
	}
)
