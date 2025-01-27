// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/loginrecord"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/predicate"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/user"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoginRecordQuery is the builder for querying LoginRecord entities.
type LoginRecordQuery struct {
	config
	ctx        *QueryContext
	order      []loginrecord.OrderOption
	inters     []Interceptor
	predicates []predicate.LoginRecord
	withUsers  *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LoginRecordQuery builder.
func (lrq *LoginRecordQuery) Where(ps ...predicate.LoginRecord) *LoginRecordQuery {
	lrq.predicates = append(lrq.predicates, ps...)
	return lrq
}

// Limit the number of records to be returned by this query.
func (lrq *LoginRecordQuery) Limit(limit int) *LoginRecordQuery {
	lrq.ctx.Limit = &limit
	return lrq
}

// Offset to start from.
func (lrq *LoginRecordQuery) Offset(offset int) *LoginRecordQuery {
	lrq.ctx.Offset = &offset
	return lrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lrq *LoginRecordQuery) Unique(unique bool) *LoginRecordQuery {
	lrq.ctx.Unique = &unique
	return lrq
}

// Order specifies how the records should be ordered.
func (lrq *LoginRecordQuery) Order(o ...loginrecord.OrderOption) *LoginRecordQuery {
	lrq.order = append(lrq.order, o...)
	return lrq
}

// QueryUsers chains the current query on the "users" edge.
func (lrq *LoginRecordQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: lrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(loginrecord.Table, loginrecord.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, loginrecord.UsersTable, loginrecord.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(lrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LoginRecord entity from the query.
// Returns a *NotFoundError when no LoginRecord was found.
func (lrq *LoginRecordQuery) First(ctx context.Context) (*LoginRecord, error) {
	nodes, err := lrq.Limit(1).All(setContextOp(ctx, lrq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{loginrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lrq *LoginRecordQuery) FirstX(ctx context.Context) *LoginRecord {
	node, err := lrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LoginRecord ID from the query.
// Returns a *NotFoundError when no LoginRecord ID was found.
func (lrq *LoginRecordQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lrq.Limit(1).IDs(setContextOp(ctx, lrq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{loginrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lrq *LoginRecordQuery) FirstIDX(ctx context.Context) int64 {
	id, err := lrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LoginRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LoginRecord entity is found.
// Returns a *NotFoundError when no LoginRecord entities are found.
func (lrq *LoginRecordQuery) Only(ctx context.Context) (*LoginRecord, error) {
	nodes, err := lrq.Limit(2).All(setContextOp(ctx, lrq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{loginrecord.Label}
	default:
		return nil, &NotSingularError{loginrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lrq *LoginRecordQuery) OnlyX(ctx context.Context) *LoginRecord {
	node, err := lrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LoginRecord ID in the query.
// Returns a *NotSingularError when more than one LoginRecord ID is found.
// Returns a *NotFoundError when no entities are found.
func (lrq *LoginRecordQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lrq.Limit(2).IDs(setContextOp(ctx, lrq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{loginrecord.Label}
	default:
		err = &NotSingularError{loginrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lrq *LoginRecordQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := lrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LoginRecords.
func (lrq *LoginRecordQuery) All(ctx context.Context) ([]*LoginRecord, error) {
	ctx = setContextOp(ctx, lrq.ctx, ent.OpQueryAll)
	if err := lrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LoginRecord, *LoginRecordQuery]()
	return withInterceptors[[]*LoginRecord](ctx, lrq, qr, lrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lrq *LoginRecordQuery) AllX(ctx context.Context) []*LoginRecord {
	nodes, err := lrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LoginRecord IDs.
func (lrq *LoginRecordQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if lrq.ctx.Unique == nil && lrq.path != nil {
		lrq.Unique(true)
	}
	ctx = setContextOp(ctx, lrq.ctx, ent.OpQueryIDs)
	if err = lrq.Select(loginrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lrq *LoginRecordQuery) IDsX(ctx context.Context) []int64 {
	ids, err := lrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lrq *LoginRecordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lrq.ctx, ent.OpQueryCount)
	if err := lrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lrq, querierCount[*LoginRecordQuery](), lrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lrq *LoginRecordQuery) CountX(ctx context.Context) int {
	count, err := lrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lrq *LoginRecordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lrq.ctx, ent.OpQueryExist)
	switch _, err := lrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lrq *LoginRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := lrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LoginRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lrq *LoginRecordQuery) Clone() *LoginRecordQuery {
	if lrq == nil {
		return nil
	}
	return &LoginRecordQuery{
		config:     lrq.config,
		ctx:        lrq.ctx.Clone(),
		order:      append([]loginrecord.OrderOption{}, lrq.order...),
		inters:     append([]Interceptor{}, lrq.inters...),
		predicates: append([]predicate.LoginRecord{}, lrq.predicates...),
		withUsers:  lrq.withUsers.Clone(),
		// clone intermediate query.
		sql:  lrq.sql.Clone(),
		path: lrq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (lrq *LoginRecordQuery) WithUsers(opts ...func(*UserQuery)) *LoginRecordQuery {
	query := (&UserClient{config: lrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lrq.withUsers = query
	return lrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LoginRecord.Query().
//		GroupBy(loginrecord.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lrq *LoginRecordQuery) GroupBy(field string, fields ...string) *LoginRecordGroupBy {
	lrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LoginRecordGroupBy{build: lrq}
	grbuild.flds = &lrq.ctx.Fields
	grbuild.label = loginrecord.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.LoginRecord.Query().
//		Select(loginrecord.FieldCreatedAt).
//		Scan(ctx, &v)
func (lrq *LoginRecordQuery) Select(fields ...string) *LoginRecordSelect {
	lrq.ctx.Fields = append(lrq.ctx.Fields, fields...)
	sbuild := &LoginRecordSelect{LoginRecordQuery: lrq}
	sbuild.label = loginrecord.Label
	sbuild.flds, sbuild.scan = &lrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LoginRecordSelect configured with the given aggregations.
func (lrq *LoginRecordQuery) Aggregate(fns ...AggregateFunc) *LoginRecordSelect {
	return lrq.Select().Aggregate(fns...)
}

func (lrq *LoginRecordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lrq); err != nil {
				return err
			}
		}
	}
	for _, f := range lrq.ctx.Fields {
		if !loginrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lrq.path != nil {
		prev, err := lrq.path(ctx)
		if err != nil {
			return err
		}
		lrq.sql = prev
	}
	return nil
}

func (lrq *LoginRecordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LoginRecord, error) {
	var (
		nodes       = []*LoginRecord{}
		withFKs     = lrq.withFKs
		_spec       = lrq.querySpec()
		loadedTypes = [1]bool{
			lrq.withUsers != nil,
		}
	)
	if lrq.withUsers != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, loginrecord.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LoginRecord).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LoginRecord{config: lrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lrq.withUsers; query != nil {
		if err := lrq.loadUsers(ctx, query, nodes, nil,
			func(n *LoginRecord, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lrq *LoginRecordQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*LoginRecord, init func(*LoginRecord), assign func(*LoginRecord, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*LoginRecord)
	for i := range nodes {
		if nodes[i].user_login_records == nil {
			continue
		}
		fk := *nodes[i].user_login_records
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_login_records" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lrq *LoginRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lrq.querySpec()
	_spec.Node.Columns = lrq.ctx.Fields
	if len(lrq.ctx.Fields) > 0 {
		_spec.Unique = lrq.ctx.Unique != nil && *lrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lrq.driver, _spec)
}

func (lrq *LoginRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(loginrecord.Table, loginrecord.Columns, sqlgraph.NewFieldSpec(loginrecord.FieldID, field.TypeInt64))
	_spec.From = lrq.sql
	if unique := lrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lrq.path != nil {
		_spec.Unique = true
	}
	if fields := lrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loginrecord.FieldID)
		for i := range fields {
			if fields[i] != loginrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lrq *LoginRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lrq.driver.Dialect())
	t1 := builder.Table(loginrecord.Table)
	columns := lrq.ctx.Fields
	if len(columns) == 0 {
		columns = loginrecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lrq.sql != nil {
		selector = lrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lrq.ctx.Unique != nil && *lrq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lrq.predicates {
		p(selector)
	}
	for _, p := range lrq.order {
		p(selector)
	}
	if offset := lrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LoginRecordGroupBy is the group-by builder for LoginRecord entities.
type LoginRecordGroupBy struct {
	selector
	build *LoginRecordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lrgb *LoginRecordGroupBy) Aggregate(fns ...AggregateFunc) *LoginRecordGroupBy {
	lrgb.fns = append(lrgb.fns, fns...)
	return lrgb
}

// Scan applies the selector query and scans the result into the given value.
func (lrgb *LoginRecordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lrgb.build.ctx, ent.OpQueryGroupBy)
	if err := lrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoginRecordQuery, *LoginRecordGroupBy](ctx, lrgb.build, lrgb, lrgb.build.inters, v)
}

func (lrgb *LoginRecordGroupBy) sqlScan(ctx context.Context, root *LoginRecordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lrgb.fns))
	for _, fn := range lrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lrgb.flds)+len(lrgb.fns))
		for _, f := range *lrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LoginRecordSelect is the builder for selecting fields of LoginRecord entities.
type LoginRecordSelect struct {
	*LoginRecordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lrs *LoginRecordSelect) Aggregate(fns ...AggregateFunc) *LoginRecordSelect {
	lrs.fns = append(lrs.fns, fns...)
	return lrs
}

// Scan applies the selector query and scans the result into the given value.
func (lrs *LoginRecordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lrs.ctx, ent.OpQuerySelect)
	if err := lrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoginRecordQuery, *LoginRecordSelect](ctx, lrs.LoginRecordQuery, lrs, lrs.inters, v)
}

func (lrs *LoginRecordSelect) sqlScan(ctx context.Context, root *LoginRecordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lrs.fns))
	for _, fn := range lrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
