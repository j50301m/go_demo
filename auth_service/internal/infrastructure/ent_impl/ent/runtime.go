// Code generated by ent, DO NOT EDIT.

package ent

import (
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/authclient"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/loginrecord"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/role"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/schema"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authclientMixin := schema.AuthClient{}.Mixin()
	authclientMixinFields0 := authclientMixin[0].Fields()
	_ = authclientMixinFields0
	authclientFields := schema.AuthClient{}.Fields()
	_ = authclientFields
	// authclientDescCreatedAt is the schema descriptor for created_at field.
	authclientDescCreatedAt := authclientMixinFields0[0].Descriptor()
	// authclient.DefaultCreatedAt holds the default value on creation for the created_at field.
	authclient.DefaultCreatedAt = authclientDescCreatedAt.Default.(func() time.Time)
	// authclientDescUpdatedAt is the schema descriptor for updated_at field.
	authclientDescUpdatedAt := authclientMixinFields0[1].Descriptor()
	// authclient.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	authclient.DefaultUpdatedAt = authclientDescUpdatedAt.Default.(func() time.Time)
	// authclient.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	authclient.UpdateDefaultUpdatedAt = authclientDescUpdatedAt.UpdateDefault.(func() time.Time)
	loginrecordMixin := schema.LoginRecord{}.Mixin()
	loginrecordMixinFields0 := loginrecordMixin[0].Fields()
	_ = loginrecordMixinFields0
	loginrecordFields := schema.LoginRecord{}.Fields()
	_ = loginrecordFields
	// loginrecordDescCreatedAt is the schema descriptor for created_at field.
	loginrecordDescCreatedAt := loginrecordMixinFields0[0].Descriptor()
	// loginrecord.DefaultCreatedAt holds the default value on creation for the created_at field.
	loginrecord.DefaultCreatedAt = loginrecordDescCreatedAt.Default.(func() time.Time)
	// loginrecordDescUpdatedAt is the schema descriptor for updated_at field.
	loginrecordDescUpdatedAt := loginrecordMixinFields0[1].Descriptor()
	// loginrecord.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	loginrecord.DefaultUpdatedAt = loginrecordDescUpdatedAt.Default.(func() time.Time)
	// loginrecord.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	loginrecord.UpdateDefaultUpdatedAt = loginrecordDescUpdatedAt.UpdateDefault.(func() time.Time)
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleMixinFields0[0].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleMixinFields0[1].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roleDescIsSystem is the schema descriptor for is_system field.
	roleDescIsSystem := roleFields[3].Descriptor()
	// role.DefaultIsSystem holds the default value on creation for the is_system field.
	role.DefaultIsSystem = roleDescIsSystem.Default.(bool)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
