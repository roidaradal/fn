package fn

import "strings"

type hasTime interface {
	GetTime() string
}

type hasID interface {
	GetID() uint
}

func SortByTime[T hasTime](item1, item2 T) int {
	return strings.Compare(item1.GetTime(), item2.GetTime())
}

func SortByTimeDesc[T hasTime](item1, item2 T) int {
	return -1 * SortByTime(item1, item2)
}

func SortByID[T hasID](item1, item2 T) int {
	return SortID(item1.GetID(), item2.GetID())
}

func SortByIDDesc[T hasID](item1, item2 T) int {
	return -1 * SortByID(item1, item2)
}

func SortID(item1, item2 uint) int {
	id1, id2 := int(item1), int(item2)
	return id1 - id2
}

func SortIDDesc(item1, item2 uint) int {
	return -1 * SortID(item1, item2)
}
