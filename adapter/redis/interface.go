package redis

import "time"

type (
	Redis interface {
		String() Strings
		Lists() Lists
		Sets() Sets
		Hashes() Hashes
		Streams() Streams
	}

	Strings interface {
		Get(key string) (string, error)
		PGet(key string, pointer interface{}) error
		Set(key, value string, expires time.Duration) error
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
