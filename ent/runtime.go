// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/riadafridishibly/vault-app/ent/file"
	"github.com/riadafridishibly/vault-app/ent/key"
	"github.com/riadafridishibly/vault-app/ent/masterpassword"
	"github.com/riadafridishibly/vault-app/ent/passworditem"
	"github.com/riadafridishibly/vault-app/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	fileMixin := schema.File{}.Mixin()
	fileMixinFields0 := fileMixin[0].Fields()
	_ = fileMixinFields0
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescCreateTime is the schema descriptor for create_time field.
	fileDescCreateTime := fileMixinFields0[0].Descriptor()
	// file.DefaultCreateTime holds the default value on creation for the create_time field.
	file.DefaultCreateTime = fileDescCreateTime.Default.(func() time.Time)
	// fileDescUpdateTime is the schema descriptor for update_time field.
	fileDescUpdateTime := fileMixinFields0[1].Descriptor()
	// file.DefaultUpdateTime holds the default value on creation for the update_time field.
	file.DefaultUpdateTime = fileDescUpdateTime.Default.(func() time.Time)
	// file.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	file.UpdateDefaultUpdateTime = fileDescUpdateTime.UpdateDefault.(func() time.Time)
	keyMixin := schema.Key{}.Mixin()
	keyMixinFields0 := keyMixin[0].Fields()
	_ = keyMixinFields0
	keyFields := schema.Key{}.Fields()
	_ = keyFields
	// keyDescCreateTime is the schema descriptor for create_time field.
	keyDescCreateTime := keyMixinFields0[0].Descriptor()
	// key.DefaultCreateTime holds the default value on creation for the create_time field.
	key.DefaultCreateTime = keyDescCreateTime.Default.(func() time.Time)
	// keyDescUpdateTime is the schema descriptor for update_time field.
	keyDescUpdateTime := keyMixinFields0[1].Descriptor()
	// key.DefaultUpdateTime holds the default value on creation for the update_time field.
	key.DefaultUpdateTime = keyDescUpdateTime.Default.(func() time.Time)
	// key.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	key.UpdateDefaultUpdateTime = keyDescUpdateTime.UpdateDefault.(func() time.Time)
	masterpasswordFields := schema.MasterPassword{}.Fields()
	_ = masterpasswordFields
	// masterpasswordDescPasswordHash is the schema descriptor for password_hash field.
	masterpasswordDescPasswordHash := masterpasswordFields[0].Descriptor()
	// masterpassword.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	masterpassword.PasswordHashValidator = masterpasswordDescPasswordHash.Validators[0].(func(string) error)
	passworditemMixin := schema.PasswordItem{}.Mixin()
	passworditemMixinFields0 := passworditemMixin[0].Fields()
	_ = passworditemMixinFields0
	passworditemFields := schema.PasswordItem{}.Fields()
	_ = passworditemFields
	// passworditemDescCreateTime is the schema descriptor for create_time field.
	passworditemDescCreateTime := passworditemMixinFields0[0].Descriptor()
	// passworditem.DefaultCreateTime holds the default value on creation for the create_time field.
	passworditem.DefaultCreateTime = passworditemDescCreateTime.Default.(func() time.Time)
	// passworditemDescUpdateTime is the schema descriptor for update_time field.
	passworditemDescUpdateTime := passworditemMixinFields0[1].Descriptor()
	// passworditem.DefaultUpdateTime holds the default value on creation for the update_time field.
	passworditem.DefaultUpdateTime = passworditemDescUpdateTime.Default.(func() time.Time)
	// passworditem.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	passworditem.UpdateDefaultUpdateTime = passworditemDescUpdateTime.UpdateDefault.(func() time.Time)
}
