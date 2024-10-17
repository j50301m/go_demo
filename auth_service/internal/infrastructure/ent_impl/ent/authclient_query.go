// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/authclient"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/predicate"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/role"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/user"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthClientQuery is the builder for querying AuthClient entities.
type AuthClientQuery struct {
	config
	ctx        *QueryContext
	order      []authclient.OrderOption
	inters     []Interceptor
	predicates []predicate.AuthClient
	withUsers  *UserQuery
	withRoles  *RoleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthClientQuery builder.
func (acq *AuthClientQuery) Where(ps ...predicate.AuthClient) *AuthClientQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit the number of records to be returned by this query.
func (acq *AuthClientQuery) Limit(limit int) *AuthClientQuery {
	acq.ctx.Limit = &limit
	return acq
}

// Offset to start from.
func (acq *AuthClientQuery) Offset(offset int) *AuthClientQuery {
	acq.ctx.Offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *AuthClientQuery) Unique(unique bool) *AuthClientQuery {
	acq.ctx.Unique = &unique
	return acq
}

// Order specifies how the records should be ordered.
func (acq *AuthClientQuery) Order(o ...authclient.OrderOption) *AuthClientQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// QueryUsers chains the current query on the "users" edge.
func (acq *AuthClientQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(authclient.Table, authclient.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, authclient.UsersTable, authclient.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoles chains the current query on the "roles" edge.
func (acq *AuthClientQuery) QueryRoles() *RoleQuery {
	query := (&RoleClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(authclient.Table, authclient.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, authclient.RolesTable, authclient.RolesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AuthClient entity from the query.
// Returns a *NotFoundError when no AuthClient was found.
func (acq *AuthClientQuery) First(ctx context.Context) (*AuthClient, error) {
	nodes, err := acq.Limit(1).All(setContextOp(ctx, acq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authclient.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *AuthClientQuery) FirstX(ctx context.Context) *AuthClient {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthClient ID from the query.
// Returns a *NotFoundError when no AuthClient ID was found.
func (acq *AuthClientQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = acq.Limit(1).IDs(setContextOp(ctx, acq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authclient.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *AuthClientQuery) FirstIDX(ctx context.Context) int64 {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthClient entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AuthClient entity is found.
// Returns a *NotFoundError when no AuthClient entities are found.
func (acq *AuthClientQuery) Only(ctx context.Context) (*AuthClient, error) {
	nodes, err := acq.Limit(2).All(setContextOp(ctx, acq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authclient.Label}
	default:
		return nil, &NotSingularError{authclient.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *AuthClientQuery) OnlyX(ctx context.Context) *AuthClient {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthClient ID in the query.
// Returns a *NotSingularError when more than one AuthClient ID is found.
// Returns a *NotFoundError when no entities are found.
func (acq *AuthClientQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = acq.Limit(2).IDs(setContextOp(ctx, acq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authclient.Label}
	default:
		err = &NotSingularError{authclient.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *AuthClientQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthClients.
func (acq *AuthClientQuery) All(ctx context.Context) ([]*AuthClient, error) {
	ctx = setContextOp(ctx, acq.ctx, ent.OpQueryAll)
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AuthClient, *AuthClientQuery]()
	return withInterceptors[[]*AuthClient](ctx, acq, qr, acq.inters)
}

// AllX is like All, but panics if an error occurs.
func (acq *AuthClientQuery) AllX(ctx context.Context) []*AuthClient {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthClient IDs.
func (acq *AuthClientQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if acq.ctx.Unique == nil && acq.path != nil {
		acq.Unique(true)
	}
	ctx = setContextOp(ctx, acq.ctx, ent.OpQueryIDs)
	if err = acq.Select(authclient.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *AuthClientQuery) IDsX(ctx context.Context) []int64 {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *AuthClientQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, acq.ctx, ent.OpQueryCount)
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, acq, querierCount[*AuthClientQuery](), acq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (acq *AuthClientQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *AuthClientQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, acq.ctx, ent.OpQueryExist)
	switch _, err := acq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *AuthClientQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthClientQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *AuthClientQuery) Clone() *AuthClientQuery {
	if acq == nil {
		return nil
	}
	return &AuthClientQuery{
		config:     acq.config,
		ctx:        acq.ctx.Clone(),
		order:      append([]authclient.OrderOption{}, acq.order...),
		inters:     append([]Interceptor{}, acq.inters...),
		predicates: append([]predicate.AuthClient{}, acq.predicates...),
		withUsers:  acq.withUsers.Clone(),
		withRoles:  acq.withRoles.Clone(),
		// clone intermediate query.
		sql:  acq.sql.Clone(),
		path: acq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AuthClientQuery) WithUsers(opts ...func(*UserQuery)) *AuthClientQuery {
	query := (&UserClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withUsers = query
	return acq
}

// WithRoles tells the query-builder to eager-load the nodes that are connected to
// the "roles" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AuthClientQuery) WithRoles(opts ...func(*RoleQuery)) *AuthClientQuery {
	query := (&RoleClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withRoles = query
	return acq
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
//	client.AuthClient.Query().
//		GroupBy(authclient.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (acq *AuthClientQuery) GroupBy(field string, fields ...string) *AuthClientGroupBy {
	acq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AuthClientGroupBy{build: acq}
	grbuild.flds = &acq.ctx.Fields
	grbuild.label = authclient.Label
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
//	client.AuthClient.Query().
//		Select(authclient.FieldCreatedAt).
//		Scan(ctx, &v)
func (acq *AuthClientQuery) Select(fields ...string) *AuthClientSelect {
	acq.ctx.Fields = append(acq.ctx.Fields, fields...)
	sbuild := &AuthClientSelect{AuthClientQuery: acq}
	sbuild.label = authclient.Label
	sbuild.flds, sbuild.scan = &acq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AuthClientSelect configured with the given aggregations.
func (acq *AuthClientQuery) Aggregate(fns ...AggregateFunc) *AuthClientSelect {
	return acq.Select().Aggregate(fns...)
}

func (acq *AuthClientQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range acq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, acq); err != nil {
				return err
			}
		}
	}
	for _, f := range acq.ctx.Fields {
		if !authclient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *AuthClientQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AuthClient, error) {
	var (
		nodes       = []*AuthClient{}
		_spec       = acq.querySpec()
		loadedTypes = [2]bool{
			acq.withUsers != nil,
			acq.withRoles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AuthClient).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AuthClient{config: acq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := acq.withUsers; query != nil {
		if err := acq.loadUsers(ctx, query, nodes,
			func(n *AuthClient) { n.Edges.Users = []*User{} },
			func(n *AuthClient, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	if query := acq.withRoles; query != nil {
		if err := acq.loadRoles(ctx, query, nodes,
			func(n *AuthClient) { n.Edges.Roles = []*Role{} },
			func(n *AuthClient, e *Role) { n.Edges.Roles = append(n.Edges.Roles, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (acq *AuthClientQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*AuthClient, init func(*AuthClient), assign func(*AuthClient, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*AuthClient)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(authclient.UsersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.auth_client_users
		if fk == nil {
			return fmt.Errorf(`foreign-key "auth_client_users" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "auth_client_users" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (acq *AuthClientQuery) loadRoles(ctx context.Context, query *RoleQuery, nodes []*AuthClient, init func(*AuthClient), assign func(*AuthClient, *Role)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*AuthClient)
	nids := make(map[int64]map[*AuthClient]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(authclient.RolesTable)
		s.Join(joinT).On(s.C(role.FieldID), joinT.C(authclient.RolesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(authclient.RolesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(authclient.RolesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullInt64).Int64
				inValue := values[1].(*sql.NullInt64).Int64
				if nids[inValue] == nil {
					nids[inValue] = map[*AuthClient]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Role](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "roles" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (acq *AuthClientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	_spec.Node.Columns = acq.ctx.Fields
	if len(acq.ctx.Fields) > 0 {
		_spec.Unique = acq.ctx.Unique != nil && *acq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *AuthClientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(authclient.Table, authclient.Columns, sqlgraph.NewFieldSpec(authclient.FieldID, field.TypeInt64))
	_spec.From = acq.sql
	if unique := acq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if acq.path != nil {
		_spec.Unique = true
	}
	if fields := acq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authclient.FieldID)
		for i := range fields {
			if fields[i] != authclient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *AuthClientQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(authclient.Table)
	columns := acq.ctx.Fields
	if len(columns) == 0 {
		columns = authclient.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.ctx.Unique != nil && *acq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthClientGroupBy is the group-by builder for AuthClient entities.
type AuthClientGroupBy struct {
	selector
	build *AuthClientQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *AuthClientGroupBy) Aggregate(fns ...AggregateFunc) *AuthClientGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the selector query and scans the result into the given value.
func (acgb *AuthClientGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acgb.build.ctx, ent.OpQueryGroupBy)
	if err := acgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthClientQuery, *AuthClientGroupBy](ctx, acgb.build, acgb, acgb.build.inters, v)
}

func (acgb *AuthClientGroupBy) sqlScan(ctx context.Context, root *AuthClientQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*acgb.flds)+len(acgb.fns))
		for _, f := range *acgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*acgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AuthClientSelect is the builder for selecting fields of AuthClient entities.
type AuthClientSelect struct {
	*AuthClientQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (acs *AuthClientSelect) Aggregate(fns ...AggregateFunc) *AuthClientSelect {
	acs.fns = append(acs.fns, fns...)
	return acs
}

// Scan applies the selector query and scans the result into the given value.
func (acs *AuthClientSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acs.ctx, ent.OpQuerySelect)
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthClientQuery, *AuthClientSelect](ctx, acs.AuthClientQuery, acs, acs.inters, v)
}

func (acs *AuthClientSelect) sqlScan(ctx context.Context, root *AuthClientQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(acs.fns))
	for _, fn := range acs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*acs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
