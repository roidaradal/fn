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
	IDCodeLookup          = map[uint]string // ID => Code lookup
	CodeIDLookup          = map[string]uint // Code => ID lookup
	LookupID[T hasID]     = map[uint]T      // ID => Item lookup
	LookupCode[T hasCode] = map[string]T    // Code => Item lookup
)

// Create new lookup from list of items, using entry function
func NewLookup[T any, K comparable, V any](items []T, entry func(T) (K, V)) map[K]V {
	lookup := make(map[K]V, len(items))
	for _, item := range items {
		k, v := entry(item)
		lookup[k] = v
	}
	return lookup
}

// Create new IDCodeLookup from list of items
func NewIDCodeLookup[T identifiable](items []T) IDCodeLookup {
	lookup := make(IDCodeLookup)
	for _, item := range items {
		lookup[item.GetID()] = item.GetCode()
	}
	return lookup
}

// Create new CodeIDLookup from list of items
func NewCodeIDLookup[T identifiable](items []T) CodeIDLookup {
	lookup := make(CodeIDLookup)
	for _, item := range items {
		lookup[item.GetCode()] = item.GetID()
	}
	return lookup
}

// Create new LookupID from list of items
func NewLookupID[T hasID](items []T) LookupID[T] {
	lookup := make(LookupID[T])
	for _, item := range items {
		lookup[item.GetID()] = item
	}
	return lookup
}

// Create new LookupCode from list of items
func NewLookupCode[T hasCode](items []T) LookupCode[T] {
	lookup := make(LookupCode[T])
	for _, item := range items {
		lookup[item.GetCode()] = item
	}
	return lookup
}

// Create new LookupID from given LookupCode,
// Use nil set to skip id validation
func LookupIDFromCode[T identifiable](codeLookup LookupCode[T], validIDs *Set[uint]) LookupID[T] {
	idLookup := make(LookupID[T])
	for _, item := range codeLookup {
		id := item.GetID()
		if validIDs != nil && validIDs.HasNo(id) {
			continue
		}
		idLookup[id] = item
	}
	return idLookup
}

// Create new LookupCode from given LookupID,
// Use nil set to skip code validation
func LookupCodeFromID[T identifiable](idLookup LookupID[T], validCodes *Set[string]) LookupCode[T] {
	codeLookup := make(LookupCode[T])
	for _, item := range idLookup {
		code := item.GetCode()
		if validCodes != nil && validCodes.HasNo(code) {
			continue
		}
		codeLookup[code] = item
	}
	return codeLookup
}
