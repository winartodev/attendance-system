package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			SchemaType(map[string]string{
				dialect.MySQL: "int(5)",
			}),
		field.String("username").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}),
		field.String("password").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),
		field.Int64("role").
			SchemaType(map[string]string{
				dialect.MySQL: "int(2)",
			}).Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
