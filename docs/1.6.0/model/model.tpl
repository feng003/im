package {{.pkg}}
{{if .withCache}}
import (
    "fmt"
    "context"
    "database/sql"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/Masterminds/squirrel"
)
{{else}}
import (
    "fmt"
    "context"
    "database/sql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/Masterminds/squirrel"
)

{{end}}
var _ {{.upperStartCamelObject}}Model = (*custom{{.upperStartCamelObject}}Model)(nil)

type (
	// {{.upperStartCamelObject}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Model.
	{{.upperStartCamelObject}}Model interface {
		{{.lowerStartCamelObject}}Model
		Trans(ctx context.Context,fn func(context context.Context,session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		FindOneByQuery(ctx context.Context,rowBuilder squirrel.SelectBuilder) (*{{.upperStartCamelObject}},error)
		FindSum(ctx context.Context,sumBuilder squirrel.SelectBuilder) (float64,error)
		FindCount(ctx context.Context,countBuilder squirrel.SelectBuilder) (int64,error)
		FindAll(ctx context.Context,rowBuilder squirrel.SelectBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error)
		FindPageListByPage(ctx context.Context,rowBuilder squirrel.SelectBuilder,page ,pageSize int64,orderBy string) ([]*{{.upperStartCamelObject}},error)
		FindPageListByIdDESC(ctx context.Context,rowBuilder squirrel.SelectBuilder ,preMinId ,pageSize int64) ([]*{{.upperStartCamelObject}},error)
		FindPageListByIdASC(ctx context.Context,rowBuilder squirrel.SelectBuilder,preMaxId ,pageSize int64) ([]*{{.upperStartCamelObject}},error)

        BulkInsert(ctx context.Context, field string, data [][]interface{}) (sql.Result, error)
        BulkUpdate(ctx context.Context, data map[string]map[interface{}]interface{}, where map[string]interface{}) (sql.Result, error)
}

	custom{{.upperStartCamelObject}}Model struct {
		*default{{.upperStartCamelObject}}Model
	}
)

// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf{{end}}) {{.upperStartCamelObject}}Model {
	return &custom{{.upperStartCamelObject}}Model{
		default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(conn{{if .withCache}}, c{{end}}),
	}
}

func (m *default{{.upperStartCamelObject}}Model) FindOneByQuery(ctx context.Context,rowBuilder squirrel.SelectBuilder) (*{{.upperStartCamelObject}},error) {

	query, values, err := rowBuilder.Where("1 = 1").ToSql()
	if err != nil {
		return nil, err
	}

	var resp {{.upperStartCamelObject}}
	{{if .withCache}}err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return &resp, nil
    case sqlx.ErrNotFound:
        return nil, ErrNotFound
	default:
		return nil, err
	}
}


func (m *default{{.upperStartCamelObject}}Model) FindSum(ctx context.Context,sumBuilder squirrel.SelectBuilder) (float64,error) {

	query, values, err := sumBuilder.Where("1 = 1").ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	{{if .withCache}}err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) FindCount(ctx context.Context,countBuilder squirrel.SelectBuilder) (int64,error) {

	query, values, err := countBuilder.Where("1 = 1").ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	{{if .withCache}}err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) FindAll(ctx context.Context,rowBuilder squirrel.SelectBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error) {

	if orderBy == ""{
		rowBuilder = rowBuilder.OrderBy("id DESC")
	}else{
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("1 = 1").ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	{{if .withCache}}err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) FindPageListByPage(ctx context.Context,rowBuilder squirrel.SelectBuilder,page ,pageSize int64,orderBy string) ([]*{{.upperStartCamelObject}},error) {

	if orderBy == ""{
		rowBuilder = rowBuilder.OrderBy("id DESC")
	}else{
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1{
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Where("1 = 1").Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	{{if .withCache}}err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) FindPageListByIdDESC(ctx context.Context,rowBuilder squirrel.SelectBuilder ,preMinId ,pageSize int64) ([]*{{.upperStartCamelObject}},error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? " , preMinId)
	}

	query, values, err := rowBuilder.Where("1 = 1").OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	{{if .withCache}}err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//按照id升序分页查询数据，不支持排序
func (m *default{{.upperStartCamelObject}}Model) FindPageListByIdASC(ctx context.Context,rowBuilder squirrel.SelectBuilder,preMaxId ,pageSize int64) ([]*{{.upperStartCamelObject}},error)  {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? " , preMaxId)
	}

	query, values, err := rowBuilder.Where("1 = 1").OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	{{if .withCache}}err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}


func (m *default{{.upperStartCamelObject}}Model) BulkInsert(ctx context.Context, field string, data [][]interface{}) (sql.Result, error) {

	if len(data) <= 0 || len(field) <= 0 {
		return nil, ErrNotFound
	}
	builder := m.InsertBuilder(field)
	for _, v := range data {
		builder = builder.Values(v...)
	}
	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
    result, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        return conn.ExecCtx(ctx, query, values...)
    })
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (m *default{{.upperStartCamelObject}}Model) BulkUpdate(ctx context.Context, data map[string]map[interface{}]interface{}, where map[string]interface{}) (sql.Result, error) {
	if len(data) <= 0 || len(where) <= 0 {
		return nil, ErrNotFound
	}

	builder := m.UpdateBuilder()
	for column, cases := range data {
		caseStmt := squirrel.Case("id")
		for id, value := range cases {
			var strId, strValue string
			switch v := id.(type) {
			case string:
				strId = fmt.Sprintf("'%s'", v) // Ensure strings are quoted
			default:
				strId = fmt.Sprintf("%v", v)
			}
			switch val := value.(type) {
            case string:
                strValue = fmt.Sprintf("'%s'", val)
            default:
                strValue = fmt.Sprintf("%v", val)
            }
			caseStmt = caseStmt.When(strId, strValue)
		}
		builder = builder.Set(column, caseStmt)
	}
	query, values, err := builder.Where(where).ToSql()
	if err != nil {
		return nil, err
	}

    result, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        return conn.ExecCtx(ctx, query, values...)
    })
	if err != nil {
		return nil, err
	}

	return result, nil
}

// export logic
func (m *default{{.upperStartCamelObject}}Model) Trans(ctx context.Context,fn func(ctx context.Context,session sqlx.Session) error) error {
	{{if .withCache}}
	return m.TransactCtx(ctx,func(ctx context.Context,session sqlx.Session) error {
		return  fn(ctx,session)
	})
	{{else}}
	return m.conn.TransactCtx(ctx,func(ctx context.Context,session sqlx.Session) error {
		return  fn(ctx,session)
	})
	{{end}}
}

// export logic
func (m *default{{.upperStartCamelObject}}Model) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select({{.lowerStartCamelObject}}Rows).From(m.table)
}

// export logic
func (m *default{{.upperStartCamelObject}}Model) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT("+field+")").From(m.table)
}

// export logic
func (m *default{{.upperStartCamelObject}}Model) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM("+field+"),0)").From(m.table)
}

//export logic
func (m *default{{.upperStartCamelObject}}Model) InsertBuilder(field string) squirrel.InsertBuilder {
	return squirrel.Insert(m.table).Columns(field)
}

func (m *default{{.upperStartCamelObject}}Model) UpdateBuilder() squirrel.UpdateBuilder {
	return squirrel.Update(m.table)
}