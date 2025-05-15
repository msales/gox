package setx_test

import (
	"testing"

	"github.com/msales/gox/gofp/setx"
	"github.com/stretchr/testify/assert"
)

type sumConditions struct {
	operator setx.Operator
	values   []string
}
type sumTestCase struct {
	name            string
	conditions      []sumConditions
	wantResult      setx.Sum[string]
	wantMergeResult setx.MergeResult
	wantErr         bool
}

func TestResult_MergeSum(t *testing.T) {
	tests := sumTestCases()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := setx.NewSum[string]()
			var mergeResult setx.MergeResult
			var err error
			for _, condition := range tt.conditions {
				mergeResult, err = result.Merge(condition.operator, condition.values...)
			}
			assert.Equal(t, tt.wantMergeResult, mergeResult)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.wantResult.Operator, result.Operator)
			assert.ElementsMatch(t, tt.wantResult.Values, result.Values)
		})
	}
}

func TestResult_MergeSumWithAll(t *testing.T) {
	tests := sumTestCases()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// any test no error case if there is one more condition with all added to the conditions list, should result with
			// immutable result with ALL operator
			tt.conditions = append(tt.conditions, addToSumMerge(setx.All))
			tt.wantMergeResult = setx.Immutable
			tt.wantResult.Operator = setx.All
			tt.wantResult.Values = []string{}
			if tt.wantErr {
				return
			}

			result := setx.NewSum[string]()
			var mergeResult setx.MergeResult
			var err error
			for _, condition := range tt.conditions {
				mergeResult, err = result.Merge(condition.operator, condition.values...)
			}
			assert.Equal(t, tt.wantMergeResult, mergeResult)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.wantResult.Operator, result.Operator)
			assert.ElementsMatch(t, tt.wantResult.Values, result.Values)
		})
	}
}

// sumTestCases is providing cases for 2 tests
// Unfortunately in goland it prevents running separate case test, but nevertheless it's better than
// copy-paste of all the cases.
func sumTestCases() []sumTestCase {
	tests := []sumTestCase{
		{
			name: "One IN condition",
			conditions: []sumConditions{
				addToSumMerge(setx.In, "PL", "ES"),
			},
			wantResult:      sumResult(setx.In, "PL", "ES"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "One NOT IN condition",
			conditions: []sumConditions{
				addToSumMerge(setx.NotIn, "PL", "ES"),
			},
			wantResult:      sumResult(setx.NotIn, "PL", "ES"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "Two IN conditions of disjoint sets",
			conditions: []sumConditions{
				addToSumMerge(setx.In, "PL", "ES"),
				addToSumMerge(setx.In, "DE", "US"),
			},
			wantResult:      sumResult(setx.In, "PL", "ES", "US", "DE"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "Two NOT IN conditions of disjoint sets",
			conditions: []sumConditions{
				addToSumMerge(setx.NotIn, "PL", "ES"),
				addToSumMerge(setx.NotIn, "DE", "US"),
			},
			wantResult:      sumResult(setx.All),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},

		{
			name: "Two IN conditions of intersecting sets",
			conditions: []sumConditions{
				addToSumMerge(setx.In, "PL", "ES"),
				addToSumMerge(setx.In, "PL", "US"),
			},
			wantResult:      sumResult(setx.In, "PL", "ES", "US"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "Two NOT IN conditions of intersecting sets",
			conditions: []sumConditions{
				addToSumMerge(setx.NotIn, "PL", "ES"),
				addToSumMerge(setx.NotIn, "PL", "US"),
			},
			wantResult:      sumResult(setx.NotIn, "PL"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "Three NOT IN conditions of intersecting sets, but no value in all conditions",
			conditions: []sumConditions{
				addToSumMerge(setx.NotIn, "PL", "ES", "UK"),
				addToSumMerge(setx.NotIn, "PL", "US", "UK"),
				addToSumMerge(setx.NotIn, "IS", "US", "CZ"),
			},
			wantResult:      sumResult(setx.All),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
		{
			name: "Three IN conditions of intersecting sets, but no value in all conditions",
			conditions: []sumConditions{
				addToSumMerge(setx.In, "PL", "ES", "UK"),
				addToSumMerge(setx.In, "PL", "US", "UK"),
				addToSumMerge(setx.In, "IS", "US", "CZ"),
			},
			wantResult:      sumResult(setx.In, "PL", "ES", "UK", "US", "IS", "CZ"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "Three NOT IN conditions of intersecting sets but at least one value in all conditions",
			conditions: []sumConditions{
				addToSumMerge(setx.NotIn, "PL", "CA", "ES", "UK", "SK"),
				addToSumMerge(setx.NotIn, "PL", "SK", "US", "UK", "CA"),
				addToSumMerge(setx.NotIn, "CA", "IS", "US", "SK", "CZ"),
			},
			wantResult:      sumResult(setx.NotIn, "CA", "SK"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "Three IN conditions of intersecting sets but at least one value in all conditions",
			conditions: []sumConditions{
				addToSumMerge(setx.In, "PL", "CA", "ES", "UK", "SK"),
				addToSumMerge(setx.In, "PL", "SK", "US", "UK", "CA"),
				addToSumMerge(setx.In, "CA", "IS", "US", "SK", "CZ"),
			},
			wantResult:      sumResult(setx.In, "PL", "CA", "ES", "UK", "SK", "US", "IS", "CZ"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "IN then NOT IN intersecting conditions",
			conditions: []sumConditions{
				addToSumMerge(setx.In, "PL", "ES"),
				addToSumMerge(setx.NotIn, "PL", "US"),
			},
			wantResult:      sumResult(setx.NotIn, "US"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "NOT IN then IN intersecting conditions",
			conditions: []sumConditions{
				addToSumMerge(setx.NotIn, "PL", "US"),
				addToSumMerge(setx.In, "PL", "ES"),
			},
			wantResult:      sumResult(setx.NotIn, "US"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "IN then NOT IN disjoint conditions",
			conditions: []sumConditions{
				addToSumMerge(setx.In, "PL", "ES"),
				addToSumMerge(setx.NotIn, "CZ", "US"),
			},
			wantResult:      sumResult(setx.NotIn, "CZ", "US"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "NOT IN then IN disjoint conditions",
			conditions: []sumConditions{
				addToSumMerge(setx.NotIn, "CZ", "US"),
				addToSumMerge(setx.In, "PL", "ES"),
			},
			wantResult:      sumResult(setx.NotIn, "CZ", "US"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "IN then NOT IN identical conditions",
			conditions: []sumConditions{
				addToSumMerge(setx.In, "PL"),
				addToSumMerge(setx.NotIn, "PL"),
			},
			wantResult:      sumResult(setx.All),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
		{
			name: "NOT IN then IN identical conditions",
			conditions: []sumConditions{
				addToSumMerge(setx.NotIn, "PL"),
				addToSumMerge(setx.In, "PL"),
			},
			wantResult:      sumResult(setx.All),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},

		{
			name: "IN then NOT IN conditions, NOT IN set contained in IN set",
			conditions: []sumConditions{
				addToSumMerge(setx.In, "PL", "US"),
				addToSumMerge(setx.NotIn, "PL"),
			},
			wantResult:      sumResult(setx.All),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
		{
			name: "NOT IN then IN conditions, NOT IN set contained in IN set",
			conditions: []sumConditions{
				addToSumMerge(setx.NotIn, "PL"),
				addToSumMerge(setx.In, "PL", "US"),
			},
			wantResult:      sumResult(setx.All),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},

		{
			name: "wrong ALL condition with values",
			conditions: []sumConditions{
				addToSumMerge(setx.All, "PL"),
			},
			wantErr: true,
		},
		{
			name: "wrong NONE condition with values",
			conditions: []sumConditions{
				addToSumMerge(setx.None, "PL"),
			},
			wantErr: true,
		},

		{
			name: "only NONE",
			conditions: []sumConditions{
				addToSumMerge(setx.None),
			},
			wantResult:      sumResult(setx.None),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "twice NONE",
			conditions: []sumConditions{
				addToSumMerge(setx.None),
				addToSumMerge(setx.None),
			},
			wantResult:      sumResult(setx.None),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "IN + NONE",
			conditions: []sumConditions{
				addToSumMerge(setx.In, "PL"),
				addToSumMerge(setx.None),
			},
			wantResult:      sumResult(setx.In, "PL"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "NONE + IN",
			conditions: []sumConditions{
				addToSumMerge(setx.None),
				addToSumMerge(setx.In, "PL"),
			},
			wantResult:      sumResult(setx.In, "PL"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "NOT IN + NONE",
			conditions: []sumConditions{
				addToSumMerge(setx.NotIn, "PL"),
				addToSumMerge(setx.None),
			},
			wantResult:      sumResult(setx.NotIn, "PL"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "NONE + IN",
			conditions: []sumConditions{
				addToSumMerge(setx.None),
				addToSumMerge(setx.NotIn, "PL"),
			},
			wantResult:      sumResult(setx.NotIn, "PL"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "ALL + NONE",
			conditions: []sumConditions{
				addToSumMerge(setx.All),
				addToSumMerge(setx.None),
			},
			wantResult:      sumResult(setx.All),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
		{
			name: "NONE + ALL",
			conditions: []sumConditions{
				addToSumMerge(setx.None),
				addToSumMerge(setx.All),
			},
			wantResult:      sumResult(setx.All),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
	}
	return tests
}

func sumResult(operator setx.Operator, values ...string) setx.Sum[string] {
	return setx.Sum[string]{
		Result: setx.Result[string]{
			Operator: operator,
			Values:   values,
		},
	}
}

func addToSumMerge(operator setx.Operator, values ...string) sumConditions {
	return sumConditions{operator: operator, values: values}
}
