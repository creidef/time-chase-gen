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

func newProjectactivity(db *gorm.DB, opts ...gen.DOOption) projectactivity {
	_projectactivity := projectactivity{}

	_projectactivity.projectactivityDo.UseDB(db, opts...)
	_projectactivity.projectactivityDo.UseModel(&model.Projectactivity{})

	tableName := _projectactivity.projectactivityDo.TableName()
	_projectactivity.ALL = field.NewAsterisk(tableName)
	_projectactivity.ID = field.NewInt32(tableName, "id")
	_projectactivity.Project = field.NewInt32(tableName, "project")
	_projectactivity.Activity = field.NewInt32(tableName, "activity")

	_projectactivity.fillFieldMap()

	return _projectactivity
}

type projectactivity struct {
	projectactivityDo projectactivityDo

	ALL      field.Asterisk
	ID       field.Int32
	Project  field.Int32
	Activity field.Int32

	fieldMap map[string]field.Expr
}

func (p projectactivity) Table(newTableName string) *projectactivity {
	p.projectactivityDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p projectactivity) As(alias string) *projectactivity {
	p.projectactivityDo.DO = *(p.projectactivityDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *projectactivity) updateTableName(table string) *projectactivity {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.Project = field.NewInt32(table, "project")
	p.Activity = field.NewInt32(table, "activity")

	p.fillFieldMap()

	return p
}

func (p *projectactivity) WithContext(ctx context.Context) *projectactivityDo {
	return p.projectactivityDo.WithContext(ctx)
}

func (p projectactivity) TableName() string { return p.projectactivityDo.TableName() }

func (p projectactivity) Alias() string { return p.projectactivityDo.Alias() }

func (p *projectactivity) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *projectactivity) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["id"] = p.ID
	p.fieldMap["project"] = p.Project
	p.fieldMap["activity"] = p.Activity
}

func (p projectactivity) clone(db *gorm.DB) projectactivity {
	p.projectactivityDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p projectactivity) replaceDB(db *gorm.DB) projectactivity {
	p.projectactivityDo.ReplaceDB(db)
	return p
}

type projectactivityDo struct{ gen.DO }

func (p projectactivityDo) Debug() *projectactivityDo {
	return p.withDO(p.DO.Debug())
}

func (p projectactivityDo) WithContext(ctx context.Context) *projectactivityDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p projectactivityDo) ReadDB() *projectactivityDo {
	return p.Clauses(dbresolver.Read)
}

func (p projectactivityDo) WriteDB() *projectactivityDo {
	return p.Clauses(dbresolver.Write)
}

func (p projectactivityDo) Session(config *gorm.Session) *projectactivityDo {
	return p.withDO(p.DO.Session(config))
}

func (p projectactivityDo) Clauses(conds ...clause.Expression) *projectactivityDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p projectactivityDo) Returning(value interface{}, columns ...string) *projectactivityDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p projectactivityDo) Not(conds ...gen.Condition) *projectactivityDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p projectactivityDo) Or(conds ...gen.Condition) *projectactivityDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p projectactivityDo) Select(conds ...field.Expr) *projectactivityDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p projectactivityDo) Where(conds ...gen.Condition) *projectactivityDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p projectactivityDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *projectactivityDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p projectactivityDo) Order(conds ...field.Expr) *projectactivityDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p projectactivityDo) Distinct(cols ...field.Expr) *projectactivityDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p projectactivityDo) Omit(cols ...field.Expr) *projectactivityDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p projectactivityDo) Join(table schema.Tabler, on ...field.Expr) *projectactivityDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p projectactivityDo) LeftJoin(table schema.Tabler, on ...field.Expr) *projectactivityDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p projectactivityDo) RightJoin(table schema.Tabler, on ...field.Expr) *projectactivityDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p projectactivityDo) Group(cols ...field.Expr) *projectactivityDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p projectactivityDo) Having(conds ...gen.Condition) *projectactivityDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p projectactivityDo) Limit(limit int) *projectactivityDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p projectactivityDo) Offset(offset int) *projectactivityDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p projectactivityDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *projectactivityDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p projectactivityDo) Unscoped() *projectactivityDo {
	return p.withDO(p.DO.Unscoped())
}

func (p projectactivityDo) Create(values ...*model.Projectactivity) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p projectactivityDo) CreateInBatches(values []*model.Projectactivity, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p projectactivityDo) Save(values ...*model.Projectactivity) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p projectactivityDo) First() (*model.Projectactivity, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Projectactivity), nil
	}
}

func (p projectactivityDo) Take() (*model.Projectactivity, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Projectactivity), nil
	}
}

func (p projectactivityDo) Last() (*model.Projectactivity, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Projectactivity), nil
	}
}

func (p projectactivityDo) Find() ([]*model.Projectactivity, error) {
	result, err := p.DO.Find()
	return result.([]*model.Projectactivity), err
}

func (p projectactivityDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Projectactivity, err error) {
	buf := make([]*model.Projectactivity, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p projectactivityDo) FindInBatches(result *[]*model.Projectactivity, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p projectactivityDo) Attrs(attrs ...field.AssignExpr) *projectactivityDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p projectactivityDo) Assign(attrs ...field.AssignExpr) *projectactivityDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p projectactivityDo) Joins(fields ...field.RelationField) *projectactivityDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p projectactivityDo) Preload(fields ...field.RelationField) *projectactivityDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p projectactivityDo) FirstOrInit() (*model.Projectactivity, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Projectactivity), nil
	}
}

func (p projectactivityDo) FirstOrCreate() (*model.Projectactivity, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Projectactivity), nil
	}
}

func (p projectactivityDo) FindByPage(offset int, limit int) (result []*model.Projectactivity, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p projectactivityDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p projectactivityDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p projectactivityDo) Delete(models ...*model.Projectactivity) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *projectactivityDo) withDO(do gen.Dao) *projectactivityDo {
	p.DO = *do.(*gen.DO)
	return p
}
