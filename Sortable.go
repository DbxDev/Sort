package Sort

type Sortable interface {
	CompareTo(Sortable) int
}

func Less(a, b Sortable) bool {
	return a.CompareTo(b) < 0
}
