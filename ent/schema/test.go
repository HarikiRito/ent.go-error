package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Test holds the schema definition for the Test entity.
type Test struct {
	ent.Schema
}

func (Test) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Test.
func (Test) Fields() []ent.Field {
	return []ent.Field{
		field.String("note"),
	}
}

// Edges of the Test.
func (Test) Edges() []ent.Edge {
	return nil
}

func (Test) Hooks() []ent.Hook {
	return []ent.Hook{
		SampleHook(),
	}
}
