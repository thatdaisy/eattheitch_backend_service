package utils

func Paginate[T any](items []T, page, pageSize int) ([]T, int, int) {
	total := len(items)
	totalPages := (total + pageSize - 1) / pageSize

	if page < 1 {
		page = 1
	}
	if page > totalPages {
		return []T{}, total, totalPages
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > total {
		end = total
	}

	return items[start:end], total, totalPages
}
