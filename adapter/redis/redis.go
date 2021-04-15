package redis

import r "github.com/go-redis/redis"

type (
	Connector struct {
		instance *r.Client
		address  string
		password string
		db       int
		strings  Strings
		lists    Lists
		sets     Sets
		hashes   Hashes
		streams  Streams
	}
)

func (con *Connector) String() Strings {
	// Success
	return con.strings
}

func (con *Connector) Lists() Lists {
	// Success
	return con.lists
}

func (con *Connector) Sets() Sets {
	// Success
	return con.sets
}

func (con *Connector) Hashes() Hashes {
	// Success
	return con.hashes
}

func (con *Connector) Streams() Streams {
	// Success
	return con.streams
}

func NewClient(conf Config) (Redis, error) {
	con := &Connector{
		instance: r.NewClient(&r.Options{
			Addr:     conf.Address,
			Password: conf.Password,
			DB:       conf.Database,
		}),
		address:  conf.Address,
		password: conf.Password,
		db:       conf.Database,
	}
	if err := con.Ping(); err != nil {
		return nil, err
	}
	con.strings = NewStringsHelper(con)
	// Success
	return con, nil
}

func (con *Connector) Ping() error {
	// Success
	return con.instance.Ping().Err()
}
