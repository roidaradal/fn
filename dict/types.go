package dict

// Key-Value pair
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

type (
	Object        = map[string]any
	UintMap       = map[string]uint
	IntMap        = map[string]int
	BoolMap       = map[string]bool
	StringMap     = map[string]string
	StringListMap = map[string][]string
	StringCounter = map[string]int
	IntCounter    = map[int]int
)

// Create new counter, with each item initialized to count=0
func NewCounter[T comparable](items []T) map[T]int {
	count := make(map[T]int, len(items))
	for _, item := range items {
		count[item] = 0
	}
	return count
}

// Create new boolean map, with each item initialized to flag boolean
func Flags[T comparable](items []T, flag bool) map[T]bool {
	flags := make(map[T]bool)
	for _, item := range items {
		flags[item] = flag
	}
	return flags
}
