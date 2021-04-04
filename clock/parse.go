package clock

import (
	"errors"
	"time"
)

var (
	ParseTimeError = errors.New("parse time error")
)

func Parse(value, layout string) (time.Time, error) {
	if layout != "" {
		result, err := time.Parse(layout, value)
		if err != nil {
			return time.Time{}, ParseTimeError
		} else {
			return result, nil
		}
	}
	if result, err := ParseRFC3339(value); err == nil {
		return result, nil
	}
	if result, err := ParseRFC882(value); err == nil {
		return result, nil
	}
	if result, err := ParseRFC850(value); err == nil {
		return result, nil
	}
	if result, err := ParseRFC1123(value); err == nil {
		return result, nil
	}
	if result, err := ParseCommon(value); err == nil {
		return result, nil
	}
	if result, err := ParseStamps(value); err == nil {
		return result, nil
	}
	// Error
	return time.Time{}, ParseTimeError
}

func ParseTimestamp(timestamp int64, timezone string) (time.Time, error) {
	// Success
	return ParseMilliTimestamp(timestamp*1000, timezone)
}

func ParseMilliTimestamp(timestamp int64, timezone string) (time.Time, error) {
	// Success
	return ParseMicroTimestamp(timestamp*1000, timezone)
}

func ParseMicroTimestamp(timestamp int64, timezone string) (time.Time, error) {
	// Success
	return ParseNanoTimestamp(timestamp*1000, timezone)
}

func ParseNanoTimestamp(timestamp int64, timezone string) (time.Time, error) {
	// Success
	return InTimezone(time.Unix(0, timestamp), timezone)
}

func ParseRFC3339(value string) (time.Time, error) {
	for _, layout := range GroupRFC3339 {
		if result, err := time.Parse(layout, value); err == nil {
			return result, nil
		}
	}
	// Error
	return time.Time{}, ParseTimeError
}

func ParseRFC882(value string) (time.Time, error) {
	for _, layout := range GroupRFC822 {
		if result, err := time.Parse(layout, value); err == nil {
			return result, nil
		}
	}
	// Error
	return time.Time{}, ParseTimeError
}

func ParseRFC850(value string) (time.Time, error) {
	for _, layout := range GroupRFC850 {
		if result, err := time.Parse(layout, value); err == nil {
			return result, nil
		}
	}
	// Error
	return time.Time{}, ParseTimeError
}

func ParseRFC1123(value string) (time.Time, error) {
	for _, layout := range GroupRFC1123 {
		if result, err := time.Parse(layout, value); err == nil {
			return result, nil
		}
	}
	// Error
	return time.Time{}, ParseTimeError
}

func ParseStamps(value string) (time.Time, error) {
	for _, layout := range GroupStamps {
		if result, err := time.Parse(layout, value); err == nil {
			return result, nil
		}
	}
	// Error
	return time.Time{}, ParseTimeError
}

func ParseCommon(value string) (time.Time, error) {
	for _, layout := range GroupCommon {
		if result, err := time.Parse(layout, value); err == nil {
			return result, nil
		}
	}
	// Error
	return time.Time{}, ParseTimeError
}
