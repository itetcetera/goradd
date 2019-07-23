package control

import (
	"fmt"
	"github.com/goradd/goradd/pkg/config"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// IDer is an object that can embed a list.
type IDer interface {
	ID() string
}

// IDSetter is an interface for an item that sets an id.
type IDSetter interface {
	SetID(id string)
}

// ItemListI is the interface for all controls that display a list of ListItems.
type ItemListI interface {
	AddItem(label string, value ...interface{}) ListItemI
	AddItemAt(index int, label string, value ...interface{})
	AddListItemAt(index int, item ListItemI)
	AddListItems(items ...interface{})
	GetItemAt(index int) ListItemI
	ListItems() []ListItemI
	Clear()
	RemoveItemAt(index int)
	Len() int
	GetItem(id string) (foundItem ListItemI)
	GetItemByValue(value interface{}) (id string, foundItem ListItemI)
	reindex(start int)
	findItemByValue(value interface{}) (container *ItemList, index int)
}

// ItemList manages a list of ListItemI list items. ItemList is designed to be embedded in another structure, and will
// turn that object into a manager of list items. ItemList will manage the id's of the items in its list, you do not
// have control of that, and it needs to manage those ids in order to efficiently manage the selection process.
type ItemList struct {
	owner IDer
	items []ListItemI
}

// NewItemList creates a new item list. "owner" is the object that has the list embedded in it, and must be
// an IDer.
func NewItemList(owner IDer) ItemList {
	return ItemList{owner: owner}
}

// AddItem adds the given item to the end of the list. The value is optional, but should only be one or zero values.
func (l *ItemList) AddItem(label string, value ...interface{}) ListItemI {
	i := NewListItem(label, value...)
	l.AddListItemAt(len(l.items), i)
	return i
}

// AddItemAt adds the item at the given index.
// If the index is negative, it counts from the end. -1 would therefore put the item before the last item.
// If the index is bigger or equal to the number of items, it adds it to the end. If the index is zero, or is negative and smaller than
// the negative value of the number of items, it adds to the beginning. This can be an expensive operation in a long
// hierarchical list, so use sparingly.
func (l *ItemList) AddItemAt(index int, label string, value ...interface{}) {
	l.AddListItemAt(index, NewListItem(label, value...))
}

// AddItemAt adds the item at the given index. If the index is negative, it counts from the end. If the index is
// -1 or bigger than the number of items, it adds it to the end. If the index is zero, or is negative and smaller than
// the negative value of the number of items, it adds to the beginning. This can be an expensive operation in a long
// hierarchical list, so use sparingly.
func (l *ItemList) AddListItemAt(index int, item ListItemI) {
	if index < 0 {
		index = len(l.items) + index
		if index < 0 {
			index = 0
		}
	} else if index > len(l.items) {
		index = len(l.items)
	}
	l.items = append(l.items, nil)
	copy(l.items[index+1:], l.items[index:])
	l.items[index] = item
	l.reindex(index)
}

// AddListItems adds one or more objects to the end of the list. items should be a list of ListItemI,
// ItemLister, ItemIDer, Labeler or Stringer types. This function can accept one or more lists of items, or
// single items
func (l *ItemList) AddListItems(items ...interface{}) {
	var start int

	if items == nil {
		return
	}
	for _, item := range items {
		kind := reflect.TypeOf(item).Kind()
		if kind == reflect.Array || kind == reflect.Slice {
			listValue := reflect.ValueOf(item)
			for i := 0; i < listValue.Len(); i++ {
				itemI := listValue.Index(i).Interface()
				l.addListItem(itemI)
			}
		} else {
			l.addListItem(item)
		}
	}
	l.reindex(start)
}

// Private function to add an interface item to the end of the list. Will need to be reindexed eventually.
func (l *ItemList) addListItem(item interface{}) {
	switch v := item.(type) {
	case ListItemI:
		l.items = append(l.items, v)
	case ItemLister:
		item := NewItemFromItemLister(v)
		l.items = append(l.items, item)
	case ItemIDer:
		item := NewItemFromItemIDer(v)
		l.items = append(l.items, item)
	case Labeler:
		item := NewItemFromLabeler(v)
		l.items = append(l.items, item)
	case fmt.Stringer:
		item := NewItemFromStringer(v)
		l.items = append(l.items, item)
	default:
		panic("Unknown object type")
	}
}

// reindex is internal and should get called whenever an item gets added to the list out of order or an id changes.
func (l *ItemList) reindex(start int) {
	if l.owner.ID() == "" || l.items == nil || len(l.items) == 0 || start >= len(l.items) {
		return
	}
	for i := start; i < len(l.items); i++ {
		id := l.owner.ID() + "_" + strconv.Itoa(i)
		l.items[i].SetID(id)
	}
}

// GetItemAt retrieves an item by index.
func (l *ItemList) GetItemAt(index int) ListItemI {
	if index >= len(l.items) {
		return nil
	}
	return l.items[index]
}

// ListItems returns a slice of the ListItemI items, in the order they were added or arranged.
func (l *ItemList) ListItems() []ListItemI {
	return l.items
}

// Clear removes all the items from the list.
func (l *ItemList) Clear() {
	l.items = nil
}

// RemoveItemAt removes an item at the given index.
func (l *ItemList) RemoveItemAt(index int) {
	if index < 0 || index >= len(l.items) {
		panic("Index out of range.")
	}
	l.items = append(l.items[:index], l.items[index+1:]...)
}

// Len returns the length of the item list at the current level. In other words, it does not count items in sublists.
func (l *ItemList) Len() int {
	if l.items == nil {
		return 0
	}
	return len(l.items)
}

// GetItem recursively searches for and returns the item corresponding to the given id. Since we are managing the
// id numbers, we can efficiently find the item. Note that if you add items to the list, the ids may change.
func (l *ItemList) GetItem(id string) (foundItem ListItemI) {
	if l.items == nil {
		return nil
	}

	parts := strings.SplitN(id, "_", 3) // first item is our own id, 2nd is id from the list, 3rd is a level beyond the list
	l1Id, err := strconv.Atoi(parts[1])
	if err != nil || l1Id < 0 {
		panic("Bad id")
	}

	var countParts int
	if countParts = len(parts); countParts <= 1 || l1Id >= len(l.items) {
		return nil
	}

	item := l.items[l1Id]

	if countParts == 2 {
		return item
	}

	return item.GetItem(parts[1] + "_" + parts[2])
}

// GetItemByValue recursively searches the list to find the item with the given value.
// It starts with the current list, and if not found, will search in sublists.
func (l *ItemList) GetItemByValue(value interface{}) (id string, foundItem ListItemI) {
	container, index := l.findItemByValue(value)

	if container != nil {
		foundItem = container.items[index]
		id = foundItem.ID()
		return
	}
	return "", nil
}

// findItemByValue searches for the item by value, and returns the index of the found item,
// and the ItemList that the item was found in. The returned ItemList could be the current
// item list, or a sublist.
func (l *ItemList) findItemByValue(value interface{}) (container *ItemList, index int) {
	if len(l.items) == 0 {
		return nil, -1 // no sub items, so its not here
	}
	var item ListItemI

	for index, item = range l.items {
		v := item.Value()
		if v == value {
			container = l
			return
		}
	}

	for index, item = range l.items {
		container, index = item.findItemByValue(value)
		if container != nil {
			return
		}
	}

	return nil, -1 // not found
}

// SortIds sorts a list of auto-generated ids in numerical and hierarchical order.
// This is normally just called by the framework.
func SortIds(ids []string) {
	if len(ids) > 1 {
		sort.Sort(IdSlice(ids))
	}
}

// IdSlice is a slice of string ids, and is used to sort a list of ids
// that the item list uses.
type IdSlice []string

func (p IdSlice) Len() int { return len(p) }
func (p IdSlice) Less(i, j int) bool {
	// First ones are always the main control id, and should be equal
	vals1 := strings.SplitN(p[i], "_", 2)
	vals2 := strings.SplitN(p[j], "_", 2)
	if vals1[0] != vals2[0] {
		panic("The first part of an id should be equal when sorting.")
	}

	for {
		vals1 = strings.SplitN(vals1[1], "_", 2)
		vals2 = strings.SplitN(vals2[1], "_", 2)
		i1, _ := strconv.Atoi(vals1[0])
		i2, _ := strconv.Atoi(vals2[0])
		if i1 < i2 {
			return true
		} else if i1 > i2 {
			return false
		} else if len(vals1) < len(vals2) {
			return true
		} else if len(vals1) > len(vals2) || len(vals2) <= 1 {
			return false
		}
	}
}

func (p IdSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }


// NoSelectionItemList returns a default item list to start a selection list that allows no selection
func NoSelectionItemList() []interface{} {
	return []interface{}{NewListItem(config.NoSelectionString, nil)}
}

// SelectOneItemList returns a default item list to start a selection list that asks the user to select an item
func SelectOneItemList() []interface{} {
	return []interface{}{NewListItem(config.SelectOneString, nil)}
}