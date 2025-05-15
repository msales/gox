package setx

import (
	"cmp"
	"fmt"

	"github.com/msales/gox/gofp"
)

type Product[Value cmp.Ordered] struct {
	Result[Value]
	immutable bool
}

func NewProduct[Value cmp.Ordered]() Product[Value] {
	return Product[Value]{
		Result: Result[Value]{
			Operator: All,
			Values:   []Value{},
		},
	}
}

// Merge merges a condition to existing result - check package description for algorithm description
func (r *Product[Value]) Merge(operator Operator, values ...Value) (MergeResult, error) {
	if r.immutable {
		return Immutable, nil
	}

	switch operator {
	case All:
		return r.mergeAll(values)
	case None:
		return r.mergeNone(values)
	case In:
		return r.mergeIn(values)
	case NotIn:
		return r.mergeNotIn(values)
	default:
		return Error, fmt.Errorf("unknown operator")
	}
}

func (r *Product[Value]) mergeNone(values []Value) (MergeResult, error) {
	if len(values) > 0 {
		return Error, fmt.Errorf("cannot merge any values when using none operator")
	}
	return r.makeImmutable()
}

func (r *Product[Value]) mergeAll(values []Value) (MergeResult, error) {
	if len(values) > 0 {
		return Error, fmt.Errorf("cannot merge any values when using all operator")
	}
	// ToDo: implement when needed - the case we have will never use this
	return r.makeImmutable()
}

func (r *Product[Value]) mergeIn(values []Value) (MergeResult, error) {
	switch r.Operator {
	case In:
		// operator was IN and merging IN => IN
		// so the resulting list is just an intersection of both
		r.Result.intersect(values...)
		if len(r.Values) == 0 {
			return r.makeImmutable()
		}
		return Continue, nil
	case NotIn:
		// operator was NOT_IN and merging IN => IN
		r.Values = r.mergeInNotIn(r.Values, values)
		r.Operator = In
		if len(r.Values) == 0 {
			// If resulting set is empty (IN is completely contained in NOT IN condition)
			// then this pair of conditions makes the result always false, so NONE and immutable.
			return r.makeImmutable()
		}
		return Continue, nil
	case All:
		// if the result is mutable and operator is ALL, this means it was not used yet
		// with first usage, operator and values are set
		r.Operator = In
		r.Values = values
		return Continue, nil
	default:
		return Error, fmt.Errorf("unknown operator in result")
	}
}

func (r *Product[Value]) mergeNotIn(values []Value) (MergeResult, error) {
	switch r.Operator {
	case In:
		// operator was IN and merging NOT IN => IN
		r.Values = r.mergeInNotIn(values, r.Values)
		// the operator was already IN
		if len(r.Values) == 0 {
			// If resulting set is empty (IN is completely contained in NOT IN condition)
			// then this pair of conditions makes the result always false, so NONE and immutable.
			return r.makeImmutable()
		}
		return Continue, nil
	case NotIn:
		// operator was NOT IN and merging NOT IN => NOT IN
		// Result is a sum of both sets.
		r.Result.merge(values...)
		return Continue, nil
	case All:
		// if the result is mutable and operator is ALL, this means it was not used yet
		// with first usage, operator and values are set
		r.Operator = NotIn
		r.Values = values
		return Continue, nil
	default:
		return Error, fmt.Errorf("unknown operator in result")
	}
}

// makeImmutable changes the result to immutable with all possible values.
// Once this state is reached it can't be changed.
// If conditions so far make the result completely open, no other condition can narrow it anymore
// because all conditions are connected with OR. ANY OR SOMETHING => ANY.
func (r *Product[Value]) makeImmutable() (MergeResult, error) {
	r.Values = []Value{}
	r.immutable = true
	r.Operator = None
	return Immutable, nil
}

func (r *Product[Value]) mergeInNotIn(inValues, notInValues []Value) []Value {
	diff := gofp.DiffRight(inValues, notInValues)
	return diff
}
