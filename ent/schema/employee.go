package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			SchemaType(map[string]string{
				dialect.MySQL: "int(5)",
			}).Unique(),
		field.String("name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}),
		field.String("email").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return nil

}
