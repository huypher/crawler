package components

type Frontier interface {
	Push(item *Item)
	Pop() *Item
	Len() int
}
