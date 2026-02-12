package fail

import "errors"

var (
	InactiveItem    = errors.New("public: Inactive item")
	InvalidCode     = errors.New("public: Invalid code")
	InvalidDate     = errors.New("public: Invalid date")
	InvalidField    = errors.New("public: Invalid field")
	InvalidOption   = errors.New("public: Invalid option")
	InvalidType     = errors.New("public: Invalid type")
	MissingConfig   = errors.New("public: Missing config")
	MissingField    = errors.New("public: Missing field")
	MissingParams   = errors.New("public: Missing required parameters")
	NotFoundAccount = errors.New("public: Account not found")
	NotFoundItem    = errors.New("public: Item not found")
	NotFoundPath    = errors.New("public: Path not found")
	WrongType       = errors.New("public: Wrong type")
)
