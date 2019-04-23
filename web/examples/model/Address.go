package model

// This is the implementation file for the Address ORM object.
// This is where you build the api to your data model for you web application and potentially mobile apps.
// Your edits to this file will be preserved.

import (
	"context"
	"fmt"

	"github.com/goradd/goradd/pkg/orm/query"
)

type Address struct {
	addressBase
}

// Create a new Address object and initialize to default values.
func NewAddress(ctx context.Context) *Address {
	o := new(Address)
	o.Initialize(ctx)
	return o
}

// Initialize or re-initialize a Address database object to default values.
func (o *Address) Initialize(ctx context.Context) {
	o.addressBase.Initialize()
	// Add your own initializations here
}

// String implements the Stringer interface and returns the default label for the object as it appears in html lists.
// Typically you would change this to whatever was pertinent to your application.
func (o *Address) String() string {
	if o == nil {
		return "" // Possibly - Select One -?
	}
	return fmt.Sprintf("Object id %v", o.PrimaryKey())
}

// QueryAddresses returns a new builder that gives you general purpose access to the Address records
// in the database. This is useful for quick queries of the database during development, but eventually you
// should remove this function and move those queries to more specific calls in this file.
func QueryAddresses() *AddressesBuilder {
	return queryAddresses()
}

// LoadAddress queries for a single Address object by primary key.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
// If you need a more elaborate query, use QueryAddresses() to start a query builder.
func LoadAddress(ctx context.Context, pk string, joinOrSelectNodes ...query.NodeI) *Address {
	return loadAddress(ctx, pk, joinOrSelectNodes...)
}

// DeleteAddress deletes the give record from the database. Note that you can also delete
// loaded Address objects by calling Delete on them.
func DeleteAddress(ctx context.Context, pk string) {
	deleteAddress(ctx, pk)
}
