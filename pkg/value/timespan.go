package value

import (
	"fmt"
	"strings"
	"time"
)

const tick = 100 * time.Nanosecond

type Timespan struct {
	Value time.Duration

	Valid bool
}

func (Timespan) isKustoVal() {}

func (t Timespan) String() string {
	if !t.Valid {
		return ""
	}
	return t.Value.String()
}

func (t Timespan) Marshal() string {
	const (
		day = 24 * time.Hour
	)

	if !t.Valid {
		return "00:00:00"
	}

	// val is used to track the duration value as we move our parts of our time into our string format.
	// For example, after we write to our string the number of days that value had, we remove those days
	// from the duration. We continue doing this until val only holds values < 10 millionth of a second (tick)
	// as that is the lowest precision in our string representation.
	val := t.Value

	sb := strings.Builder{}

	// Add a - sign if we have a negative value. Convert our value to positive for easier processing.
	if t.Value < 0 {
		sb.WriteString("-")
		val = val * -1
	}

	// Only include the day if the duration is 1+ days.
	days := val / day
	val = val - (days * day)
	if days > 0 {
		sb.WriteString(fmt.Sprintf("%d.", int(days)))
	}

	// Add our hours:minutes:seconds section.
	hours := val / time.Hour
	val = val - (hours * time.Hour)
	minutes := val / time.Minute
	val = val - (minutes * time.Minute)
	seconds := val / time.Second
	val = val - (seconds * time.Second)
	sb.WriteString(fmt.Sprintf("%02d:%02d:%02d", int(hours), int(minutes), int(seconds)))

	// Add our sub-second string representation that is proceeded with a ".".
	milliseconds := val / time.Millisecond
	val = val - (milliseconds * time.Millisecond)
	ticks := val / tick
	if milliseconds > 0 || ticks > 0 {
		sb.WriteString(fmt.Sprintf(".%03d%d", milliseconds, ticks))
	}

	// Remove any trailing 0's.
	str := strings.TrimRight(sb.String(), "0")
	if strings.HasSuffix(str, ":") {
		str = str + "00"
	}

	return str
}
