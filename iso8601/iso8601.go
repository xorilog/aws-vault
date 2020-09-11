// The iso8601 package provides functionality for dealing with ISO-8601 times
// in a strict format compatible with all the AWS SDKs.
package iso8601

import (
	"time"
)

const format = "2006-01-02T15:04:05Z"
const jsonFormat = `"` + format + `"`

// Format outputs an AWS-compatible ISO-8601 string from the given time.
func Format(t time.Time) string {
	return New(t).String()
}

// Parse parses an AWS-compatible ISO-8601 string
func Parse(s string) (it Time, err error) {
	t, err := time.Parse(format, s)
	if err != nil {
		return
	}
	return New(t), nil
}

// Time represents an instant in time
type Time time.Time

// New constructs a new Time instance
func New(t time.Time) Time {
	return Time(t.UTC())
}

func (it Time) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(it).Format(jsonFormat)), nil
}

func (it *Time) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(jsonFormat, string(data))
	if err == nil {
		*it = New(t)
	}

	return err
}

// String outputs the AWS-compatible ISO-8601 representation of Time.
func (it Time) String() string {
	return time.Time(it).Format(format)
}
