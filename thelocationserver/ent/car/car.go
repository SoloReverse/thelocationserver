// Code generated by ent, DO NOT EDIT.

package car

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the car type in the database.
	Label = "car"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldIconURL holds the string denoting the iconurl field in the database.
	FieldIconURL = "icon_url"
	// EdgeCarsowned holds the string denoting the carsowned edge name in mutations.
	EdgeCarsowned = "carsowned"
	// Table holds the table name of the car in the database.
	Table = "cars"
	// CarsownedTable is the table that holds the carsowned relation/edge. The primary key declared below.
	CarsownedTable = "car_carsowned"
	// CarsownedInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CarsownedInverseTable = "users"
)

// Columns holds all SQL columns for car fields.
var Columns = []string{
	FieldID,
	FieldModel,
	FieldCountry,
	FieldIconURL,
}

var (
	// CarsownedPrimaryKey and CarsownedColumn2 are the table columns denoting the
	// primary key for the carsowned relation (M2M).
	CarsownedPrimaryKey = []string{"car_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Car queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByModel orders the results by the model field.
func ByModel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModel, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// ByIconURL orders the results by the iconURL field.
func ByIconURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIconURL, opts...).ToFunc()
}

// ByCarsownedCount orders the results by carsowned count.
func ByCarsownedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCarsownedStep(), opts...)
	}
}

// ByCarsowned orders the results by carsowned terms.
func ByCarsowned(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCarsownedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCarsownedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CarsownedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, CarsownedTable, CarsownedPrimaryKey...),
	)
}
