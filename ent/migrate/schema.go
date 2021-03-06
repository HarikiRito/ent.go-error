// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TestsColumns holds the columns for the "tests" table.
	TestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "note", Type: field.TypeString},
	}
	// TestsTable holds the schema information for the "tests" table.
	TestsTable = &schema.Table{
		Name:        "tests",
		Columns:     TestsColumns,
		PrimaryKey:  []*schema.Column{TestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TestsTable,
	}
)

func init() {
}
