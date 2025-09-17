package conv

import (
	"fmt"

	"github.com/roidaradal/fn"
)

func IntToString(x int) string {
	return fmt.Sprintf("%d", x)
}

func FloatToString[T ~float32 | ~float64](x T) string {
	return fmt.Sprintf("%f", x)
}

func BooleanToString(flag bool) string {
	return fn.Ternary(flag, "1", "0")
}

func StringToBoolean(flag string) bool {
	return flag != "0"
}
