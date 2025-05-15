package setx

import (
	"cmp"
	"fmt"

	"github.com/msales/gox/gofp"
)

type Sum[Value cmp.Ordered] struct {
	Result[Value]
	immutable bool
}

func NewSum[Value cmp.Ordered]() Sum[Value] {
	return Sum[Value]{
		Result: Result[Value]{
			Operator: All,
			Values:   []Value{},
		},
	}
}

// Merge merges a condition to existing result - check package description for algorithm description
func (r *Sum[Value]) Merge(operator Operator, values ...Value) (MergeResult, error) {
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

func (r *Sum[Value]) mergeAll(values []Value) (MergeResult, error) {
	if len(values) > 0 {
		return Error, fmt.Errorf("cannot merge any values when using all operator")
	}
	return r.makeImmutable()
}

func (r *Sum[Value]) mergeNone(values []Value) (MergeResult, error) {
	if len(values) > 0 {
		return Error, fmt.Errorf("cannot merge any values when using none operator")
	}

	switch r.Operator {
	case In:
		return Continue, nil
	case NotIn:
		return Continue, nil
	case All:
		// if the result is mutable and operator is ALL, this means it was not used yet
		// with first usage, operator and values are set
		r.Operator = None
		return Continue, nil
	case None:
		return Continue, nil
	default:
		return Error, fmt.Errorf("unknown operator %d in result", r.Operator)

	}
}

func (r *Sum[Value]) mergeIn(values []Value) (MergeResult, error) {
	switch r.Operator {
	case In:
		// operator was IN and merging IN => IN
		// so the resulting list is just a sum of both
		r.Result.merge(values...)
		return Continue, nil
	case NotIn:
		// operator was NOT_IN and merging IN => NOT IN
		// since one of the operators is NOT in the result must be NOT IN too
		// if there are any elements in the NOT IN set that are not in IN set then those
		// elements are in resulting NOT IN set
		r.Values = r.mergeInNotIn(values, r.Values)
		// it is already NOT IN operator
		if len(r.Values) == 0 {
			// If resulting set is empty (NOT IN is completely contained in IN condition)
			// then this pair of conditions makes the result always true, so ALL and immutable.
			return r.makeImmutable()
		}
		return Continue, nil
	case All:
		// if the result is mutable and operator is ALL, this means it was not used yet
		// with first usage, operator and values are set
		r.Operator = In
		r.Values = values
		return Continue, nil
	case None:
		// so far result was none, but NONE result added to any result becomes that any result
		r.Operator = In
		r.Values = values
		return Continue, nil
	default:
		return Error, fmt.Errorf("unknown operator %d in result", r.Operator)
	}
}

func (r *Sum[Value]) mergeNotIn(values []Value) (MergeResult, error) {
	switch r.Operator {
	case In:
		// operator was NOT_IN and merging IN => NOT IN
		// since one of the operators is NOT in the result must be NOT IN too
		// if there are any elements in the NOT IN set that are not in IN set then those
		// elements are in resulting NOT IN set
		r.Values = r.mergeInNotIn(r.Values, values)
		r.Operator = NotIn
		if len(r.Values) == 0 {
			// If resulting set is empty (NOT IN is completely contained in IN condition)
			// then this pair of conditions makes the result always true, so ALL and immutable.
			return r.makeImmutable()
		}
		return Continue, nil
	case NotIn:
		// operator was NOT IN and merging NOT IN => NOT IN
		// so the resulting list is just an intersection of both, in case there is no intersection, all values
		// are possible and result becomes immutable
		r.Result.intersect(values...)
		if len(r.Values) > 0 {
			return Continue, nil
		}
		return r.makeImmutable()
	case All:
		// if the result is mutable and operator is ALL, this means it was not used yet
		// with first usage, operator and values are set
		r.Operator = NotIn
		r.Values = values
		return Continue, nil
	case None:
		// so far result was none, but NONE result added to any result becomes that any result
		r.Operator = NotIn
		r.Values = values
		return Continue, nil
	default:
		return Error, fmt.Errorf("unknown operator %d in result", r.Operator)

	}
}

// makeImmutable changes the result to immutable with all possible values.
// Once this state is reached it can't be changed.
// If conditions so far make the result completely open, no other condition can narrow it anymore
// because all conditions are connected with OR. ANY OR SOMETHING => ANY.
func (r *Sum[Value]) makeImmutable() (MergeResult, error) {
	r.Values = []Value{}
	r.immutable = true
	r.Operator = All
	return Immutable, nil
}

func (r *Sum[Value]) mergeInNotIn(inValues, notInValues []Value) []Value {
	diff := gofp.DiffRight(inValues, notInValues)
	return diff
}
