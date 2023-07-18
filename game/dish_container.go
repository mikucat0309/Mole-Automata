package game

type (
	DishContainerLoc uint32
	CookStatus       uint32
)

type DishContainer struct {
	CookID       CookID
	Loc          DishContainerLoc
	Dish         DishID
	Count        uint32
	state        uint32
	CookDuration uint32
}

const (
	COOK_STATUS_EMPTY     CookStatus = 0
	COOK_STATUS_PREPARE_1 CookStatus = 1
	COOK_STATUS_PREPARE_2 CookStatus = 2
	COOK_STATUS_COOKING   CookStatus = 3
	COOK_STATUS_COMPLETED CookStatus = 4
	COOK_STATUS_EXPIRED   CookStatus = 5
)

type Stove struct {
	Dish *DishContainer
}

type DishTable struct {
	Dish *DishContainer
}

func (d *DishContainer) IsEmpty() bool {
	return d.Count == 0
}

func (d *DishContainer) Info() *DishInfo {
	return DishInfos[d.Dish]
}

func (d *DishContainer) Status() CookStatus {
	if d.IsEmpty() {
		return COOK_STATUS_EMPTY
	}
	info := d.Info()
	if d.state == 1 {
		return COOK_STATUS_PREPARE_1
	}
	if d.state == 2 {
		return COOK_STATUS_PREPARE_2
	}
	if d.CookDuration < info.CompleteDuration {
		return COOK_STATUS_COOKING
	}
	if d.CookDuration < info.CompleteDuration+info.ExpireDuration {
		return COOK_STATUS_COMPLETED
	}
	return COOK_STATUS_EXPIRED
}
