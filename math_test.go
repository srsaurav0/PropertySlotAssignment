package main

import (
	"sort"
	"testing"
)

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]float64
		expected map[string]float64
	}{
		{
			name: "equal percentages sum to 100",
			input: map[string]float64{
				"11": 33.33,
				"12": 33.33,
				"24": 33.34,
			},
			expected: map[string]float64{
				"11": 33.33,
				"12": 33.33,
				"24": 33.34,
			},
		},
		{
			name: "remove zero percentages",
			input: map[string]float64{
				"11": 50.0,
				"12": 0.0,
				"24": 50.0,
			},
			expected: map[string]float64{
				"11": 50.0,
				"24": 50.0,
			},
		},
		{
			name: "normalize percentages to 100",
			input: map[string]float64{
				"11": 30.0,
				"12": 30.0,
				"24": 30.0,
			},
			expected: map[string]float64{
				"11": 33.33333333333333,
				"12": 33.33333333333333,
				"24": 33.33333333333333,
			},
		},
		{
			name: "single feed 100 percent",
			input: map[string]float64{
				"11": 100.0,
			},
			expected: map[string]float64{
				"11": 100.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validate_input(tt.input)

			// Compare map lengths
			if len(result) != len(tt.expected) {
				t.Errorf("validate_input() got length = %v, want length = %v", len(result), len(tt.expected))
				return
			}

			// Compare keys
			resultKeys := getSortedKeys(result)
			expectedKeys := getSortedKeys(tt.expected)
			if !compareStringSlices(resultKeys, expectedKeys) {
				t.Errorf("validate_input() got keys = %v, want keys = %v", resultKeys, expectedKeys)
				return
			}
		})
	}
}

func TestAssignFeed(t *testing.T) {
	tests := []struct {
		name         string
		input        map[string]float64
		limit        int
		priorityFeed string
		expected     map[string]int
	}{
		{
			name: "equal distribution with limit 3",
			input: map[string]float64{
				"11": 33.33,
				"12": 33.33,
				"24": 33.34,
			},
			limit:        3,
			priorityFeed: "12-11-24",
			expected: map[string]int{
				"11": 1,
				"12": 1,
				"24": 1,
			},
		},
		{
			name: "priority based assignment with limit 2",
			input: map[string]float64{
				"11": 30.0,
				"12": 30.0,
				"24": 40.0,
			},
			limit:        2,
			priorityFeed: "12-11-24",
			expected: map[string]int{
				"11": 0,
				"12": 1,
				"24": 1,
			},
		},
		{
			name: "single feed gets all",
			input: map[string]float64{
				"11": 100.0,
			},
			limit:        5,
			priorityFeed: "11",
			expected: map[string]int{
				"11": 5,
			},
		},
		{
			name: "handle equal percentages with priority",
			input: map[string]float64{
				"11": 50.0,
				"12": 50.0,
			},
			limit:        3,
			priorityFeed: "12-11",
			expected: map[string]int{
				"11": 1,
				"12": 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := assign_feed(tt.input, tt.limit, tt.priorityFeed)

			// Compare map lengths
			if len(result) != len(tt.expected) {
				t.Errorf("assign_feed() got length = %v, want length = %v", len(result), len(tt.expected))
				return
			}

			// Compare keys and values
			for k, expectedVal := range tt.expected {
				resultVal, exists := result[k]
				if !exists {
					t.Errorf("assign_feed() missing key = %v", k)
					continue
				}
				if resultVal != expectedVal {
					t.Errorf("assign_feed() for key %s got = %v, want = %v", k, resultVal, expectedVal)
				}
			}

			// Verify total matches limit
			total := 0
			for _, v := range result {
				total += v
			}
			if total != tt.limit {
				t.Errorf("assign_feed() total = %v, want limit = %v", total, tt.limit)
			}
		})
	}
}

// Helper function to get sorted keys from map
func getSortedKeys(m map[string]float64) []string {
	keys := make([]string, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Helper function to compare string slices
func compareStringSlices(a, b []string) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
