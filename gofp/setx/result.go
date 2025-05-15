// Package setx calculates value for approximated condition for product (A AND B AND C) and sum (A OR B OR C)
//
// SUM
// Input:
// There is a list of conditions related to the same value type, where all conditions are summed up:
// A OR B OR C
// Each condition has an operator IN, NOT, ALL.
// Example expressions:
// country IN ("PL")
// country NOT IN  ("PL")
// country IN ("PL") OR country IN ("US")
// country IN ("PL") OR country NOT IN ("US")
// country IN ("PL") OR ANY country // operator ALL
// The algorithm leverages the fact, that each time we can calculate a pair of conditions and use the result to
// calculate result with next one:
// A OR B OR C OR D -> ((A OR B) OR C) OR D
// Algorithm allows merging consecutive conditions to a result and each time calculates new operator and list of values.
// Each time Merge is called it returns MergeResult which may be used to optimise processing. It is enough that at
// any point of processing the list of values becomes ALL to ensure that no consecutive condition can change that
// because they are all connected with OR. Consecutive conditions can extend the possible list of resulting values
// but never limit them.
// If at any point of processing result becomes ALL values, MergeResult is set to Immutable and no additional Merge
// call will change that. In that case additional Merge calls may be skipped.
//
// PRODUCT
// ...
package setx

import (
	"cmp"
	"slices"

	"github.com/msales/gox/gofp"
)

type MergeResult int
type Operator int

const (
	Error MergeResult = iota
	Continue
	Immutable
)

const (
	In Operator = iota
	NotIn
	All
	None
)

func (o Operator) String() string {
	switch o {
	case In:
		return "IN"
	case NotIn:
		return "NOT_IN"
	case All:
		return "ALL"
	default:
		return "not implemented"
	}
}

type Result[Value cmp.Ordered] struct {
	Operator Operator
	Values   []Value
}

func (r *Result[Value]) merge(values ...Value) {
	r.Values = append(r.Values, values...)
	slices.Sort(r.Values)
	r.Values = slices.Compact(r.Values)
}

func (r *Result[Value]) intersect(values ...Value) {
	r.Values = gofp.Intersection(r.Values, values, func(elementOne, elementTwo Value) bool {
		return elementOne == elementTwo
	})
}
