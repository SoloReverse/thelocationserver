package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

type Car struct {
	ent.Schema
}

func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("model"),
		field.String("country"),
		field.String("iconURL"),
	}
}

func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("carsowned", User.Type),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cars", Car.Type).Ref("carsowned"),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("username").
			Unique(),
		field.String("phonenumber"),
		field.String("photoURL"),
		field.String("car"),
	}

}
