package game

type (
	InnerStyleId uint32
)

type InnerStyle struct {
	ID        int
	Name      string
	Stove     int
	Table     int
	DishTable int
}

var InnerStyles = map[InnerStyleId]*InnerStyle{
	1330004: {
		ID:        1330004,
		Name:      "普拉內飾1",
		Stove:     3,
		Table:     3,
		DishTable: 2,
	},
	1330006: {
		ID:        1330006,
		Name:      "普拉內飾2",
		Stove:     3,
		Table:     4,
		DishTable: 2,
	},
	1330009: {
		ID:        1330009,
		Name:      "超拉內飾2",
		Stove:     3,
		Table:     4,
		DishTable: 2,
	},
	1330010: {
		ID:        1330010,
		Name:      "普拉內飾3",
		Stove:     3,
		Table:     4,
		DishTable: 3,
	},
	1330011: {
		ID:        1330011,
		Name:      "超拉內飾3",
		Stove:     3,
		Table:     4,
		DishTable: 3,
	},
	1330012: {
		ID:        1330012,
		Name:      "普拉內飾4",
		Stove:     4,
		Table:     5,
		DishTable: 3,
	},
	1330013: {
		ID:        1330013,
		Name:      "超拉內飾4",
		Stove:     4,
		Table:     5,
		DishTable: 3,
	},
	1330014: {
		ID:        1330014,
		Name:      "石質內飾1",
		Stove:     4,
		Table:     4,
		DishTable: 3,
	},
	1330015: {
		ID:        1330015,
		Name:      "石質內飾2",
		Stove:     6,
		Table:     6,
		DishTable: 4,
	},
	1330017: {
		ID:        1330017,
		Name:      "普拉內飾5",
		Stove:     4,
		Table:     5,
		DishTable: 4,
	},
	1330018: {
		ID:        1330018,
		Name:      "超拉內飾5",
		Stove:     4,
		Table:     5,
		DishTable: 4,
	},
	1330019: {
		ID:        1330019,
		Name:      "石質內飾3",
		Stove:     4,
		Table:     5,
		DishTable: 4,
	},
	1330020: {
		ID:        1330020,
		Name:      "石質內飾4",
		Stove:     7,
		Table:     7,
		DishTable: 4,
	},
	1330023: {
		ID:        1330023,
		Name:      "騎士風格內飾(豪華版)",
		Stove:     7,
		Table:     8,
		DishTable: 6,
	},
	1330024: {
		ID:        1330024,
		Name:      "騎士風格內飾(專業版)",
		Stove:     5,
		Table:     6,
		DishTable: 6,
	},
}
