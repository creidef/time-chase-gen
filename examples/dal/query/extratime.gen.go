// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gorm.io/gen/examples/dal/model"
)

func newExtratime(db *gorm.DB, opts ...gen.DOOption) extratime {
	_extratime := extratime{}

	_extratime.extratimeDo.UseDB(db, opts...)
	_extratime.extratimeDo.UseModel(&model.Extratime{})

	tableName := _extratime.extratimeDo.TableName()
	_extratime.ALL = field.NewAsterisk(tableName)
	_extratime.ID = field.NewInt32(tableName, "id")
	_extratime.Name = field.NewString(tableName, "name")
	_extratime.Factor = field.NewFloat64(tableName, "factor")

	_extratime.fillFieldMap()

	return _extratime
}

type extratime struct {
	extratimeDo extratimeDo

	ALL    field.Asterisk
	ID     field.Int32
	Name   field.String
	Factor field.Float64

	fieldMap map[string]field.Expr
}

func (e extratime) Table(newTableName string) *extratime {
	e.extratimeDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e extratime) As(alias string) *extratime {
	e.extratimeDo.DO = *(e.extratimeDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *extratime) updateTableName(table string) *extratime {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Name = field.NewString(table, "name")
	e.Factor = field.NewFloat64(table, "factor")

	e.fillFieldMap()

	return e
}

func (e *extratime) WithContext(ctx context.Context) *extratimeDo {
	return e.extratimeDo.WithContext(ctx)
}

func (e extratime) TableName() string { return e.extratimeDo.TableName() }

func (e extratime) Alias() string { return e.extratimeDo.Alias() }

func (e *extratime) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *extratime) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 3)
	e.fieldMap["id"] = e.ID
	e.fieldMap["name"] = e.Name
	e.fieldMap["factor"] = e.Factor
}

func (e extratime) clone(db *gorm.DB) extratime {
	e.extratimeDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e extratime) replaceDB(db *gorm.DB) extratime {
	e.extratimeDo.ReplaceDB(db)
	return e
}

type extratimeDo struct{ gen.DO }

func (e extratimeDo) Debug() *extratimeDo {
	return e.withDO(e.DO.Debug())
}

func (e extratimeDo) WithContext(ctx context.Context) *extratimeDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e extratimeDo) ReadDB() *extratimeDo {
	return e.Clauses(dbresolver.Read)
}

func (e extratimeDo) WriteDB() *extratimeDo {
	return e.Clauses(dbresolver.Write)
}

func (e extratimeDo) Session(config *gorm.Session) *extratimeDo {
	return e.withDO(e.DO.Session(config))
}

func (e extratimeDo) Clauses(conds ...clause.Expression) *extratimeDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e extratimeDo) Returning(value interface{}, columns ...string) *extratimeDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e extratimeDo) Not(conds ...gen.Condition) *extratimeDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e extratimeDo) Or(conds ...gen.Condition) *extratimeDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e extratimeDo) Select(conds ...field.Expr) *extratimeDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e extratimeDo) Where(conds ...gen.Condition) *extratimeDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e extratimeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *extratimeDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e extratimeDo) Order(conds ...field.Expr) *extratimeDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e extratimeDo) Distinct(cols ...field.Expr) *extratimeDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e extratimeDo) Omit(cols ...field.Expr) *extratimeDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e extratimeDo) Join(table schema.Tabler, on ...field.Expr) *extratimeDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e extratimeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *extratimeDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e extratimeDo) RightJoin(table schema.Tabler, on ...field.Expr) *extratimeDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e extratimeDo) Group(cols ...field.Expr) *extratimeDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e extratimeDo) Having(conds ...gen.Condition) *extratimeDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e extratimeDo) Limit(limit int) *extratimeDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e extratimeDo) Offset(offset int) *extratimeDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e extratimeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *extratimeDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e extratimeDo) Unscoped() *extratimeDo {
	return e.withDO(e.DO.Unscoped())
}

func (e extratimeDo) Create(values ...*model.Extratime) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e extratimeDo) CreateInBatches(values []*model.Extratime, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e extratimeDo) Save(values ...*model.Extratime) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e extratimeDo) First() (*model.Extratime, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Extratime), nil
	}
}

func (e extratimeDo) Take() (*model.Extratime, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Extratime), nil
	}
}

func (e extratimeDo) Last() (*model.Extratime, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Extratime), nil
	}
}

func (e extratimeDo) Find() ([]*model.Extratime, error) {
	result, err := e.DO.Find()
	return result.([]*model.Extratime), err
}

func (e extratimeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Extratime, err error) {
	buf := make([]*model.Extratime, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e extratimeDo) FindInBatches(result *[]*model.Extratime, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e extratimeDo) Attrs(attrs ...field.AssignExpr) *extratimeDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e extratimeDo) Assign(attrs ...field.AssignExpr) *extratimeDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e extratimeDo) Joins(fields ...field.RelationField) *extratimeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e extratimeDo) Preload(fields ...field.RelationField) *extratimeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e extratimeDo) FirstOrInit() (*model.Extratime, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Extratime), nil
	}
}

func (e extratimeDo) FirstOrCreate() (*model.Extratime, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Extratime), nil
	}
}

func (e extratimeDo) FindByPage(offset int, limit int) (result []*model.Extratime, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e extratimeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e extratimeDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e extratimeDo) Delete(models ...*model.Extratime) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *extratimeDo) withDO(do gen.Dao) *extratimeDo {
	e.DO = *do.(*gen.DO)
	return e
}
