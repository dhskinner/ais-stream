package encode

import (
	"log/slog"
	"math"
)

func VerificationError(source string, err string, expected interface{}, actual interface{}) bool {
	slog.Error("verification failed", "source", source, "error", err, "expected", expected, "actual", actual)
	return false
}

func VerifyFloat(expected float64, actual float64, precision int) bool {
	return math.Round(expected*float64(precision)) == math.Round(actual*float64(precision))
}
