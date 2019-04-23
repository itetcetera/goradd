package model

// Code generated by goradd. DO NOT EDIT.

import (
	"context"
	"fmt"
	"github.com/goradd/goradd/web/examples/model/node"

	"github.com/goradd/goradd/pkg/orm/db"
	. "github.com/goradd/goradd/pkg/orm/op"
	"github.com/goradd/goradd/pkg/orm/query"

	//"./node"
	"bytes"
	"encoding/gob"

	"github.com/goradd/goradd/pkg/datetime"
)

// typeTestBase is a base structure to be embedded in a "subclass" and provides the ORM access to the database.
// Do not directly access the internal variables, but rather use the accessor functions, since this class maintains internal state
// related to the variables.

type typeTestBase struct {
	id        string
	idIsValid bool
	idIsDirty bool

	date        datetime.DateTime
	dateIsNull  bool
	dateIsValid bool
	dateIsDirty bool

	time        datetime.DateTime
	timeIsNull  bool
	timeIsValid bool
	timeIsDirty bool

	dateTime        datetime.DateTime
	dateTimeIsNull  bool
	dateTimeIsValid bool
	dateTimeIsDirty bool

	ts        datetime.DateTime
	tsIsNull  bool
	tsIsValid bool
	tsIsDirty bool

	testInt        int
	testIntIsNull  bool
	testIntIsValid bool
	testIntIsDirty bool

	testFloat        float32
	testFloatIsNull  bool
	testFloatIsValid bool
	testFloatIsDirty bool

	testText        string
	testTextIsNull  bool
	testTextIsValid bool
	testTextIsDirty bool

	testBit        bool
	testBitIsNull  bool
	testBitIsValid bool
	testBitIsDirty bool

	testVarchar        string
	testVarcharIsNull  bool
	testVarcharIsValid bool
	testVarcharIsDirty bool

	// Custom aliases, if specified
	_aliases map[string]interface{}

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update
	_restored bool
}

const (
	TypeTestIDDefault          = ""
	TypeTestDateDefault        = datetime.Zero
	TypeTestTimeDefault        = datetime.Zero
	TypeTestDateTimeDefault    = datetime.Zero
	TypeTestTsDefault          = datetime.Zero
	TypeTestTestIntDefault     = 5
	TypeTestTestFloatDefault   = 0.0
	TypeTestTestTextDefault    = ""
	TypeTestTestBitDefault     = false
	TypeTestTestVarcharDefault = ""
)

const (
	TypeTestID          = `ID`
	TypeTestDate        = `Date`
	TypeTestTime        = `Time`
	TypeTestDateTime    = `DateTime`
	TypeTestTs          = `Ts`
	TypeTestTestInt     = `TestInt`
	TypeTestTestFloat   = `TestFloat`
	TypeTestTestText    = `TestText`
	TypeTestTestBit     = `TestBit`
	TypeTestTestVarchar = `TestVarchar`
)

// Initialize or re-initialize a TypeTest database object to default values.
func (o *typeTestBase) Initialize() {

	o.id = ""
	o.idIsValid = false
	o.idIsDirty = false

	o.date = datetime.DateTime{}
	o.dateIsNull = true
	o.dateIsValid = true
	o.dateIsDirty = true

	o.time = datetime.DateTime{}
	o.timeIsNull = true
	o.timeIsValid = true
	o.timeIsDirty = true

	o.dateTime = datetime.DateTime{}
	o.dateTimeIsNull = true
	o.dateTimeIsValid = true
	o.dateTimeIsDirty = true

	o.ts = datetime.DateTime{}
	o.tsIsNull = true
	o.tsIsValid = true
	o.tsIsDirty = true

	o.testInt = 5
	o.testIntIsNull = false
	o.testIntIsValid = true
	o.testIntIsDirty = true

	o.testFloat = 0.0
	o.testFloatIsNull = true
	o.testFloatIsValid = true
	o.testFloatIsDirty = true

	o.testText = ""
	o.testTextIsNull = true
	o.testTextIsValid = true
	o.testTextIsDirty = true

	o.testBit = false
	o.testBitIsNull = true
	o.testBitIsValid = true
	o.testBitIsDirty = true

	o.testVarchar = ""
	o.testVarcharIsNull = true
	o.testVarcharIsValid = true
	o.testVarcharIsDirty = true

	o._restored = false
}

func (o *typeTestBase) PrimaryKey() string {
	return o.id
}

// ID returns the loaded value of ID.
func (o *typeTestBase) ID() string {
	return fmt.Sprint(o.id)
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *typeTestBase) IDIsValid() bool {
	return o._restored && o.idIsValid
}

func (o *typeTestBase) Date() datetime.DateTime {
	if o._restored && !o.dateIsValid {
		panic("date was not selected in the last query and so is not valid")
	}
	return o.date
}

// DateIsValid returns true if the value was loaded from the database or has been set.
func (o *typeTestBase) DateIsValid() bool {
	return o.dateIsValid
}

// DateIsNull returns true if the related database value is null.
func (o *typeTestBase) DateIsNull() bool {
	return o.dateIsNull
}

func (o *typeTestBase) SetDate(i interface{}) {
	if i == nil {
		if !o.dateIsNull {
			o.dateIsNull = true
			o.dateIsDirty = true
			o.date = datetime.DateTime{}
		}
	} else {
		v := i.(datetime.DateTime)
		if o.dateIsNull ||
			!o._restored ||
			o.date != v {

			o.dateIsNull = false
			o.date = v
			o.dateIsDirty = true
		}
	}
}

func (o *typeTestBase) Time() datetime.DateTime {
	if o._restored && !o.timeIsValid {
		panic("time was not selected in the last query and so is not valid")
	}
	return o.time
}

// TimeIsValid returns true if the value was loaded from the database or has been set.
func (o *typeTestBase) TimeIsValid() bool {
	return o.timeIsValid
}

// TimeIsNull returns true if the related database value is null.
func (o *typeTestBase) TimeIsNull() bool {
	return o.timeIsNull
}

func (o *typeTestBase) SetTime(i interface{}) {
	if i == nil {
		if !o.timeIsNull {
			o.timeIsNull = true
			o.timeIsDirty = true
			o.time = datetime.DateTime{}
		}
	} else {
		v := i.(datetime.DateTime)
		if o.timeIsNull ||
			!o._restored ||
			o.time != v {

			o.timeIsNull = false
			o.time = v
			o.timeIsDirty = true
		}
	}
}

func (o *typeTestBase) DateTime() datetime.DateTime {
	if o._restored && !o.dateTimeIsValid {
		panic("dateTime was not selected in the last query and so is not valid")
	}
	return o.dateTime
}

// DateTimeIsValid returns true if the value was loaded from the database or has been set.
func (o *typeTestBase) DateTimeIsValid() bool {
	return o.dateTimeIsValid
}

// DateTimeIsNull returns true if the related database value is null.
func (o *typeTestBase) DateTimeIsNull() bool {
	return o.dateTimeIsNull
}

func (o *typeTestBase) SetDateTime(i interface{}) {
	if i == nil {
		if !o.dateTimeIsNull {
			o.dateTimeIsNull = true
			o.dateTimeIsDirty = true
			o.dateTime = datetime.DateTime{}
		}
	} else {
		v := i.(datetime.DateTime)
		if o.dateTimeIsNull ||
			!o._restored ||
			o.dateTime != v {

			o.dateTimeIsNull = false
			o.dateTime = v
			o.dateTimeIsDirty = true
		}
	}
}

func (o *typeTestBase) Ts() datetime.DateTime {
	if o._restored && !o.tsIsValid {
		panic("ts was not selected in the last query and so is not valid")
	}
	return o.ts
}

// TsIsValid returns true if the value was loaded from the database or has been set.
func (o *typeTestBase) TsIsValid() bool {
	return o.tsIsValid
}

// TsIsNull returns true if the related database value is null.
func (o *typeTestBase) TsIsNull() bool {
	return o.tsIsNull
}

func (o *typeTestBase) SetTs(i interface{}) {
	if i == nil {
		if !o.tsIsNull {
			o.tsIsNull = true
			o.tsIsDirty = true
			o.ts = datetime.DateTime{}
		}
	} else {
		v := i.(datetime.DateTime)
		if o.tsIsNull ||
			!o._restored ||
			o.ts != v {

			o.tsIsNull = false
			o.ts = v
			o.tsIsDirty = true
		}
	}
}

func (o *typeTestBase) TestInt() int {
	if o._restored && !o.testIntIsValid {
		panic("testInt was not selected in the last query and so is not valid")
	}
	return o.testInt
}

// TestIntIsValid returns true if the value was loaded from the database or has been set.
func (o *typeTestBase) TestIntIsValid() bool {
	return o.testIntIsValid
}

// TestIntIsNull returns true if the related database value is null.
func (o *typeTestBase) TestIntIsNull() bool {
	return o.testIntIsNull
}

func (o *typeTestBase) SetTestInt(i interface{}) {
	if i == nil {
		if !o.testIntIsNull {
			o.testIntIsNull = true
			o.testIntIsDirty = true
			o.testInt = 5
		}
	} else {
		v := i.(int)
		if o.testIntIsNull ||
			!o._restored ||
			o.testInt != v {

			o.testIntIsNull = false
			o.testInt = v
			o.testIntIsDirty = true
		}
	}
}

func (o *typeTestBase) TestFloat() float32 {
	if o._restored && !o.testFloatIsValid {
		panic("testFloat was not selected in the last query and so is not valid")
	}
	return o.testFloat
}

// TestFloatIsValid returns true if the value was loaded from the database or has been set.
func (o *typeTestBase) TestFloatIsValid() bool {
	return o.testFloatIsValid
}

// TestFloatIsNull returns true if the related database value is null.
func (o *typeTestBase) TestFloatIsNull() bool {
	return o.testFloatIsNull
}

func (o *typeTestBase) SetTestFloat(i interface{}) {
	if i == nil {
		if !o.testFloatIsNull {
			o.testFloatIsNull = true
			o.testFloatIsDirty = true
			o.testFloat = 0.0
		}
	} else {
		v := i.(float32)
		if o.testFloatIsNull ||
			!o._restored ||
			o.testFloat != v {

			o.testFloatIsNull = false
			o.testFloat = v
			o.testFloatIsDirty = true
		}
	}
}

func (o *typeTestBase) TestText() string {
	if o._restored && !o.testTextIsValid {
		panic("testText was not selected in the last query and so is not valid")
	}
	return o.testText
}

// TestTextIsValid returns true if the value was loaded from the database or has been set.
func (o *typeTestBase) TestTextIsValid() bool {
	return o.testTextIsValid
}

// TestTextIsNull returns true if the related database value is null.
func (o *typeTestBase) TestTextIsNull() bool {
	return o.testTextIsNull
}

func (o *typeTestBase) SetTestText(i interface{}) {
	if i == nil {
		if !o.testTextIsNull {
			o.testTextIsNull = true
			o.testTextIsDirty = true
			o.testText = ""
		}
	} else {
		v := i.(string)
		if o.testTextIsNull ||
			!o._restored ||
			o.testText != v {

			o.testTextIsNull = false
			o.testText = v
			o.testTextIsDirty = true
		}
	}
}

func (o *typeTestBase) TestBit() bool {
	if o._restored && !o.testBitIsValid {
		panic("testBit was not selected in the last query and so is not valid")
	}
	return o.testBit
}

// TestBitIsValid returns true if the value was loaded from the database or has been set.
func (o *typeTestBase) TestBitIsValid() bool {
	return o.testBitIsValid
}

// TestBitIsNull returns true if the related database value is null.
func (o *typeTestBase) TestBitIsNull() bool {
	return o.testBitIsNull
}

func (o *typeTestBase) SetTestBit(i interface{}) {
	if i == nil {
		if !o.testBitIsNull {
			o.testBitIsNull = true
			o.testBitIsDirty = true
			o.testBit = false
		}
	} else {
		v := i.(bool)
		if o.testBitIsNull ||
			!o._restored ||
			o.testBit != v {

			o.testBitIsNull = false
			o.testBit = v
			o.testBitIsDirty = true
		}
	}
}

func (o *typeTestBase) TestVarchar() string {
	if o._restored && !o.testVarcharIsValid {
		panic("testVarchar was not selected in the last query and so is not valid")
	}
	return o.testVarchar
}

// TestVarcharIsValid returns true if the value was loaded from the database or has been set.
func (o *typeTestBase) TestVarcharIsValid() bool {
	return o.testVarcharIsValid
}

// TestVarcharIsNull returns true if the related database value is null.
func (o *typeTestBase) TestVarcharIsNull() bool {
	return o.testVarcharIsNull
}

func (o *typeTestBase) SetTestVarchar(i interface{}) {
	if i == nil {
		if !o.testVarcharIsNull {
			o.testVarcharIsNull = true
			o.testVarcharIsDirty = true
			o.testVarchar = ""
		}
	} else {
		v := i.(string)
		if o.testVarcharIsNull ||
			!o._restored ||
			o.testVarchar != v {

			o.testVarcharIsNull = false
			o.testVarchar = v
			o.testVarcharIsDirty = true
		}
	}
}

// GetAlias returns the alias for the given key.
func (o *typeTestBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
		return query.NewAliasValue([]byte{})
	}
}

// loadTypeTest queries for a single TypeTest object by primary key.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
func loadTypeTest(ctx context.Context, pk string, joinOrSelectNodes ...query.NodeI) *TypeTest {
	return queryTypeTests().Where(Equal(node.TypeTest().ID(), pk)).joinOrSelect(joinOrSelectNodes...).Get(ctx)
}

func queryTypeTests() *TypeTestsBuilder {
	return newTypeTestBuilder()
}

// The TypeTestsBuilder uses the QueryBuilderI interface from the database to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, Count, or Delete
type TypeTestsBuilder struct {
	base                query.QueryBuilderI
	hasConditionalJoins bool
}

func newTypeTestBuilder() *TypeTestsBuilder {
	b := &TypeTestsBuilder{
		base: db.GetDatabase("goradd").
			NewBuilder(),
	}
	return b.Join(node.TypeTest())
}

// Load terminates the query builder, performs the query, and returns a slice of TypeTest objects. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice
func (b *TypeTestsBuilder) Load(ctx context.Context) (typeTestSlice []*TypeTest) {
	results := b.base.Load(ctx)
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(TypeTest)
		o.load(item, !b.hasConditionalJoins, o, nil, "")
		typeTestSlice = append(typeTestSlice, o)
	}
	return typeTestSlice
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice.
func (b *TypeTestsBuilder) LoadI(ctx context.Context) (typeTestSlice []interface{}) {
	results := b.base.Load(ctx)
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(TypeTest)
		o.load(item, !b.hasConditionalJoins, o, nil, "")
		typeTestSlice = append(typeTestSlice, o)
	}
	return typeTestSlice
}

// Get is a convenience method to return only the first item found in a query. It is equivalent to adding
// Limit(1,0) to the query, and then getting the first item from the returned slice.
// Limits with joins do not currently work, so don't try it if you have a join
// TODO: Change this to Load1 to be more descriptive and avoid confusion with other Getters
func (b *TypeTestsBuilder) Get(ctx context.Context) *TypeTest {
	results := b.Limit(1, 0).Load(ctx)
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Expand expands an array type node so that it will produce individual rows instead of an array of items
func (b *TypeTestsBuilder) Expand(n query.NodeI) *TypeTestsBuilder {
	b.base.Expand(n)
	return b
}

// Join adds a node to the node tree so that its fields will appear in the query. Optionally add conditions to filter
// what gets included. The conditions will be AND'd with the basic condition matching the primary keys of the join.
func (b *TypeTestsBuilder) Join(n query.NodeI, conditions ...query.NodeI) *TypeTestsBuilder {
	var condition query.NodeI
	if len(conditions) > 1 {
		condition = And(conditions)
	} else if len(conditions) == 1 {
		condition = conditions[0]
	}
	b.base.Join(n, condition)
	if condition != nil {
		b.hasConditionalJoins = true
	}
	return b
}

// JoinOn adds a node to the node tree so that its fields will appear in the query. Optionally add conditions to filter
// what gets included. The conditions will be AND'd with the basic condition matching the primary keys of the join.
func (b *TypeTestsBuilder) JoinOn(n query.NodeI, conditions ...query.NodeI) *TypeTestsBuilder {
	var condition query.NodeI
	if len(conditions) > 1 {
		condition = And(conditions)
	} else if len(conditions) == 1 {
		condition = conditions[0]
	}
	b.base.Join(n, condition)
	if condition != nil {
		b.hasConditionalJoins = true
	}
	return b
}

// Where adds a condition to filter what gets selected.
func (b *TypeTestsBuilder) Where(c query.NodeI) *TypeTestsBuilder {
	b.base.Condition(c)
	return b
}

// OrderBy  spedifies how the resulting data should be sorted.
func (b *TypeTestsBuilder) OrderBy(nodes ...query.NodeI) *TypeTestsBuilder {
	b.base.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified
func (b *TypeTestsBuilder) Limit(maxRowCount int, offset int) *TypeTestsBuilder {
	b.base.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields. Once you put a Select in your query, you must
// specify all the fields that you will eventually read out. Be careful when selecting fields in joined tables, as joined
// tables will also contain pointers back to the parent table, and so the parent node should have the same field selected
// as the child node if you are querying those fields.
func (b *TypeTestsBuilder) Select(nodes ...query.NodeI) *TypeTestsBuilder {
	b.base.Select(nodes...)
	return b
}

// Alias lets you add a node with a custom name. After the query, you can read out the data using GetAlias() on a
// returned object. Alias is useful for adding calculations or subqueries to the query.
func (b *TypeTestsBuilder) Alias(name string, n query.NodeI) *TypeTestsBuilder {
	b.base.Alias(name, n)
	return b
}

// Distinct removes duplicates from the results of the query. Adding a Select() may help you get to the data you want, although
// using Distinct with joined tables is often not effective, since we force joined tables to include primary keys in the query, and this
// often ruins the effect of Distinct.
func (b *TypeTestsBuilder) Distinct() *TypeTestsBuilder {
	b.base.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions in an Alias() call.
func (b *TypeTestsBuilder) GroupBy(nodes ...query.NodeI) *TypeTestsBuilder {
	b.base.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query.
func (b *TypeTestsBuilder) Having(node query.NodeI) *TypeTestsBuilder {
	b.base.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
func (b *TypeTestsBuilder) Count(ctx context.Context, distinct bool, nodes ...query.NodeI) uint {
	return b.base.Count(ctx, distinct, nodes...)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *TypeTestsBuilder) Delete(ctx context.Context) {
	b.base.Delete(ctx)
}

// Subquery uses the query builder to define a subquery within a larger query. You MUST include what
// you are selecting by adding Alias or Select functions on the subquery builder. Generally you would use
// this as a node to an Alias function on the surrounding query builder.
func (b *TypeTestsBuilder) Subquery() *query.SubqueryNode {
	return b.base.Subquery()
}

// joinOrSelect us a private helper function for the Load* functions
func (b *TypeTestsBuilder) joinOrSelect(nodes ...query.NodeI) *TypeTestsBuilder {
	for _, n := range nodes {
		switch n.(type) {
		case query.TableNodeI:
			b.base.Join(n, nil)
		case *query.ColumnNode:
			b.Select(n)
		}
	}
	return b
}

func CountTypeTestByID(ctx context.Context, id string) uint {
	return queryTypeTests().Where(Equal(node.TypeTest().ID(), id)).Count(ctx, false)
}

func CountTypeTestByDate(ctx context.Context, date datetime.DateTime) uint {
	return queryTypeTests().Where(Equal(node.TypeTest().Date(), date)).Count(ctx, false)
}

func CountTypeTestByTime(ctx context.Context, time datetime.DateTime) uint {
	return queryTypeTests().Where(Equal(node.TypeTest().Time(), time)).Count(ctx, false)
}

func CountTypeTestByDateTime(ctx context.Context, dateTime datetime.DateTime) uint {
	return queryTypeTests().Where(Equal(node.TypeTest().DateTime(), dateTime)).Count(ctx, false)
}

func CountTypeTestByTs(ctx context.Context, ts datetime.DateTime) uint {
	return queryTypeTests().Where(Equal(node.TypeTest().Ts(), ts)).Count(ctx, false)
}

func CountTypeTestByTestInt(ctx context.Context, testInt int) uint {
	return queryTypeTests().Where(Equal(node.TypeTest().TestInt(), testInt)).Count(ctx, false)
}

func CountTypeTestByTestFloat(ctx context.Context, testFloat float32) uint {
	return queryTypeTests().Where(Equal(node.TypeTest().TestFloat(), testFloat)).Count(ctx, false)
}

func CountTypeTestByTestText(ctx context.Context, testText string) uint {
	return queryTypeTests().Where(Equal(node.TypeTest().TestText(), testText)).Count(ctx, false)
}

func CountTypeTestByTestBit(ctx context.Context, testBit bool) uint {
	return queryTypeTests().Where(Equal(node.TypeTest().TestBit(), testBit)).Count(ctx, false)
}

func CountTypeTestByTestVarchar(ctx context.Context, testVarchar string) uint {
	return queryTypeTests().Where(Equal(node.TypeTest().TestVarchar(), testVarchar)).Count(ctx, false)
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// If linkParent is true we will have child relationships use a pointer back to the parent object. If false, it will create a separate object.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
// Also, if any joins are conditional, that might affect which child objects are included, so in this situation, linkParent should be false
func (o *typeTestBase) load(m map[string]interface{}, linkParent bool, objThis *TypeTest, objParent interface{}, parentKey string) {
	if v, ok := m["id"]; ok && v != nil {
		if o.id, ok = v.(string); ok {
			o.idIsValid = true
			o.idIsDirty = false
		} else {
			panic("Wrong type found for id.")
		}
	} else {
		o.idIsValid = false
		o.id = ""
	}

	if v, ok := m["date"]; ok {
		if v == nil {
			o.date = datetime.DateTime{}
			o.dateIsNull = true
			o.dateIsValid = true
			o.dateIsDirty = false
		} else if o.date, ok = v.(datetime.DateTime); ok {
			o.dateIsNull = false
			o.dateIsValid = true
			o.dateIsDirty = false
		} else {
			panic("Wrong type found for date.")
		}
	} else {
		o.dateIsValid = false
		o.dateIsNull = true
		o.date = datetime.DateTime{}
	}
	if v, ok := m["time"]; ok {
		if v == nil {
			o.time = datetime.DateTime{}
			o.timeIsNull = true
			o.timeIsValid = true
			o.timeIsDirty = false
		} else if o.time, ok = v.(datetime.DateTime); ok {
			o.timeIsNull = false
			o.timeIsValid = true
			o.timeIsDirty = false
		} else {
			panic("Wrong type found for time.")
		}
	} else {
		o.timeIsValid = false
		o.timeIsNull = true
		o.time = datetime.DateTime{}
	}
	if v, ok := m["date_time"]; ok {
		if v == nil {
			o.dateTime = datetime.DateTime{}
			o.dateTimeIsNull = true
			o.dateTimeIsValid = true
			o.dateTimeIsDirty = false
		} else if o.dateTime, ok = v.(datetime.DateTime); ok {
			o.dateTimeIsNull = false
			o.dateTimeIsValid = true
			o.dateTimeIsDirty = false
		} else {
			panic("Wrong type found for date_time.")
		}
	} else {
		o.dateTimeIsValid = false
		o.dateTimeIsNull = true
		o.dateTime = datetime.DateTime{}
	}
	if v, ok := m["ts"]; ok {
		if v == nil {
			o.ts = datetime.DateTime{}
			o.tsIsNull = true
			o.tsIsValid = true
			o.tsIsDirty = false
		} else if o.ts, ok = v.(datetime.DateTime); ok {
			o.tsIsNull = false
			o.tsIsValid = true
			o.tsIsDirty = false
			o.ts.SetIsTimestamp(true)
		} else {
			panic("Wrong type found for ts.")
		}
	} else {
		o.tsIsValid = false
		o.tsIsNull = true
		o.ts = datetime.DateTime{}
	}
	if v, ok := m["test_int"]; ok {
		if v == nil {
			o.testInt = 5
			o.testIntIsNull = true
			o.testIntIsValid = true
			o.testIntIsDirty = false
		} else if o.testInt, ok = v.(int); ok {
			o.testIntIsNull = false
			o.testIntIsValid = true
			o.testIntIsDirty = false
		} else {
			panic("Wrong type found for test_int.")
		}
	} else {
		o.testIntIsValid = false
		o.testIntIsNull = true
		o.testInt = 5
	}
	if v, ok := m["test_float"]; ok {
		if v == nil {
			o.testFloat = 0.0
			o.testFloatIsNull = true
			o.testFloatIsValid = true
			o.testFloatIsDirty = false
		} else if o.testFloat, ok = v.(float32); ok {
			o.testFloatIsNull = false
			o.testFloatIsValid = true
			o.testFloatIsDirty = false
		} else {
			panic("Wrong type found for test_float.")
		}
	} else {
		o.testFloatIsValid = false
		o.testFloatIsNull = true
		o.testFloat = 0.0
	}
	if v, ok := m["test_text"]; ok {
		if v == nil {
			o.testText = ""
			o.testTextIsNull = true
			o.testTextIsValid = true
			o.testTextIsDirty = false
		} else if o.testText, ok = v.(string); ok {
			o.testTextIsNull = false
			o.testTextIsValid = true
			o.testTextIsDirty = false
		} else {
			panic("Wrong type found for test_text.")
		}
	} else {
		o.testTextIsValid = false
		o.testTextIsNull = true
		o.testText = ""
	}
	if v, ok := m["test_bit"]; ok {
		if v == nil {
			o.testBit = false
			o.testBitIsNull = true
			o.testBitIsValid = true
			o.testBitIsDirty = false
		} else if o.testBit, ok = v.(bool); ok {
			o.testBitIsNull = false
			o.testBitIsValid = true
			o.testBitIsDirty = false
		} else {
			panic("Wrong type found for test_bit.")
		}
	} else {
		o.testBitIsValid = false
		o.testBitIsNull = true
		o.testBit = false
	}
	if v, ok := m["test_varchar"]; ok {
		if v == nil {
			o.testVarchar = ""
			o.testVarcharIsNull = true
			o.testVarcharIsValid = true
			o.testVarcharIsDirty = false
		} else if o.testVarchar, ok = v.(string); ok {
			o.testVarcharIsNull = false
			o.testVarcharIsValid = true
			o.testVarcharIsDirty = false
		} else {
			panic("Wrong type found for test_varchar.")
		}
	} else {
		o.testVarcharIsValid = false
		o.testVarcharIsNull = true
		o.testVarchar = ""
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = map[string]interface{}(v.(db.ValueMap))
	}
	o._restored = true
}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *typeTestBase) Save(ctx context.Context) {
	if o._restored {
		o.Update(ctx)
	} else {
		o.Insert(ctx)
	}
}

// Update will update the values in the database, saving any changed values.
func (o *typeTestBase) Update(ctx context.Context) {
	if !o._restored {
		panic("Cannot update a record that was not originally read from the database.")
	}
	m := o.getModifiedFields()
	if len(m) == 0 {
		return
	}
	d := db.GetDatabase("goradd")
	d.Update(ctx, "type_test", m, "id", fmt.Sprint(o.id))
	o.resetDirtyStatus()
}

// Insert forces the object to be inserted into the database. If the object was loaded from the database originally,
// this will create a duplicate in the database.
func (o *typeTestBase) Insert(ctx context.Context) {
	m := o.getModifiedFields()
	if len(m) == 0 {
		return
	}
	d := db.GetDatabase("goradd")
	id := d.Insert(ctx, "type_test", m)
	o.id = id
	o.resetDirtyStatus()
	o._restored = true
}

func (o *typeTestBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}

	if o.dateIsDirty {
		if o.dateIsNull {
			fields["date"] = nil
		} else {
			fields["date"] = o.date.GoTime()
		}
	}

	if o.timeIsDirty {
		if o.timeIsNull {
			fields["time"] = nil
		} else {
			fields["time"] = o.time.GoTime()
		}
	}

	if o.dateTimeIsDirty {
		if o.dateTimeIsNull {
			fields["date_time"] = nil
		} else {
			fields["date_time"] = o.dateTime.GoTime()
		}
	}

	if o.tsIsDirty {
		if o.tsIsNull {
			fields["ts"] = nil
		} else {
			fields["ts"] = o.ts.GoTime()
		}
	}

	if o.testIntIsDirty {
		if o.testIntIsNull {
			fields["test_int"] = nil
		} else {
			fields["test_int"] = o.testInt
		}
	}

	if o.testFloatIsDirty {
		if o.testFloatIsNull {
			fields["test_float"] = nil
		} else {
			fields["test_float"] = o.testFloat
		}
	}

	if o.testTextIsDirty {
		if o.testTextIsNull {
			fields["test_text"] = nil
		} else {
			fields["test_text"] = o.testText
		}
	}

	if o.testBitIsDirty {
		if o.testBitIsNull {
			fields["test_bit"] = nil
		} else {
			fields["test_bit"] = o.testBit
		}
	}

	if o.testVarcharIsDirty {
		if o.testVarcharIsNull {
			fields["test_varchar"] = nil
		} else {
			fields["test_varchar"] = o.testVarchar
		}
	}

	return
}

func (o *typeTestBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.dateIsDirty = false
	o.timeIsDirty = false
	o.dateTimeIsDirty = false
	o.tsIsDirty = false
	o.testIntIsDirty = false
	o.testFloatIsDirty = false
	o.testTextIsDirty = false
	o.testBitIsDirty = false
	o.testVarcharIsDirty = false
}

// Delete deletes the associated record from the database.
func (o *typeTestBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "type_test", "id", o.id)
}

// deleteTypeTest deletes the associated record from the database.
func deleteTypeTest(ctx context.Context, pk string) {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "type_test", "id", pk)
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *typeTestBase) Get(key string) interface{} {

	switch key {
	case "ID":
		if !o.idIsValid {
			return nil
		}
		return o.id

	case "Date":
		if !o.dateIsValid {
			return nil
		}
		return o.date

	case "Time":
		if !o.timeIsValid {
			return nil
		}
		return o.time

	case "DateTime":
		if !o.dateTimeIsValid {
			return nil
		}
		return o.dateTime

	case "Ts":
		if !o.tsIsValid {
			return nil
		}
		return o.ts

	case "TestInt":
		if !o.testIntIsValid {
			return nil
		}
		return o.testInt

	case "TestFloat":
		if !o.testFloatIsValid {
			return nil
		}
		return o.testFloat

	case "TestText":
		if !o.testTextIsValid {
			return nil
		}
		return o.testText

	case "TestBit":
		if !o.testBitIsValid {
			return nil
		}
		return o.testBit

	case "TestVarchar":
		if !o.testVarcharIsValid {
			return nil
		}
		return o.testVarchar

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database object over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *typeTestBase) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err = encoder.Encode(o.id); err != nil {
		return
	}
	if err = encoder.Encode(o.idIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.idIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.date); err != nil {
		return
	}
	if err = encoder.Encode(o.dateIsNull); err != nil {
		return
	}
	if err = encoder.Encode(o.dateIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.dateIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.time); err != nil {
		return
	}
	if err = encoder.Encode(o.timeIsNull); err != nil {
		return
	}
	if err = encoder.Encode(o.timeIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.timeIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.dateTime); err != nil {
		return
	}
	if err = encoder.Encode(o.dateTimeIsNull); err != nil {
		return
	}
	if err = encoder.Encode(o.dateTimeIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.dateTimeIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.ts); err != nil {
		return
	}
	if err = encoder.Encode(o.tsIsNull); err != nil {
		return
	}
	if err = encoder.Encode(o.tsIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.tsIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.testInt); err != nil {
		return
	}
	if err = encoder.Encode(o.testIntIsNull); err != nil {
		return
	}
	if err = encoder.Encode(o.testIntIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.testIntIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.testFloat); err != nil {
		return
	}
	if err = encoder.Encode(o.testFloatIsNull); err != nil {
		return
	}
	if err = encoder.Encode(o.testFloatIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.testFloatIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.testText); err != nil {
		return
	}
	if err = encoder.Encode(o.testTextIsNull); err != nil {
		return
	}
	if err = encoder.Encode(o.testTextIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.testTextIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.testBit); err != nil {
		return
	}
	if err = encoder.Encode(o.testBitIsNull); err != nil {
		return
	}
	if err = encoder.Encode(o.testBitIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.testBitIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.testVarchar); err != nil {
		return
	}
	if err = encoder.Encode(o.testVarcharIsNull); err != nil {
		return
	}
	if err = encoder.Encode(o.testVarcharIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.testVarcharIsDirty); err != nil {
		return
	}

	if o._aliases == nil {
		if err = encoder.Encode(false); err != nil {
			return
		}
	} else {
		if err = encoder.Encode(true); err != nil {
			return
		}
		if err = encoder.Encode(o._aliases); err != nil {
			return
		}
	}

	if err = encoder.Encode(o._restored); err != nil {
		return
	}

	return
}

func (o *typeTestBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&o.id); err != nil {
		return
	}
	if err = dec.Decode(&o.idIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.date); err != nil {
		return
	}
	if err = dec.Decode(&o.dateIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.dateIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.dateIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.time); err != nil {
		return
	}
	if err = dec.Decode(&o.timeIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.timeIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.timeIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.dateTime); err != nil {
		return
	}
	if err = dec.Decode(&o.dateTimeIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.dateTimeIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.dateTimeIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.ts); err != nil {
		return
	}
	if err = dec.Decode(&o.tsIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.tsIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.tsIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.testInt); err != nil {
		return
	}
	if err = dec.Decode(&o.testIntIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.testIntIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.testIntIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.testFloat); err != nil {
		return
	}
	if err = dec.Decode(&o.testFloatIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.testFloatIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.testFloatIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.testText); err != nil {
		return
	}
	if err = dec.Decode(&o.testTextIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.testTextIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.testTextIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.testBit); err != nil {
		return
	}
	if err = dec.Decode(&o.testBitIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.testBitIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.testBitIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.testVarchar); err != nil {
		return
	}
	if err = dec.Decode(&o.testVarcharIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.testVarcharIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.testVarcharIsDirty); err != nil {
		return
	}

	var hasAliases bool
	if err = dec.Decode(&hasAliases); err != nil {
		return
	}
	if hasAliases {
		if err = dec.Decode(&o._aliases); err != nil {
			return
		}
	}

	if err = dec.Decode(&o._restored); err != nil {
		return
	}

	return err
}
