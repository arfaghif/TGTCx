package time_helper

import (
	"encoding/json"
	"time"

	"github.com/arfaghif/TGTCx/backend/dictionary"
)

func ParseTimestamp(ts string) (time.Time, error) {
	parsed_timestamp, err := time.Parse(time.RFC3339, ts)

	if err != nil {
		return time.Time{}, err
	}

	return parsed_timestamp, err
}

func BuildResponse(v dictionary.APIResponse) string {
	byte, _ := json.Marshal(v)
	return string(byte)
}
