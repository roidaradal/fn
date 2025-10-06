package ds

type identifiable interface {
	hasID
	hasCode
}

type hasID interface {
	GetID() uint
}

type hasCode interface {
	GetCode() string
}

type (
	IDCodeMap             map[uint]string // ID => Code mapping
	CodeIDMap             map[string]uint // Code => ID mapping
	LookupID[T hasID]     map[uint]T      // ID => Item lookup
	LookupCode[T hasCode] map[string]T    // Code => Item lookup
)

// Create a new lookup map from list of items, using the entry function
func NewLookup[T any, K comparable, V any](items []T, entry func(T) (K, V)) map[K]V {
	lookup := make(map[K]V, len(items))
	for _, item := range items {
		k, v := entry(item)
		lookup[k] = v
	}
	return lookup
}

// Create new IDCodeMap: map[uint]string from list of items
func NewIDCodeMap[T identifiable](items []T) IDCodeMap {
	lookup := make(IDCodeMap)
	for _, item := range items {
		lookup[item.GetID()] = item.GetCode()
	}
	return lookup
}

// Create new CodeIDMap: map[string]uint from list of items
func NewCodeIDMap[T identifiable](items []T) CodeIDMap {
	lookup := make(CodeIDMap)
	for _, item := range items {
		lookup[item.GetCode()] = item.GetID()
	}
	return lookup
}

// Create new LookupID: map[uint]T from list of items
func NewLookupID[T hasID](items []T) LookupID[T] {
	lookup := make(LookupID[T])
	for _, item := range items {
		lookup[item.GetID()] = item
	}
	return lookup
}

// Create new LookupCode: map[string]T from list of items
func NewLookupCode[T hasCode](items []T) LookupCode[T] {
	lookup := make(LookupCode[T])
	for _, item := range items {
		lookup[item.GetCode()] = item
	}
	return lookup
}

// Create new LookupID: map[uint]T from given LookupCode
func NewLookupIDFromCode[T identifiable](codeLookup LookupCode[T], validIDs *Set[uint]) LookupID[T] {
	idLookup := make(LookupID[T])
	for _, item := range codeLookup {
		id := item.GetID()
		if validIDs != nil && !validIDs.Contains(id) {
			continue
		}
		idLookup[id] = item
	}
	return idLookup
}

// Create new LookupCode: map[string]T from given LookupID
func NewLookupCodeFromID[T identifiable](idLookup LookupID[T], validCodes *Set[string]) LookupCode[T] {
	codeLookup := make(LookupCode[T])
	for _, item := range idLookup {
		code := item.GetCode()
		if validCodes != nil && !validCodes.Contains(code) {
			continue
		}
		codeLookup[code] = item
	}
	return codeLookup
}
