package time_now

import "time"

var defaultFunctionTime = time.Now().UTC

func Now() time.Time {
	return defaultFunctionTime()
}

func ReplaceFunctionTime(functionTime func() time.Time) {
	defaultFunctionTime = functionTime
}
