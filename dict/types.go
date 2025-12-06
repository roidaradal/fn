package dict

// Key-Value pair
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

type (
	Object                = map[string]any
	UintMap               = map[string]uint
	IntMap                = map[string]int
	BoolMap               = map[string]bool
	StringMap             = map[string]string
	StringListMap         = map[string][]string
	StringCounter         = map[string]int
	IntCounter            = map[int]int
	Counter[T comparable] = map[T]int
)

// Create new counter, with each item initialized to count=0
func NewCounter[T comparable](items []T) Counter[T] {
	count := make(Counter[T], len(items))
	for _, item := range items {
		count[item] = 0
	}
	return count
}

// Create counter, with keys produced from keyFn
func CounterFunc[T any, K comparable](items []T, key func(T) K) Counter[K] {
	count := make(Counter[K], len(items))
	for _, item := range items {
		count[key(item)] += 1
	}
	return count
}

// Update the counter with the incoming items
func UpdateCounter[T comparable](counter Counter[T], items []T) {
	for _, item := range items {
		counter[item] += 1
	}
}

// Create new boolean map, with each item initialized to flag boolean
func Flags[T comparable](items []T, flag bool) map[T]bool {
	flags := make(map[T]bool)
	for _, item := range items {
		flags[item] = flag
	}
	return flags
}

// Return Entry as Key, Value
func (e Entry[K, V]) Tuple() (K, V) {
	return e.Key, e.Value
}
