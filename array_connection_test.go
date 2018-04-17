package relay_test

import (
	"reflect"
	"testing"

	"github.com/graphql-go/graphql/testutil"
	"github.com/stratumn/relay"
)

var arrayListTestLetters = []interface{}{
	"A", "B", "C", "D", "E",
}

func TestListFromArray_HandlesBasicSlicing_ReturnsAllElementsWithoutFilters(t *testing.T) {
	args := relay.NewListArguments(nil)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesBasicSlicing_RespectsASmallerFirst(t *testing.T) {
	// Create list arguments from map[string]interface{},
	// which you usually get from types.GQLParams.Args
	filter := map[string]interface{}{
		"first": 2,
	}
	args := relay.NewListArguments(filter)

	// Alternatively, you can create list arg the following way.
	// args := relay.NewListArguments(filter)
	// args.First = 2

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjE=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesBasicSlicing_RespectsAnOverlyLargeFirst(t *testing.T) {

	filter := map[string]interface{}{
		"first": 10,
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesBasicSlicing_RespectsASmallerLast(t *testing.T) {

	filter := map[string]interface{}{
		"last": 2,
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjM=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: true,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesBasicSlicing_RespectsAnOverlyLargeLast(t *testing.T) {

	filter := map[string]interface{}{
		"last": 10,
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}

func TestListFromArray_HandlesPagination_RespectsFirstAndAfter(t *testing.T) {

	filter := map[string]interface{}{
		"first": 2,
		"after": "YXJyYXljb25uZWN0aW9uOjE=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjI=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesPagination_RespectsFirstAndAfterWithLongFirst(t *testing.T) {

	filter := map[string]interface{}{
		"first": 10,
		"after": "YXJyYXljb25uZWN0aW9uOjE=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjI=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesPagination_RespectsLastAndBefore(t *testing.T) {
	filter := map[string]interface{}{
		"last":   2,
		"before": "YXJyYXljb25uZWN0aW9uOjM=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjI=",
			HasPreviousPage: true,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesPagination_RespectsLastAndBeforeWithLongLast(t *testing.T) {
	filter := map[string]interface{}{
		"last":   10,
		"before": "YXJyYXljb25uZWN0aW9uOjM=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjI=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesPagination_RespectsFirstAndAfterAndBefore_TooFew(t *testing.T) {
	filter := map[string]interface{}{
		"first":  2,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjI=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesPagination_RespectsFirstAndAfterAndBefore_TooMany(t *testing.T) {
	filter := map[string]interface{}{
		"first":  4,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesPagination_RespectsFirstAndAfterAndBefore_ExactlyRight(t *testing.T) {
	filter := map[string]interface{}{
		"first":  3,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesPagination_RespectsLastAndAfterAndBefore_TooFew(t *testing.T) {
	filter := map[string]interface{}{
		"last":   2,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjI=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: true,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesPagination_RespectsLasttAndAfterAndBefore_TooMany(t *testing.T) {
	filter := map[string]interface{}{
		"last":   4,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesPagination_RespectsLastAndAfterAndBefore_ExactlyRight(t *testing.T) {
	filter := map[string]interface{}{
		"last":   3,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}

func TestListFromArray_HandlesCursorEdgeCases_ReturnsNoElementsIfFirstIsZero(t *testing.T) {
	filter := map[string]interface{}{
		"first": 0,
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{},
		PageInfo: relay.PageInfo{
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesCursorEdgeCases_ReturnsAllElementsIfCursorsAreInvalid(t *testing.T) {
	filter := map[string]interface{}{
		"before": "invalid",
		"after":  "invalid",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesCursorEdgeCases_ReturnsAllElementsIfCursorsAreOnTheOutside(t *testing.T) {
	filter := map[string]interface{}{
		"before": "YXJyYXljb25uZWN0aW9uOjYK",     // ==> offset: int(6)
		"after":  "YXJyYXljb25uZWN0aW9uOi0xCg==", // ==> offset: int(-1)
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}

func TestListFromArray_HandlesCursorEdgeCases_ReturnsNullIfCursorsIsConsecutive(t *testing.T) {
	filter := map[string]interface{}{
		"before": "YXJyYXljb25uZWN0aW9uOjM=", // ==> offset: int(3)
		"after":  "YXJyYXljb25uZWN0aW9uOjI=", // ==> offset: int(2)
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges:    []*relay.Edge{},
		PageInfo: relay.PageInfo{},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_HandlesCursorEdgeCases_ReturnsNoElementsIfCursorsCross(t *testing.T) {
	filter := map[string]interface{}{
		"before": "YXJyYXljb25uZWN0aW9uOjI=", // ==> offset: int(2)
		"after":  "YXJyYXljb25uZWN0aW9uOjQ=", // ==> offset: int(4)
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges:    []*relay.Edge{},
		PageInfo: relay.PageInfo{},
	}

	result := relay.ListFromArray(arrayListTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
func TestListFromArray_CursorForObjectInList_ReturnsAnEdgeCursor_GivenAnArrayAndAMemberObject(t *testing.T) {
	letterBCursor := relay.CursorForObjectInList(arrayListTestLetters, "B")
	expected := relay.ListCursor("YXJyYXljb25uZWN0aW9uOjE=")
	if !reflect.DeepEqual(letterBCursor, expected) {
		t.Fatalf("wrong result, cursor result diff: %v", testutil.Diff(expected, letterBCursor))
	}
}
func TestListFromArray_CursorForObjectInList_ReturnsEmptyCursor_GivenAnArrayAndANonMemberObject(t *testing.T) {
	letterFCursor := relay.CursorForObjectInList(arrayListTestLetters, "F")
	if letterFCursor != "" {
		t.Fatalf("wrong result, expected empty cursor, got: %v", letterFCursor)
	}
}

func TestListFromArraySlice_JustRightArraySlice(t *testing.T) {
	filter := map[string]interface{}{
		"first": 2,
		"after": "YXJyYXljb25uZWN0aW9uOjA=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjI=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := relay.ListFromArraySlice(
		arrayListTestLetters[1:3],
		args,
		relay.ArraySliceMetaInfo{
			SliceStart:  1,
			ArrayLength: 5,
		},
	)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}

func TestListFromArraySlice_OversizedSliceLeft(t *testing.T) {
	filter := map[string]interface{}{
		"first": 2,
		"after": "YXJyYXljb25uZWN0aW9uOjA=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjI=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := relay.ListFromArraySlice(
		arrayListTestLetters[0:3],
		args,
		relay.ArraySliceMetaInfo{
			SliceStart:  0,
			ArrayLength: 5,
		},
	)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}

func TestListFromArraySlice_OversizedSliceRight(t *testing.T) {
	filter := map[string]interface{}{
		"first": 1,
		"after": "YXJyYXljb25uZWN0aW9uOjE=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjI=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjI=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := relay.ListFromArraySlice(
		arrayListTestLetters[2:4],
		args,
		relay.ArraySliceMetaInfo{
			SliceStart:  2,
			ArrayLength: 5,
		},
	)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}

func TestListFromArraySlice_OversizedSliceBoth(t *testing.T) {
	filter := map[string]interface{}{
		"first": 1,
		"after": "YXJyYXljb25uZWN0aW9uOjE=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjI=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjI=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := relay.ListFromArraySlice(
		arrayListTestLetters[1:4],
		args,
		relay.ArraySliceMetaInfo{
			SliceStart:  1,
			ArrayLength: 5,
		},
	)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}

func TestListFromArraySlice_UndersizedSliceLeft(t *testing.T) {
	filter := map[string]interface{}{
		"first": 3,
		"after": "YXJyYXljb25uZWN0aW9uOjE=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjM=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := relay.ListFromArraySlice(
		arrayListTestLetters[3:5],
		args,
		relay.ArraySliceMetaInfo{
			SliceStart:  3,
			ArrayLength: 5,
		},
	)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}

func TestListFromArraySlice_UndersizedSliceRight(t *testing.T) {
	filter := map[string]interface{}{
		"first": 3,
		"after": "YXJyYXljb25uZWN0aW9uOjE=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjI=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := relay.ListFromArraySlice(
		arrayListTestLetters[2:4],
		args,
		relay.ArraySliceMetaInfo{
			SliceStart:  2,
			ArrayLength: 5,
		},
	)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}

func TestListFromArraySlice_UndersizedSliceBoth(t *testing.T) {
	filter := map[string]interface{}{
		"first": 3,
		"after": "YXJyYXljb25uZWN0aW9uOjE=",
	}
	args := relay.NewListArguments(filter)

	expected := &relay.List{
		Edges: []*relay.Edge{
			&relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjM=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := relay.ListFromArraySlice(
		arrayListTestLetters[3:4],
		args,
		relay.ArraySliceMetaInfo{
			SliceStart:  3,
			ArrayLength: 5,
		},
	)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, list result diff: %v", testutil.Diff(expected, result))
	}
}
