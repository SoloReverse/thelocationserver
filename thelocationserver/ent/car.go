// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"thelocationserver/ent/car"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Car is the model entity for the Car schema.
type Car struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// IconURL holds the value of the "iconURL" field.
	IconURL string `json:"iconURL,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CarQuery when eager-loading is set.
	Edges        CarEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CarEdges holds the relations/edges for other nodes in the graph.
type CarEdges struct {
	// Carsowned holds the value of the carsowned edge.
	Carsowned []*User `json:"carsowned,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedCarsowned map[string][]*User
}

// CarsownedOrErr returns the Carsowned value or an error if the edge
// was not loaded in eager-loading.
func (e CarEdges) CarsownedOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Carsowned, nil
	}
	return nil, &NotLoadedError{edge: "carsowned"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Car) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case car.FieldModel, car.FieldCountry, car.FieldIconURL:
			values[i] = new(sql.NullString)
		case car.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Car fields.
func (c *Car) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case car.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case car.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				c.Model = value.String
			}
		case car.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				c.Country = value.String
			}
		case car.FieldIconURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field iconURL", values[i])
			} else if value.Valid {
				c.IconURL = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Car.
// This includes values selected through modifiers, order, etc.
func (c *Car) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryCarsowned queries the "carsowned" edge of the Car entity.
func (c *Car) QueryCarsowned() *UserQuery {
	return NewCarClient(c.config).QueryCarsowned(c)
}

// Update returns a builder for updating this Car.
// Note that you need to call Car.Unwrap() before calling this method if this Car
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Car) Update() *CarUpdateOne {
	return NewCarClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Car entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Car) Unwrap() *Car {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Car is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Car) String() string {
	var builder strings.Builder
	builder.WriteString("Car(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("model=")
	builder.WriteString(c.Model)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(c.Country)
	builder.WriteString(", ")
	builder.WriteString("iconURL=")
	builder.WriteString(c.IconURL)
	builder.WriteByte(')')
	return builder.String()
}

// NamedCarsowned returns the Carsowned named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Car) NamedCarsowned(name string) ([]*User, error) {
	if c.Edges.namedCarsowned == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCarsowned[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Car) appendNamedCarsowned(name string, edges ...*User) {
	if c.Edges.namedCarsowned == nil {
		c.Edges.namedCarsowned = make(map[string][]*User)
	}
	if len(edges) == 0 {
		c.Edges.namedCarsowned[name] = []*User{}
	} else {
		c.Edges.namedCarsowned[name] = append(c.Edges.namedCarsowned[name], edges...)
	}
}

// Cars is a parsable slice of Car.
type Cars []*Car
