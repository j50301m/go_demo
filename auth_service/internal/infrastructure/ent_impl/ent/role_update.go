// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/authclient"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/predicate"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/role"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/user"
	"hype-casino-platform/pkg/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RoleUpdate) SetUpdatedAt(t time.Time) *RoleUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetName sets the "name" field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableName(s *string) *RoleUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetPermissions sets the "permissions" field.
func (ru *RoleUpdate) SetPermissions(e []enum.Permission) *RoleUpdate {
	ru.mutation.SetPermissions(e)
	return ru
}

// AppendPermissions appends e to the "permissions" field.
func (ru *RoleUpdate) AppendPermissions(e []enum.Permission) *RoleUpdate {
	ru.mutation.AppendPermissions(e)
	return ru
}

// SetIsSystem sets the "is_system" field.
func (ru *RoleUpdate) SetIsSystem(b bool) *RoleUpdate {
	ru.mutation.SetIsSystem(b)
	return ru
}

// SetNillableIsSystem sets the "is_system" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableIsSystem(b *bool) *RoleUpdate {
	if b != nil {
		ru.SetIsSystem(*b)
	}
	return ru
}

// SetClientType sets the "client_type" field.
func (ru *RoleUpdate) SetClientType(i int) *RoleUpdate {
	ru.mutation.ResetClientType()
	ru.mutation.SetClientType(i)
	return ru
}

// SetNillableClientType sets the "client_type" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableClientType(i *int) *RoleUpdate {
	if i != nil {
		ru.SetClientType(*i)
	}
	return ru
}

// AddClientType adds i to the "client_type" field.
func (ru *RoleUpdate) AddClientType(i int) *RoleUpdate {
	ru.mutation.AddClientType(i)
	return ru
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ru *RoleUpdate) AddUserIDs(ids ...int64) *RoleUpdate {
	ru.mutation.AddUserIDs(ids...)
	return ru
}

// AddUsers adds the "users" edges to the User entity.
func (ru *RoleUpdate) AddUsers(u ...*User) *RoleUpdate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.AddUserIDs(ids...)
}

// AddAuthClientIDs adds the "auth_clients" edge to the AuthClient entity by IDs.
func (ru *RoleUpdate) AddAuthClientIDs(ids ...int64) *RoleUpdate {
	ru.mutation.AddAuthClientIDs(ids...)
	return ru
}

// AddAuthClients adds the "auth_clients" edges to the AuthClient entity.
func (ru *RoleUpdate) AddAuthClients(a ...*AuthClient) *RoleUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.AddAuthClientIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (ru *RoleUpdate) ClearUsers() *RoleUpdate {
	ru.mutation.ClearUsers()
	return ru
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (ru *RoleUpdate) RemoveUserIDs(ids ...int64) *RoleUpdate {
	ru.mutation.RemoveUserIDs(ids...)
	return ru
}

// RemoveUsers removes "users" edges to User entities.
func (ru *RoleUpdate) RemoveUsers(u ...*User) *RoleUpdate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.RemoveUserIDs(ids...)
}

// ClearAuthClients clears all "auth_clients" edges to the AuthClient entity.
func (ru *RoleUpdate) ClearAuthClients() *RoleUpdate {
	ru.mutation.ClearAuthClients()
	return ru
}

// RemoveAuthClientIDs removes the "auth_clients" edge to AuthClient entities by IDs.
func (ru *RoleUpdate) RemoveAuthClientIDs(ids ...int64) *RoleUpdate {
	ru.mutation.RemoveAuthClientIDs(ids...)
	return ru
}

// RemoveAuthClients removes "auth_clients" edges to AuthClient entities.
func (ru *RoleUpdate) RemoveAuthClients(a ...*AuthClient) *RoleUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.RemoveAuthClientIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RoleUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := role.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Permissions(); ok {
		_spec.SetField(role.FieldPermissions, field.TypeJSON, value)
	}
	if value, ok := ru.mutation.AppendedPermissions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, role.FieldPermissions, value)
		})
	}
	if value, ok := ru.mutation.IsSystem(); ok {
		_spec.SetField(role.FieldIsSystem, field.TypeBool, value)
	}
	if value, ok := ru.mutation.ClientType(); ok {
		_spec.SetField(role.FieldClientType, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedClientType(); ok {
		_spec.AddField(role.FieldClientType, field.TypeInt, value)
	}
	if ru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: []string{role.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedUsersIDs(); len(nodes) > 0 && !ru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: []string{role.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: []string{role.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.AuthClientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AuthClientsTable,
			Columns: role.AuthClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authclient.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedAuthClientsIDs(); len(nodes) > 0 && !ru.mutation.AuthClientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AuthClientsTable,
			Columns: role.AuthClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authclient.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.AuthClientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AuthClientsTable,
			Columns: role.AuthClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authclient.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RoleUpdateOne) SetUpdatedAt(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetName sets the "name" field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableName(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetPermissions sets the "permissions" field.
func (ruo *RoleUpdateOne) SetPermissions(e []enum.Permission) *RoleUpdateOne {
	ruo.mutation.SetPermissions(e)
	return ruo
}

// AppendPermissions appends e to the "permissions" field.
func (ruo *RoleUpdateOne) AppendPermissions(e []enum.Permission) *RoleUpdateOne {
	ruo.mutation.AppendPermissions(e)
	return ruo
}

// SetIsSystem sets the "is_system" field.
func (ruo *RoleUpdateOne) SetIsSystem(b bool) *RoleUpdateOne {
	ruo.mutation.SetIsSystem(b)
	return ruo
}

// SetNillableIsSystem sets the "is_system" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableIsSystem(b *bool) *RoleUpdateOne {
	if b != nil {
		ruo.SetIsSystem(*b)
	}
	return ruo
}

// SetClientType sets the "client_type" field.
func (ruo *RoleUpdateOne) SetClientType(i int) *RoleUpdateOne {
	ruo.mutation.ResetClientType()
	ruo.mutation.SetClientType(i)
	return ruo
}

// SetNillableClientType sets the "client_type" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableClientType(i *int) *RoleUpdateOne {
	if i != nil {
		ruo.SetClientType(*i)
	}
	return ruo
}

// AddClientType adds i to the "client_type" field.
func (ruo *RoleUpdateOne) AddClientType(i int) *RoleUpdateOne {
	ruo.mutation.AddClientType(i)
	return ruo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ruo *RoleUpdateOne) AddUserIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.AddUserIDs(ids...)
	return ruo
}

// AddUsers adds the "users" edges to the User entity.
func (ruo *RoleUpdateOne) AddUsers(u ...*User) *RoleUpdateOne {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.AddUserIDs(ids...)
}

// AddAuthClientIDs adds the "auth_clients" edge to the AuthClient entity by IDs.
func (ruo *RoleUpdateOne) AddAuthClientIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.AddAuthClientIDs(ids...)
	return ruo
}

// AddAuthClients adds the "auth_clients" edges to the AuthClient entity.
func (ruo *RoleUpdateOne) AddAuthClients(a ...*AuthClient) *RoleUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.AddAuthClientIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (ruo *RoleUpdateOne) ClearUsers() *RoleUpdateOne {
	ruo.mutation.ClearUsers()
	return ruo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (ruo *RoleUpdateOne) RemoveUserIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.RemoveUserIDs(ids...)
	return ruo
}

// RemoveUsers removes "users" edges to User entities.
func (ruo *RoleUpdateOne) RemoveUsers(u ...*User) *RoleUpdateOne {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.RemoveUserIDs(ids...)
}

// ClearAuthClients clears all "auth_clients" edges to the AuthClient entity.
func (ruo *RoleUpdateOne) ClearAuthClients() *RoleUpdateOne {
	ruo.mutation.ClearAuthClients()
	return ruo
}

// RemoveAuthClientIDs removes the "auth_clients" edge to AuthClient entities by IDs.
func (ruo *RoleUpdateOne) RemoveAuthClientIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.RemoveAuthClientIDs(ids...)
	return ruo
}

// RemoveAuthClients removes "auth_clients" edges to AuthClient entities.
func (ruo *RoleUpdateOne) RemoveAuthClients(a ...*AuthClient) *RoleUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.RemoveAuthClientIDs(ids...)
}

// Where appends a list predicates to the RoleUpdate builder.
func (ruo *RoleUpdateOne) Where(ps ...predicate.Role) *RoleUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoleUpdateOne) Select(field string, fields ...string) *RoleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RoleUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := role.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Role.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for _, f := range fields {
			if !role.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Permissions(); ok {
		_spec.SetField(role.FieldPermissions, field.TypeJSON, value)
	}
	if value, ok := ruo.mutation.AppendedPermissions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, role.FieldPermissions, value)
		})
	}
	if value, ok := ruo.mutation.IsSystem(); ok {
		_spec.SetField(role.FieldIsSystem, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.ClientType(); ok {
		_spec.SetField(role.FieldClientType, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedClientType(); ok {
		_spec.AddField(role.FieldClientType, field.TypeInt, value)
	}
	if ruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: []string{role.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !ruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: []string{role.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: []string{role.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.AuthClientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AuthClientsTable,
			Columns: role.AuthClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authclient.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedAuthClientsIDs(); len(nodes) > 0 && !ruo.mutation.AuthClientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AuthClientsTable,
			Columns: role.AuthClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authclient.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.AuthClientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AuthClientsTable,
			Columns: role.AuthClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authclient.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
