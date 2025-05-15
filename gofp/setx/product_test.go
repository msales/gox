package setx_test

import (
	"testing"

	"github.com/msales/gox/gofp/setx"
	"github.com/stretchr/testify/assert"
)

type productConditions struct {
	operator setx.Operator
	values   []string
}
type testCase struct {
	name            string
	conditions      []productConditions
	wantResult      setx.Product[string]
	wantMergeResult setx.MergeResult
	wantErr         bool
}

func TestResult_MergeProduct(t *testing.T) {
	tests := productTestCases()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := setx.NewProduct[string]()
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

func TestResult_MergeProductWithNone(t *testing.T) {
	tests := productTestCases()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// any test no error case if there is one more condition with all added to the conditions list, should result with
			// immutable result with ALL operator
			tt.conditions = append(tt.conditions, addToProductMerge(setx.None))
			tt.wantMergeResult = setx.Immutable
			tt.wantResult.Operator = setx.None
			tt.wantResult.Values = []string{}
			if tt.wantErr {
				return
			}

			result := setx.NewProduct[string]()
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

// productTestCases is providing cases for 2 tests
// Unfortunately in goland it prevents running separate case test, but nevertheless it's better than
// copy-paste of all the cases.
func productTestCases() []testCase {
	tests := []testCase{
		{
			name: "One IN condition",
			conditions: []productConditions{
				addToProductMerge(setx.In, "PL", "ES"),
			},
			wantResult:      productResult(setx.In, "PL", "ES"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "One NOT IN condition",
			conditions: []productConditions{
				addToProductMerge(setx.NotIn, "PL", "ES"),
			},
			wantResult:      productResult(setx.NotIn, "PL", "ES"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "Two IN conditions of disjoint sets",
			conditions: []productConditions{
				addToProductMerge(setx.In, "PL", "ES"),
				addToProductMerge(setx.In, "DE", "US"),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
		{
			name: "Two NOT IN conditions of disjoint sets",
			conditions: []productConditions{
				addToProductMerge(setx.NotIn, "PL", "ES"),
				addToProductMerge(setx.NotIn, "DE", "US"),
			},
			wantResult:      productResult(setx.NotIn, "PL", "ES", "DE", "US"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "Two IN conditions of intersecting sets",
			conditions: []productConditions{
				addToProductMerge(setx.In, "PL", "ES"),
				addToProductMerge(setx.In, "PL", "US"),
			},
			wantResult:      productResult(setx.In, "PL"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "Two NOT IN conditions of intersecting sets",
			conditions: []productConditions{
				addToProductMerge(setx.NotIn, "PL", "ES"),
				addToProductMerge(setx.NotIn, "PL", "US"),
			},
			wantResult:      productResult(setx.NotIn, "PL", "ES", "US"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "Three NOT IN conditions of intersecting sets, but no value in all conditions",
			conditions: []productConditions{
				addToProductMerge(setx.NotIn, "PL", "ES", "UK"),
				addToProductMerge(setx.NotIn, "PL", "US", "UK"),
				addToProductMerge(setx.NotIn, "IS", "US", "CZ"),
			},
			wantResult:      productResult(setx.NotIn, "PL", "ES", "UK", "US", "IS", "CZ"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "Three IN conditions of intersecting sets, but no value in all conditions",
			conditions: []productConditions{
				addToProductMerge(setx.In, "PL", "ES", "UK"),
				addToProductMerge(setx.In, "PL", "US", "UK"),
				addToProductMerge(setx.In, "IS", "US", "CZ"),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},

		{
			name: "Three NOT IN conditions of intersecting sets but at least one value in all conditions",
			conditions: []productConditions{
				addToProductMerge(setx.NotIn, "PL", "CA", "ES", "UK", "SK"),
				addToProductMerge(setx.NotIn, "PL", "SK", "US", "UK", "CA"),
				addToProductMerge(setx.NotIn, "CA", "IS", "US", "SK", "CZ"),
			},
			wantResult:      productResult(setx.NotIn, "PL", "CA", "ES", "UK", "SK", "US", "IS", "CZ"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "Three IN conditions of intersecting sets but at least one value in all conditions",
			conditions: []productConditions{
				addToProductMerge(setx.In, "PL", "CA", "ES", "UK", "SK"),
				addToProductMerge(setx.In, "PL", "SK", "US", "UK", "CA"),
				addToProductMerge(setx.In, "CA", "IS", "US", "SK", "CZ"),
			},
			wantResult:      productResult(setx.In, "CA", "SK"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "IN then NOT IN intersecting conditions",
			conditions: []productConditions{
				addToProductMerge(setx.In, "PL", "ES"),
				addToProductMerge(setx.NotIn, "PL", "US"),
			},
			wantResult:      productResult(setx.In, "ES"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "NOT IN then IN intersecting conditions",
			conditions: []productConditions{
				addToProductMerge(setx.NotIn, "PL", "US"),
				addToProductMerge(setx.In, "PL", "ES"),
			},
			wantResult:      productResult(setx.In, "ES"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "IN then NOT IN disjoint conditions",
			conditions: []productConditions{
				addToProductMerge(setx.In, "PL", "ES"),
				addToProductMerge(setx.NotIn, "CZ", "US"),
			},
			wantResult:      productResult(setx.In, "PL", "ES"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "NOT IN then IN disjoint conditions",
			conditions: []productConditions{
				addToProductMerge(setx.NotIn, "CZ", "US"),
				addToProductMerge(setx.In, "PL", "ES"),
			},
			wantResult:      productResult(setx.In, "PL", "ES"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "IN then NOT IN identical conditions",
			conditions: []productConditions{
				addToProductMerge(setx.In, "PL"),
				addToProductMerge(setx.NotIn, "PL"),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
		{
			name: "NOT IN then IN identical conditions",
			conditions: []productConditions{
				addToProductMerge(setx.NotIn, "PL"),
				addToProductMerge(setx.In, "PL"),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},

		{
			name: "IN then NOT IN conditions, NOT IN set contained in IN set",
			conditions: []productConditions{
				addToProductMerge(setx.In, "PL", "US"),
				addToProductMerge(setx.NotIn, "PL"),
			},
			wantResult:      productResult(setx.In, "US"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},
		{
			name: "NOT IN then IN conditions, NOT IN set contained in IN set",
			conditions: []productConditions{
				addToProductMerge(setx.NotIn, "PL"),
				addToProductMerge(setx.In, "PL", "US"),
			},
			wantResult:      productResult(setx.In, "US"),
			wantMergeResult: setx.Continue,
			wantErr:         false,
		},

		{
			name: "wrong ALL condition with values",
			conditions: []productConditions{
				addToProductMerge(setx.All, "PL"),
			},
			wantErr: true,
		},

		{
			name: "wrong NONE condition with values",
			conditions: []productConditions{
				addToProductMerge(setx.None, "PL"),
			},
			wantErr: true,
		},

		{
			name: "only NONE",
			conditions: []productConditions{
				addToProductMerge(setx.None),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
		{
			name: "twice NONE",
			conditions: []productConditions{
				addToProductMerge(setx.None),
				addToProductMerge(setx.None),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},

		{
			name: "IN + NONE",
			conditions: []productConditions{
				addToProductMerge(setx.In, "PL"),
				addToProductMerge(setx.None),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
		{
			name: "NONE + IN",
			conditions: []productConditions{
				addToProductMerge(setx.None),
				addToProductMerge(setx.In, "PL"),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},

		{
			name: "NOT IN + NONE",
			conditions: []productConditions{
				addToProductMerge(setx.NotIn, "PL"),
				addToProductMerge(setx.None),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
		{
			name: "NONE + IN",
			conditions: []productConditions{
				addToProductMerge(setx.None),
				addToProductMerge(setx.NotIn, "PL"),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},

		{
			name: "ALL + NONE",
			conditions: []productConditions{
				addToProductMerge(setx.All),
				addToProductMerge(setx.None),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
		{
			name: "NONE + ALL",
			conditions: []productConditions{
				addToProductMerge(setx.None),
				addToProductMerge(setx.All),
			},
			wantResult:      productResult(setx.None),
			wantMergeResult: setx.Immutable,
			wantErr:         false,
		},
	}
	return tests
}

func productResult(operator setx.Operator, values ...string) setx.Product[string] {
	return setx.Product[string]{
		Result: setx.Result[string]{
			Operator: operator,
			Values:   values,
		},
	}
}

func addToProductMerge(operator setx.Operator, values ...string) productConditions {
	return productConditions{operator: operator, values: values}
}
