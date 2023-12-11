package interfaces

type OrderingParam struct {
	OrderBy string `json:"orderby" binding:"optional" swag:"query"` // Ordering item by request (default: id)
	SortBy  string `json:"sortby" binding:"optional" swag:"query"`  // Sorting item by ascending or descending (default: asc)
}

type PaginationParam struct {
	Limit uint16 `json:"limit" binding:"optional" swag:"query"` // Limit is the maximum number of items per page (default: 10)
	Page  uint16 `json:"page" binding:"optional" swag:"query"`  // Page Number of items in page (default: 1)
}
