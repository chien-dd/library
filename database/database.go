package database

type (
	// Map
	M map[string]interface{}
	// Query
	Doc      M
	Query    M
	QueryES2 M
	QueryES7 M
	// Document
	Document interface {
		GetID() string
	}
	// Database
	Database interface {
		Get(database, collection, id string, result interface{}) error
		Exists(database, collection, id string) (bool, error)
		Count(database, collection string, query Query) (int64, error)
		FindOne(database, collection string, query Query, sorts []string, result interface{}) error
		FindPaging(database, collection string, query Query, sorts []string, page, size int, results interface{}) (int64, error)
		FindOffset(database, collection string, query Query, sorts []string, offset, size int, results interface{}) (int64, error)
		FindScroll(database, collection string, query Query, sorts []string, size int, scrollID string, results interface{}) (string, int64, error)
		InsertOne(database, collection string, doc Document) error
		InsertMany(database, collection string, docs []Document) error
		UpdateByID(database, collection, id string, update interface{}, upsert bool) error
		UpdateOne(database, collection string, query Query, update interface{}, upsert bool) error
		UpdateMany(database, collection string, query Query, update interface{}, upsert bool) error
		DeleteByID(database, collection, id string) error
		DeleteMany(database, collection string, query Query) error
	}
)

func (doc Doc) GetID() string {
	if value, ok := doc["id"]; ok {
		return value.(string)
	}
	return ""
}

func (q Query) Source() interface{} {
	return q
}

func (q QueryES2) Source() interface{} {
	return q
}

func (q QueryES7) Source() (interface{}, error) {
	return q, nil
}
