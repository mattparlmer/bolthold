/*
Package bolthold is an indexing and querying layer on top of a Bolt DB.  The goal is to allow easy, persistant storage
and retrieval of Go types.  BoltDB is an embedded key-value store, and bolthold servers a similar use case however with
a higher level interface for common uses of BoltDB.

Go Types

BoltHold deals directly with Go Types.  When inserting data, you pass in your stucture directly.  When querying data you
pass in a pointer to a slice of the type you want to return.  By default Gob encoding is used.

err := store.Insert(1234, ItemTest{
	Name:    "Test Name",
	Created: time.Now(),
})

var result []ItemTest

err := store.Find(&result, query)


Indexes

BoltHold will automatically create an index for any struct fields tags with "boldholdIndex"

type ItemTest struct {
	ID       int
	Name     string
	Category string `boltholdIndex:"Category"`
	Created  time.Time
}

The first field specified in query will be used as the index (if one exists).

Queries are chained together criteria that applies to a set of fields:
	bolthold.Where("Name").Eq("John Doe").And("DOB").Lt(time.Now())


*/
package bolthold