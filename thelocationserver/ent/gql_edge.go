// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (c *Car) Carsowned(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedCarsowned(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.CarsownedOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryCarsowned().All(ctx)
	}
	return result, err
}

func (u *User) Cars(ctx context.Context) (result []*Car, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedCars(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.CarsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryCars().All(ctx)
	}
	return result, err
}
