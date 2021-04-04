package elastic

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"strings"

	es "github.com/olivere/elastic/v7"

	"github.com/chien-dd/library/database"
)

type (
	ES struct {
		model *es.Client
	}

	BulkService struct {
		helper *es.BulkService
	}

	ScrollService struct {
		helper *es.ScrollService
	}
)

func NewElastic(addr string) (*ES, error) {
	client, err := es.NewClient(es.SetURL(addr), es.SetSniff(false))
	if err != nil {
		return nil, err
	}
	// Success
	return &ES{
		model: client,
	}, nil
}

func (con *ES) Bulk() *BulkService {
	// Success
	return &BulkService{helper: es.NewBulkService(con.model)}
}

func (bs *BulkService) Index(index, id string, doc interface{}) {
	req := es.NewBulkIndexRequest().
		Index(index).
		Id(id).
		Doc(doc)
	bs.helper = bs.helper.Add(req)
	// Success
	return
}

func (bs *BulkService) Update(index, id string, update interface{}, upsert bool) {
	req := es.NewBulkUpdateRequest().
		Index(index).
		Id(id).
		Doc(update).
		DocAsUpsert(upsert)
	bs.helper = bs.helper.Add(req)
	// Success
	return
}

func (bs *BulkService) Delete(index, id string) {
	req := es.NewBulkDeleteRequest().
		Index(index).
		Id(id)
	bs.helper = bs.helper.Add(req)
	// Success
	return
}

func (bs *BulkService) Do() error {
	defer bs.helper.Reset()
	_, err := bs.helper.Refresh("true").Do(context.Background())
	if err != nil {
		return err
	}
	// Success
	return nil
}

func (con *ES) Scroll() *ScrollService {
	// Success
	return &ScrollService{helper: es.NewScrollService(con.model)}
}

func (ss *ScrollService) Search(index string, query database.Query, sorts []string, size int, scrollID string) (*es.SearchResult, error) {
	var q database.QueryES7
	value, ok := reflect.ValueOf(query).Convert(reflect.TypeOf(q)).Interface().(database.QueryES7)
	if !ok {
		return nil, database.ReflectError
	}
	q = value
	service := ss.helper.Index(index).Query(q).Size(size).ScrollId(scrollID).Scroll("10m")
	if sorts != nil && len(sorts) > 0 {
		for _, sort := range sorts {
			if strings.HasPrefix(sort, "-") {
				service = service.Sort(strings.TrimPrefix(sort, "-"), false)
			} else if strings.HasPrefix(sort, "+") {
				service = service.Sort(strings.TrimPrefix(sort, "+"), true)
			}
		}
	}
	result, err := service.Do(context.Background())
	if err != nil {
		if err == io.EOF {
			return nil, database.NotFoundError
		}
		if ex := err.(*es.Error); ex.Status == http.StatusBadRequest || ex.Status == http.StatusNotFound {
			return nil, database.NotFoundError
		}
		return nil, err
	}
	// Success
	return result, nil
}

func (con *ES) IndexExists(index string) (bool, error) {
	// Success
	return con.model.IndexExists(index).
		Do(context.Background())
}

func (con *ES) CreateIndex(index string, mapping database.M) (*es.IndicesCreateResult, error) {
	// Success
	return con.model.CreateIndex(index).
		BodyJson(mapping).
		Do(context.Background())
}

func (con *ES) DeleteIndex(index string) (*es.IndicesDeleteResponse, error) {
	// Success
	return con.model.DeleteIndex(index).
		Do(context.Background())
}

func (con *ES) Count(index string, query database.Query) (int64, error) {
	var q database.QueryES7
	value, ok := reflect.ValueOf(query).Convert(reflect.TypeOf(q)).Interface().(database.QueryES7)
	if !ok {
		return 0, database.ReflectError
	}
	q = value
	// Success
	return con.model.Count().
		Index(index).
		Query(q).
		Do(context.Background())
}

func (con *ES) SearchScroll(index string, query database.Query, sorts []string, site int, scrollID string) (*es.SearchResult, error) {
	// Success
	return con.Scroll().Search(index, query, sorts, site, scrollID)
}

func (con *ES) SearchPaging(index string, query database.Query, sorts []string, page, size int) (*es.SearchResult, error) {
	// Success
	return con.SearchOffset(index, query, sorts, page*size, size)
}

func (con *ES) SearchOffset(index string, query database.Query, sorts []string, offset, size int) (*es.SearchResult, error) {
	var q database.QueryES7
	value, ok := reflect.ValueOf(query).Convert(reflect.TypeOf(q)).Interface().(database.QueryES7)
	if !ok {
		return nil, database.ReflectError
	}
	q = value
	service := con.model.Search().Index(index)
	if sorts != nil && len(sorts) > 0 {
		for _, sort := range sorts {
			if strings.HasPrefix(sort, "-") {
				service = service.Sort(strings.TrimPrefix(sort, "-"), false)
			} else if strings.HasPrefix(sort, "+") {
				service = service.Sort(strings.TrimPrefix(sort, "+"), true)
			}
		}
	}
	service = service.Query(q)
	if size == 0 {
		size = 10
	}
	service = service.Size(size).From(offset)
	result, err := service.Do(context.Background())
	if err != nil {
		if ex := err.(*es.Error); ex.Status == http.StatusBadRequest || ex.Status == http.StatusNotFound {
			return nil, database.NotFoundError
		}
		return nil, err
	}
	// Success
	return result, nil
}

func (con *ES) Get(index, id string) (*es.GetResult, error) {
	// Success
	return con.model.Get().
		Index(index).
		Id(id).
		Refresh("true").
		Do(context.Background())
}

func (con *ES) Exists(index, id string) (bool, error) {
	// Success
	return con.model.Exists().
		Index(index).
		Id(id).
		Refresh("true").
		Do(context.Background())
}

func (con *ES) Index(index string, doc database.Document) (*es.IndexResponse, error) {
	// Success
	return con.model.Index().
		Index(index).
		Id(doc.GetID()).
		BodyJson(doc).
		Refresh("true").
		Do(context.Background())
}

func (con *ES) Update(index, id string, update interface{}, upsert bool) (*es.UpdateResponse, error) {
	// Success
	return con.model.Update().
		Index(index).
		Id(id).
		Doc(update).
		DocAsUpsert(upsert).
		Refresh("true").
		Do(context.Background())
}

func (con *ES) DeleteByID(index, id string) (*es.DeleteResponse, error) {
	// Success
	return con.model.Delete().
		Index(index).
		Id(id).
		Refresh("true").
		Do(context.Background())
}

func (con *ES) DeleteByQuery(index string, query database.Query) (*es.BulkIndexByScrollResponse, error) {
	var q database.QueryES7
	value, ok := reflect.ValueOf(query).Convert(reflect.TypeOf(q)).Interface().(database.QueryES7)
	if !ok {
		return nil, database.ReflectError
	}
	q = value
	// Success
	return con.model.DeleteByQuery().
		Index(index).
		Query(q).
		Refresh("true").
		Do(context.Background())
}
