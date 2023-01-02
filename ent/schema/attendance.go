package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Attendance holds the schema definition for the Attendance entity.
type Attendance struct {
	ent.Schema
}

// Fields of the Attendance.
func (Attendance) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			SchemaType(map[string]string{
				dialect.MySQL: "int(5)",
			}).Unique(),
		field.Int64("employee_id").
			SchemaType(map[string]string{
				dialect.MySQL: "int(5)",
			}),
		field.Time("clocked_in").
			SchemaType(map[string]string{
				dialect.MySQL: "timestamp",
			}),
		field.Time("clocked_out").
			SchemaType(map[string]string{
				dialect.MySQL: "timestamp",
			}),
	}
}

// Edges of the Attendance.
func (Attendance) Edges() []ent.Edge {
	return nil
}
