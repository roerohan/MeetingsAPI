package utils

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

// PaginateMeetings returns 100 items in a page
func PaginateMeetings(meetings []bson.M, pageArr []string) ([]bson.M, bool) {
	itemsPerPage := 1

	page, _ := strconv.Atoi(pageArr[0])

	if page < 1 {
		return []bson.M{}, true
	}

	startIdx := (page - 1) * itemsPerPage
	endIdx := startIdx + itemsPerPage

	if startIdx >= len(meetings) {
		return []bson.M{}, false
	}

	if endIdx > len(meetings) {
		endIdx = len(meetings)
	}

	return meetings[startIdx:endIdx], false
}
