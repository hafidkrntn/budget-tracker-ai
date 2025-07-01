package pagination

type Resolved struct {
	Page   int
	Limit  int
	Offset int
}

type Pagination struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
}

type PaginatedResponse[T any] struct {
	Data       []T       `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// Resolve menghitung nilai page, limit, dan offset yang valid
func Resolve(page, limit int) Resolved {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	return Resolved{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}

func NewPaginatedResponse[T any](data []T, total, page, limit int) PaginatedResponse[T] {
	totalPages := total / limit
	if total%limit != 0 {
		totalPages++
	}

	return PaginatedResponse[T]{
		Data: data,
		Pagination: Pagination{
			Total:     total,
			Page:      page,
			Limit:     limit,
			TotalPage: totalPages,
		},
	}
}
