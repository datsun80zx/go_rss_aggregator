package rss

import (
	"database/sql"
	"time"
)

func parseTime(dateStr string) sql.NullTime {
	formats := []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC822,
		time.RFC822Z,
		time.RFC3339,
	}

	for _, format := range formats {
		if parsedTime, err := time.Parse(format, dateStr); err == nil {
			return sql.NullTime{
				Time:  parsedTime,
				Valid: true,
			}
		}
	}

	// If no format worked, return NULL
	return sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
}
