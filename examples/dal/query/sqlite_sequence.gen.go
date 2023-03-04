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

func newSqliteSequence(db *gorm.DB, opts ...gen.DOOption) sqliteSequence {
	_sqliteSequence := sqliteSequence{}

	_sqliteSequence.sqliteSequenceDo.UseDB(db, opts...)
	_sqliteSequence.sqliteSequenceDo.UseModel(&model.SqliteSequence{})

	tableName := _sqliteSequence.sqliteSequenceDo.TableName()
	_sqliteSequence.ALL = field.NewAsterisk(tableName)
	_sqliteSequence.Name = field.NewString(tableName, "name")
	_sqliteSequence.Seq = field.NewString(tableName, "seq")

	_sqliteSequence.fillFieldMap()

	return _sqliteSequence
}

type sqliteSequence struct {
	sqliteSequenceDo sqliteSequenceDo

	ALL  field.Asterisk
	Name field.String
	Seq  field.String

	fieldMap map[string]field.Expr
}

func (s sqliteSequence) Table(newTableName string) *sqliteSequence {
	s.sqliteSequenceDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sqliteSequence) As(alias string) *sqliteSequence {
	s.sqliteSequenceDo.DO = *(s.sqliteSequenceDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sqliteSequence) updateTableName(table string) *sqliteSequence {
	s.ALL = field.NewAsterisk(table)
	s.Name = field.NewString(table, "name")
	s.Seq = field.NewString(table, "seq")

	s.fillFieldMap()

	return s
}

func (s *sqliteSequence) WithContext(ctx context.Context) *sqliteSequenceDo {
	return s.sqliteSequenceDo.WithContext(ctx)
}

func (s sqliteSequence) TableName() string { return s.sqliteSequenceDo.TableName() }

func (s sqliteSequence) Alias() string { return s.sqliteSequenceDo.Alias() }

func (s *sqliteSequence) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sqliteSequence) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 2)
	s.fieldMap["name"] = s.Name
	s.fieldMap["seq"] = s.Seq
}

func (s sqliteSequence) clone(db *gorm.DB) sqliteSequence {
	s.sqliteSequenceDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sqliteSequence) replaceDB(db *gorm.DB) sqliteSequence {
	s.sqliteSequenceDo.ReplaceDB(db)
	return s
}

type sqliteSequenceDo struct{ gen.DO }

func (s sqliteSequenceDo) Debug() *sqliteSequenceDo {
	return s.withDO(s.DO.Debug())
}

func (s sqliteSequenceDo) WithContext(ctx context.Context) *sqliteSequenceDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sqliteSequenceDo) ReadDB() *sqliteSequenceDo {
	return s.Clauses(dbresolver.Read)
}

func (s sqliteSequenceDo) WriteDB() *sqliteSequenceDo {
	return s.Clauses(dbresolver.Write)
}

func (s sqliteSequenceDo) Session(config *gorm.Session) *sqliteSequenceDo {
	return s.withDO(s.DO.Session(config))
}

func (s sqliteSequenceDo) Clauses(conds ...clause.Expression) *sqliteSequenceDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sqliteSequenceDo) Returning(value interface{}, columns ...string) *sqliteSequenceDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sqliteSequenceDo) Not(conds ...gen.Condition) *sqliteSequenceDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sqliteSequenceDo) Or(conds ...gen.Condition) *sqliteSequenceDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sqliteSequenceDo) Select(conds ...field.Expr) *sqliteSequenceDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sqliteSequenceDo) Where(conds ...gen.Condition) *sqliteSequenceDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sqliteSequenceDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sqliteSequenceDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sqliteSequenceDo) Order(conds ...field.Expr) *sqliteSequenceDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sqliteSequenceDo) Distinct(cols ...field.Expr) *sqliteSequenceDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sqliteSequenceDo) Omit(cols ...field.Expr) *sqliteSequenceDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sqliteSequenceDo) Join(table schema.Tabler, on ...field.Expr) *sqliteSequenceDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sqliteSequenceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sqliteSequenceDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sqliteSequenceDo) RightJoin(table schema.Tabler, on ...field.Expr) *sqliteSequenceDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sqliteSequenceDo) Group(cols ...field.Expr) *sqliteSequenceDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sqliteSequenceDo) Having(conds ...gen.Condition) *sqliteSequenceDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sqliteSequenceDo) Limit(limit int) *sqliteSequenceDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sqliteSequenceDo) Offset(offset int) *sqliteSequenceDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sqliteSequenceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sqliteSequenceDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sqliteSequenceDo) Unscoped() *sqliteSequenceDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sqliteSequenceDo) Create(values ...*model.SqliteSequence) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sqliteSequenceDo) CreateInBatches(values []*model.SqliteSequence, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sqliteSequenceDo) Save(values ...*model.SqliteSequence) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sqliteSequenceDo) First() (*model.SqliteSequence, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqliteSequence), nil
	}
}

func (s sqliteSequenceDo) Take() (*model.SqliteSequence, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqliteSequence), nil
	}
}

func (s sqliteSequenceDo) Last() (*model.SqliteSequence, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqliteSequence), nil
	}
}

func (s sqliteSequenceDo) Find() ([]*model.SqliteSequence, error) {
	result, err := s.DO.Find()
	return result.([]*model.SqliteSequence), err
}

func (s sqliteSequenceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SqliteSequence, err error) {
	buf := make([]*model.SqliteSequence, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sqliteSequenceDo) FindInBatches(result *[]*model.SqliteSequence, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sqliteSequenceDo) Attrs(attrs ...field.AssignExpr) *sqliteSequenceDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sqliteSequenceDo) Assign(attrs ...field.AssignExpr) *sqliteSequenceDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sqliteSequenceDo) Joins(fields ...field.RelationField) *sqliteSequenceDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sqliteSequenceDo) Preload(fields ...field.RelationField) *sqliteSequenceDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sqliteSequenceDo) FirstOrInit() (*model.SqliteSequence, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqliteSequence), nil
	}
}

func (s sqliteSequenceDo) FirstOrCreate() (*model.SqliteSequence, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqliteSequence), nil
	}
}

func (s sqliteSequenceDo) FindByPage(offset int, limit int) (result []*model.SqliteSequence, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sqliteSequenceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sqliteSequenceDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sqliteSequenceDo) Delete(models ...*model.SqliteSequence) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sqliteSequenceDo) withDO(do gen.Dao) *sqliteSequenceDo {
	s.DO = *do.(*gen.DO)
	return s
}
