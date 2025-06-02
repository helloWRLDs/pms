package list

import (
	"testing"
)

type User struct {
	ID string
}

// func TestList_String(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		list     List[User]
// 		expected string
// 	}{
// 		{
// 			name: "empty list",
// 			list: List[User]{
// 				Pagination: Pagination{
// 					Page:       1,
// 					PerPage:    10,
// 					TotalItems: 0,
// 					TotalPages: 0,
// 				},
// 				Items: []User{},
// 			},
// 			expected: "page=1 per_page=10 created_from= created_to= order_by= ",
// 		},
// 		{
// 			name: "list with items",
// 			list: List[User]{
// 				Pagination: Pagination{
// 					Page:       2,
// 					PerPage:    20,
// 					TotalItems: 2,
// 					TotalPages: 1,
// 				},
// 				Items: []User{
// 					{ID: "1"},
// 					{ID: "2"},
// 				},
// 			},
// 			expected: "page=2 per_page=20 created_from= created_to= order_by= ",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.list.String(); got != tt.expected {
// 				t.Errorf("List.String() = %v, want %v", got, tt.expected)
// 			}
// 		})
// 	}
// }

func TestFilters_String(t *testing.T) {
	tests := []struct {
		name     string
		filters  Filters
		expected string
	}{
		{
			name: "empty filters",
			filters: Filters{
				Pagination: Pagination{
					Page:    1,
					PerPage: 10,
				},
				Date: Date{
					From: "",
					To:   "",
				},
				Order: Order{
					By:       "",
					Ascended: false,
				},
			},
			expected: "page=1 per_page=10 created_from= created_to= order_by= ",
		},
		{
			name: "filters with values",
			filters: Filters{
				Pagination: Pagination{
					Page:    2,
					PerPage: 20,
				},
				Date: Date{
					From: "2024-01-01",
					To:   "2024-12-31",
				},
				Order: Order{
					By:       "created_at",
					Ascended: true,
				},
				Fields: map[string]string{
					"status": "active",
				},
			},
			expected: "page=2 per_page=20 created_from=2024-01-01 created_to=2024-12-31 order_by=created_at ASC status=active ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.filters.String(); got != tt.expected {
				t.Errorf("Filters.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestOrder_String(t *testing.T) {
	tests := []struct {
		name     string
		order    Order
		expected string
	}{
		{
			name: "ascending order",
			order: Order{
				By:       "created_at",
				Ascended: true,
			},
			expected: "created_at ASC",
		},
		{
			name: "descending order",
			order: Order{
				By:       "created_at",
				Ascended: false,
			},
			expected: "created_at DESC",
		},
		{
			name: "empty order",
			order: Order{
				By:       "",
				Ascended: false,
			},
			expected: " DESC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.order.String(); got != tt.expected {
				t.Errorf("Order.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}
