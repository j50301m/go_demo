// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/authclient"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/role"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthClientCreate is the builder for creating a AuthClient entity.
type AuthClientCreate struct {
	config
	mutation *AuthClientMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (acc *AuthClientCreate) SetCreatedAt(t time.Time) *AuthClientCreate {
	acc.mutation.SetCreatedAt(t)
	return acc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acc *AuthClientCreate) SetNillableCreatedAt(t *time.Time) *AuthClientCreate {
	if t != nil {
		acc.SetCreatedAt(*t)
	}
	return acc
}

// SetUpdatedAt sets the "updated_at" field.
func (acc *AuthClientCreate) SetUpdatedAt(t time.Time) *AuthClientCreate {
	acc.mutation.SetUpdatedAt(t)
	return acc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (acc *AuthClientCreate) SetNillableUpdatedAt(t *time.Time) *AuthClientCreate {
	if t != nil {
		acc.SetUpdatedAt(*t)
	}
	return acc
}

// SetClientType sets the "client_type" field.
func (acc *AuthClientCreate) SetClientType(i int) *AuthClientCreate {
	acc.mutation.SetClientType(i)
	return acc
}

// SetMerchantID sets the "merchant_id" field.
func (acc *AuthClientCreate) SetMerchantID(i int64) *AuthClientCreate {
	acc.mutation.SetMerchantID(i)
	return acc
}

// SetSecret sets the "secret" field.
func (acc *AuthClientCreate) SetSecret(s string) *AuthClientCreate {
	acc.mutation.SetSecret(s)
	return acc
}

// SetActive sets the "active" field.
func (acc *AuthClientCreate) SetActive(b bool) *AuthClientCreate {
	acc.mutation.SetActive(b)
	return acc
}

// SetTokenExpireSecs sets the "token_expire_secs" field.
func (acc *AuthClientCreate) SetTokenExpireSecs(i int) *AuthClientCreate {
	acc.mutation.SetTokenExpireSecs(i)
	return acc
}

// SetLoginFailedTimes sets the "login_failed_times" field.
func (acc *AuthClientCreate) SetLoginFailedTimes(i int) *AuthClientCreate {
	acc.mutation.SetLoginFailedTimes(i)
	return acc
}

// SetID sets the "id" field.
func (acc *AuthClientCreate) SetID(i int64) *AuthClientCreate {
	acc.mutation.SetID(i)
	return acc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (acc *AuthClientCreate) AddUserIDs(ids ...int64) *AuthClientCreate {
	acc.mutation.AddUserIDs(ids...)
	return acc
}

// AddUsers adds the "users" edges to the User entity.
func (acc *AuthClientCreate) AddUsers(u ...*User) *AuthClientCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return acc.AddUserIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (acc *AuthClientCreate) AddRoleIDs(ids ...int64) *AuthClientCreate {
	acc.mutation.AddRoleIDs(ids...)
	return acc
}

// AddRoles adds the "roles" edges to the Role entity.
func (acc *AuthClientCreate) AddRoles(r ...*Role) *AuthClientCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return acc.AddRoleIDs(ids...)
}

// Mutation returns the AuthClientMutation object of the builder.
func (acc *AuthClientCreate) Mutation() *AuthClientMutation {
	return acc.mutation
}

// Save creates the AuthClient in the database.
func (acc *AuthClientCreate) Save(ctx context.Context) (*AuthClient, error) {
	acc.defaults()
	return withHooks(ctx, acc.sqlSave, acc.mutation, acc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (acc *AuthClientCreate) SaveX(ctx context.Context) *AuthClient {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *AuthClientCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *AuthClientCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acc *AuthClientCreate) defaults() {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		v := authclient.DefaultCreatedAt()
		acc.mutation.SetCreatedAt(v)
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		v := authclient.DefaultUpdatedAt()
		acc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acc *AuthClientCreate) check() error {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AuthClient.created_at"`)}
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AuthClient.updated_at"`)}
	}
	if _, ok := acc.mutation.ClientType(); !ok {
		return &ValidationError{Name: "client_type", err: errors.New(`ent: missing required field "AuthClient.client_type"`)}
	}
	if _, ok := acc.mutation.MerchantID(); !ok {
		return &ValidationError{Name: "merchant_id", err: errors.New(`ent: missing required field "AuthClient.merchant_id"`)}
	}
	if _, ok := acc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`ent: missing required field "AuthClient.secret"`)}
	}
	if _, ok := acc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "AuthClient.active"`)}
	}
	if _, ok := acc.mutation.TokenExpireSecs(); !ok {
		return &ValidationError{Name: "token_expire_secs", err: errors.New(`ent: missing required field "AuthClient.token_expire_secs"`)}
	}
	if _, ok := acc.mutation.LoginFailedTimes(); !ok {
		return &ValidationError{Name: "login_failed_times", err: errors.New(`ent: missing required field "AuthClient.login_failed_times"`)}
	}
	return nil
}

func (acc *AuthClientCreate) sqlSave(ctx context.Context) (*AuthClient, error) {
	if err := acc.check(); err != nil {
		return nil, err
	}
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	acc.mutation.id = &_node.ID
	acc.mutation.done = true
	return _node, nil
}

func (acc *AuthClientCreate) createSpec() (*AuthClient, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthClient{config: acc.config}
		_spec = sqlgraph.NewCreateSpec(authclient.Table, sqlgraph.NewFieldSpec(authclient.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = acc.conflict
	if id, ok := acc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := acc.mutation.CreatedAt(); ok {
		_spec.SetField(authclient.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := acc.mutation.UpdatedAt(); ok {
		_spec.SetField(authclient.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := acc.mutation.ClientType(); ok {
		_spec.SetField(authclient.FieldClientType, field.TypeInt, value)
		_node.ClientType = value
	}
	if value, ok := acc.mutation.MerchantID(); ok {
		_spec.SetField(authclient.FieldMerchantID, field.TypeInt64, value)
		_node.MerchantID = value
	}
	if value, ok := acc.mutation.Secret(); ok {
		_spec.SetField(authclient.FieldSecret, field.TypeString, value)
		_node.Secret = value
	}
	if value, ok := acc.mutation.Active(); ok {
		_spec.SetField(authclient.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := acc.mutation.TokenExpireSecs(); ok {
		_spec.SetField(authclient.FieldTokenExpireSecs, field.TypeInt, value)
		_node.TokenExpireSecs = value
	}
	if value, ok := acc.mutation.LoginFailedTimes(); ok {
		_spec.SetField(authclient.FieldLoginFailedTimes, field.TypeInt, value)
		_node.LoginFailedTimes = value
	}
	if nodes := acc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   authclient.UsersTable,
			Columns: []string{authclient.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := acc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   authclient.RolesTable,
			Columns: authclient.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AuthClient.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthClientUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (acc *AuthClientCreate) OnConflict(opts ...sql.ConflictOption) *AuthClientUpsertOne {
	acc.conflict = opts
	return &AuthClientUpsertOne{
		create: acc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AuthClient.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acc *AuthClientCreate) OnConflictColumns(columns ...string) *AuthClientUpsertOne {
	acc.conflict = append(acc.conflict, sql.ConflictColumns(columns...))
	return &AuthClientUpsertOne{
		create: acc,
	}
}

type (
	// AuthClientUpsertOne is the builder for "upsert"-ing
	//  one AuthClient node.
	AuthClientUpsertOne struct {
		create *AuthClientCreate
	}

	// AuthClientUpsert is the "OnConflict" setter.
	AuthClientUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthClientUpsert) SetUpdatedAt(v time.Time) *AuthClientUpsert {
	u.Set(authclient.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthClientUpsert) UpdateUpdatedAt() *AuthClientUpsert {
	u.SetExcluded(authclient.FieldUpdatedAt)
	return u
}

// SetClientType sets the "client_type" field.
func (u *AuthClientUpsert) SetClientType(v int) *AuthClientUpsert {
	u.Set(authclient.FieldClientType, v)
	return u
}

// UpdateClientType sets the "client_type" field to the value that was provided on create.
func (u *AuthClientUpsert) UpdateClientType() *AuthClientUpsert {
	u.SetExcluded(authclient.FieldClientType)
	return u
}

// AddClientType adds v to the "client_type" field.
func (u *AuthClientUpsert) AddClientType(v int) *AuthClientUpsert {
	u.Add(authclient.FieldClientType, v)
	return u
}

// SetMerchantID sets the "merchant_id" field.
func (u *AuthClientUpsert) SetMerchantID(v int64) *AuthClientUpsert {
	u.Set(authclient.FieldMerchantID, v)
	return u
}

// UpdateMerchantID sets the "merchant_id" field to the value that was provided on create.
func (u *AuthClientUpsert) UpdateMerchantID() *AuthClientUpsert {
	u.SetExcluded(authclient.FieldMerchantID)
	return u
}

// AddMerchantID adds v to the "merchant_id" field.
func (u *AuthClientUpsert) AddMerchantID(v int64) *AuthClientUpsert {
	u.Add(authclient.FieldMerchantID, v)
	return u
}

// SetSecret sets the "secret" field.
func (u *AuthClientUpsert) SetSecret(v string) *AuthClientUpsert {
	u.Set(authclient.FieldSecret, v)
	return u
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *AuthClientUpsert) UpdateSecret() *AuthClientUpsert {
	u.SetExcluded(authclient.FieldSecret)
	return u
}

// SetActive sets the "active" field.
func (u *AuthClientUpsert) SetActive(v bool) *AuthClientUpsert {
	u.Set(authclient.FieldActive, v)
	return u
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *AuthClientUpsert) UpdateActive() *AuthClientUpsert {
	u.SetExcluded(authclient.FieldActive)
	return u
}

// SetTokenExpireSecs sets the "token_expire_secs" field.
func (u *AuthClientUpsert) SetTokenExpireSecs(v int) *AuthClientUpsert {
	u.Set(authclient.FieldTokenExpireSecs, v)
	return u
}

// UpdateTokenExpireSecs sets the "token_expire_secs" field to the value that was provided on create.
func (u *AuthClientUpsert) UpdateTokenExpireSecs() *AuthClientUpsert {
	u.SetExcluded(authclient.FieldTokenExpireSecs)
	return u
}

// AddTokenExpireSecs adds v to the "token_expire_secs" field.
func (u *AuthClientUpsert) AddTokenExpireSecs(v int) *AuthClientUpsert {
	u.Add(authclient.FieldTokenExpireSecs, v)
	return u
}

// SetLoginFailedTimes sets the "login_failed_times" field.
func (u *AuthClientUpsert) SetLoginFailedTimes(v int) *AuthClientUpsert {
	u.Set(authclient.FieldLoginFailedTimes, v)
	return u
}

// UpdateLoginFailedTimes sets the "login_failed_times" field to the value that was provided on create.
func (u *AuthClientUpsert) UpdateLoginFailedTimes() *AuthClientUpsert {
	u.SetExcluded(authclient.FieldLoginFailedTimes)
	return u
}

// AddLoginFailedTimes adds v to the "login_failed_times" field.
func (u *AuthClientUpsert) AddLoginFailedTimes(v int) *AuthClientUpsert {
	u.Add(authclient.FieldLoginFailedTimes, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AuthClient.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(authclient.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuthClientUpsertOne) UpdateNewValues() *AuthClientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(authclient.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(authclient.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AuthClient.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AuthClientUpsertOne) Ignore() *AuthClientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthClientUpsertOne) DoNothing() *AuthClientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthClientCreate.OnConflict
// documentation for more info.
func (u *AuthClientUpsertOne) Update(set func(*AuthClientUpsert)) *AuthClientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthClientUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthClientUpsertOne) SetUpdatedAt(v time.Time) *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthClientUpsertOne) UpdateUpdatedAt() *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetClientType sets the "client_type" field.
func (u *AuthClientUpsertOne) SetClientType(v int) *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetClientType(v)
	})
}

// AddClientType adds v to the "client_type" field.
func (u *AuthClientUpsertOne) AddClientType(v int) *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.AddClientType(v)
	})
}

// UpdateClientType sets the "client_type" field to the value that was provided on create.
func (u *AuthClientUpsertOne) UpdateClientType() *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateClientType()
	})
}

// SetMerchantID sets the "merchant_id" field.
func (u *AuthClientUpsertOne) SetMerchantID(v int64) *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetMerchantID(v)
	})
}

// AddMerchantID adds v to the "merchant_id" field.
func (u *AuthClientUpsertOne) AddMerchantID(v int64) *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.AddMerchantID(v)
	})
}

// UpdateMerchantID sets the "merchant_id" field to the value that was provided on create.
func (u *AuthClientUpsertOne) UpdateMerchantID() *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateMerchantID()
	})
}

// SetSecret sets the "secret" field.
func (u *AuthClientUpsertOne) SetSecret(v string) *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetSecret(v)
	})
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *AuthClientUpsertOne) UpdateSecret() *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateSecret()
	})
}

// SetActive sets the "active" field.
func (u *AuthClientUpsertOne) SetActive(v bool) *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *AuthClientUpsertOne) UpdateActive() *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateActive()
	})
}

// SetTokenExpireSecs sets the "token_expire_secs" field.
func (u *AuthClientUpsertOne) SetTokenExpireSecs(v int) *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetTokenExpireSecs(v)
	})
}

// AddTokenExpireSecs adds v to the "token_expire_secs" field.
func (u *AuthClientUpsertOne) AddTokenExpireSecs(v int) *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.AddTokenExpireSecs(v)
	})
}

// UpdateTokenExpireSecs sets the "token_expire_secs" field to the value that was provided on create.
func (u *AuthClientUpsertOne) UpdateTokenExpireSecs() *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateTokenExpireSecs()
	})
}

// SetLoginFailedTimes sets the "login_failed_times" field.
func (u *AuthClientUpsertOne) SetLoginFailedTimes(v int) *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetLoginFailedTimes(v)
	})
}

// AddLoginFailedTimes adds v to the "login_failed_times" field.
func (u *AuthClientUpsertOne) AddLoginFailedTimes(v int) *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.AddLoginFailedTimes(v)
	})
}

// UpdateLoginFailedTimes sets the "login_failed_times" field to the value that was provided on create.
func (u *AuthClientUpsertOne) UpdateLoginFailedTimes() *AuthClientUpsertOne {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateLoginFailedTimes()
	})
}

// Exec executes the query.
func (u *AuthClientUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthClientCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthClientUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AuthClientUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AuthClientUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AuthClientCreateBulk is the builder for creating many AuthClient entities in bulk.
type AuthClientCreateBulk struct {
	config
	err      error
	builders []*AuthClientCreate
	conflict []sql.ConflictOption
}

// Save creates the AuthClient entities in the database.
func (accb *AuthClientCreateBulk) Save(ctx context.Context) ([]*AuthClient, error) {
	if accb.err != nil {
		return nil, accb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*AuthClient, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthClientMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = accb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *AuthClientCreateBulk) SaveX(ctx context.Context) []*AuthClient {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *AuthClientCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *AuthClientCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AuthClient.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthClientUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (accb *AuthClientCreateBulk) OnConflict(opts ...sql.ConflictOption) *AuthClientUpsertBulk {
	accb.conflict = opts
	return &AuthClientUpsertBulk{
		create: accb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AuthClient.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (accb *AuthClientCreateBulk) OnConflictColumns(columns ...string) *AuthClientUpsertBulk {
	accb.conflict = append(accb.conflict, sql.ConflictColumns(columns...))
	return &AuthClientUpsertBulk{
		create: accb,
	}
}

// AuthClientUpsertBulk is the builder for "upsert"-ing
// a bulk of AuthClient nodes.
type AuthClientUpsertBulk struct {
	create *AuthClientCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AuthClient.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(authclient.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuthClientUpsertBulk) UpdateNewValues() *AuthClientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(authclient.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(authclient.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AuthClient.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AuthClientUpsertBulk) Ignore() *AuthClientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthClientUpsertBulk) DoNothing() *AuthClientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthClientCreateBulk.OnConflict
// documentation for more info.
func (u *AuthClientUpsertBulk) Update(set func(*AuthClientUpsert)) *AuthClientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthClientUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthClientUpsertBulk) SetUpdatedAt(v time.Time) *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthClientUpsertBulk) UpdateUpdatedAt() *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetClientType sets the "client_type" field.
func (u *AuthClientUpsertBulk) SetClientType(v int) *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetClientType(v)
	})
}

// AddClientType adds v to the "client_type" field.
func (u *AuthClientUpsertBulk) AddClientType(v int) *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.AddClientType(v)
	})
}

// UpdateClientType sets the "client_type" field to the value that was provided on create.
func (u *AuthClientUpsertBulk) UpdateClientType() *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateClientType()
	})
}

// SetMerchantID sets the "merchant_id" field.
func (u *AuthClientUpsertBulk) SetMerchantID(v int64) *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetMerchantID(v)
	})
}

// AddMerchantID adds v to the "merchant_id" field.
func (u *AuthClientUpsertBulk) AddMerchantID(v int64) *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.AddMerchantID(v)
	})
}

// UpdateMerchantID sets the "merchant_id" field to the value that was provided on create.
func (u *AuthClientUpsertBulk) UpdateMerchantID() *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateMerchantID()
	})
}

// SetSecret sets the "secret" field.
func (u *AuthClientUpsertBulk) SetSecret(v string) *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetSecret(v)
	})
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *AuthClientUpsertBulk) UpdateSecret() *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateSecret()
	})
}

// SetActive sets the "active" field.
func (u *AuthClientUpsertBulk) SetActive(v bool) *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *AuthClientUpsertBulk) UpdateActive() *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateActive()
	})
}

// SetTokenExpireSecs sets the "token_expire_secs" field.
func (u *AuthClientUpsertBulk) SetTokenExpireSecs(v int) *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetTokenExpireSecs(v)
	})
}

// AddTokenExpireSecs adds v to the "token_expire_secs" field.
func (u *AuthClientUpsertBulk) AddTokenExpireSecs(v int) *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.AddTokenExpireSecs(v)
	})
}

// UpdateTokenExpireSecs sets the "token_expire_secs" field to the value that was provided on create.
func (u *AuthClientUpsertBulk) UpdateTokenExpireSecs() *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateTokenExpireSecs()
	})
}

// SetLoginFailedTimes sets the "login_failed_times" field.
func (u *AuthClientUpsertBulk) SetLoginFailedTimes(v int) *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.SetLoginFailedTimes(v)
	})
}

// AddLoginFailedTimes adds v to the "login_failed_times" field.
func (u *AuthClientUpsertBulk) AddLoginFailedTimes(v int) *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.AddLoginFailedTimes(v)
	})
}

// UpdateLoginFailedTimes sets the "login_failed_times" field to the value that was provided on create.
func (u *AuthClientUpsertBulk) UpdateLoginFailedTimes() *AuthClientUpsertBulk {
	return u.Update(func(s *AuthClientUpsert) {
		s.UpdateLoginFailedTimes()
	})
}

// Exec executes the query.
func (u *AuthClientUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AuthClientCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthClientCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthClientUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
