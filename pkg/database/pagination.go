package database

// PaginationOption - Option type
type PaginationOption func(*Pagination)

// Size -
func Size(size uint32) PaginationOption {
	return func(c *Pagination) {
		c.size = size
	}
}

// Page -
func Page(page uint32) PaginationOption {
	return func(c *Pagination) {
		c.page = page
	}
}

// Pagination -
type Pagination struct {
	size uint32
	page uint32
}

// NewPagination -
func NewPagination(opts ...PaginationOption) Pagination {
	pagination := &Pagination{}

	// Custom options
	for _, opt := range opts {
		opt(pagination)
	}

	return *pagination
}

// Size -
func (p Pagination) Size() uint32 {
	return p.size
}

// Page -
func (p Pagination) Page() uint32 {
	return p.page
}
