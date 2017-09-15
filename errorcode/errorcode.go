package errorcode

import "sync"

const (
	InvalidRequestData = "InvalidRequestData"
	InvalidSubscription = "InvalidSubscription"

	MemberNotFound     = "MemberNotFound"
	FriendNotFound		= "FriendNotFound"

	ValidationError = "ValidationError"
)

func initMap() {
	errMap = make(map[string]int)
	errMap[InvalidRequestData] = 400
	errMap[InvalidSubscription] = 400

	errMap[MemberNotFound] = 404
	errMap[FriendNotFound] = 404

	errMap[ValidationError] = 404
}

var errMap map[string]int
var once sync.Once

func GetErrorToHTPPStatusMap() map[string]int {
	once.Do(func() {
		initMap()
	})

	return errMap
}
