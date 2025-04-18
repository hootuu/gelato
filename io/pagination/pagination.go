package pagination

import "fmt"

const DefaultPagingSize = 100

const MaxPagingSize = 2000

const FirstPage = 1

type Page struct {
	Size int64 `json:"size"` // Page Size, default 100
	Numb int64 `json:"numb"` // Page Numb, From One
}

func PageALL() *Page {
	return &Page{
		Size: MaxPagingSize,
		Numb: FirstPage,
	}
}

func PageNormal() *Page {
	return &Page{
		Size: DefaultPagingSize,
		Numb: FirstPage,
	}
}

type Paging struct {
	Size  int64 `json:"size"`  // Page Size, default 100
	Numb  int64 `json:"numb"`  // Page Numb, From One
	Total int64 `json:"total"` // Page Total, The Page Total
	Count int64 `json:"count"` // Item Count, The Item Count
}

func (paging *Paging) ToString() string {
	return fmt.Sprintf("size: %d|numb: %d|total: %d|count: %d",
		paging.Size, paging.Numb, paging.Total, paging.Count)
}

func (paging *Paging) WithCount(count int64) *Paging {
	if paging.Size == 0 {
		paging.Size = DefaultPagingSize
	}
	if paging.Numb < FirstPage {
		paging.Numb = FirstPage
	}

	if count == 0 {
		paging.Total = 0
		paging.Numb = FirstPage
		return paging
	}
	paging.Count = count
	paging.Total = paging.Count / paging.Size
	if paging.Count%paging.Size != 0 {
		paging.Total += 1
	}

	if paging.Numb > paging.Total {
		paging.Numb = paging.Total
	}

	if paging.Numb < 1 {
		paging.Numb = 1
	}
	return paging
}

func (paging *Paging) Skip() int64 {
	return (paging.Numb - 1) * paging.Size
}

func (paging *Paging) Limit() int64 {
	return paging.Size
}

func PagingOfPage(page *Page) *Paging {
	if page == nil {
		return PagingOf(DefaultPagingSize, FirstPage)
	}
	return PagingOf(page.Size, page.Numb)
}

func PagingOf(size int64, current int64) *Paging {
	if size <= 0 {
		size = DefaultPagingSize
	}
	if size > MaxPagingSize {
		size = MaxPagingSize
	}
	if current < FirstPage {
		current = FirstPage
	}
	return &Paging{
		Size:  size,
		Numb:  current,
		Total: 0,
		Count: 0,
	}
}

func PagingALL() *Paging {
	return PagingOf(MaxPagingSize, FirstPage)
}

type Pagination[T any] struct {
	Paging *Paging `bson:"paging" json:"paging"`
	Data   []*T    `bson:"data" json:"data"`
}

func NewPagination[T any](paging *Paging, data []*T) *Pagination[T] {
	return &Pagination[T]{
		Paging: paging,
		Data:   data,
	}
}

func (p *Pagination[T]) Iter(fn func(item *T, idx int)) {
	for idx, item := range p.Data {
		fn(item, idx)
	}
}
