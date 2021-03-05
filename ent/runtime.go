// Code generated by entc, DO NOT EDIT.

package ent

import (
	"abc/ent/schema"
	"abc/ent/test"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	testMixin := schema.Test{}.Mixin()
	testMixinFields0 := testMixin[0].Fields()
	_ = testMixinFields0
	testFields := schema.Test{}.Fields()
	_ = testFields
	// testDescCreatedAt is the schema descriptor for created_at field.
	testDescCreatedAt := testMixinFields0[1].Descriptor()
	// test.DefaultCreatedAt holds the default value on creation for the created_at field.
	test.DefaultCreatedAt = testDescCreatedAt.Default.(func() time.Time)
	// testDescUpdatedAt is the schema descriptor for updated_at field.
	testDescUpdatedAt := testMixinFields0[2].Descriptor()
	// test.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	test.DefaultUpdatedAt = testDescUpdatedAt.Default.(func() time.Time)
	// test.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	test.UpdateDefaultUpdatedAt = testDescUpdatedAt.UpdateDefault.(func() time.Time)
	// testDescID is the schema descriptor for id field.
	testDescID := testMixinFields0[0].Descriptor()
	// test.IDValidator is a validator for the "id" field. It is called by the builders before save.
	test.IDValidator = testDescID.Validators[0].(func(uint64) error)
}