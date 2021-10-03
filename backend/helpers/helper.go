package helper

import "time"

func ParseTimestamp(ts string) (time.Time, error) {

	parsed_timestamp, err := time.Parse(time.RFC3339, ts)

	if err != nil {
		return time.Time{}, err
	}

	return parsed_timestamp, err
}
