package util

const (
	createdAtDesc = "created_at DESC"
	createdAtAsc  = "created_at ASC"
	updatedAtDesc = "updated_at DESC"
	updatedAtAsc  = "updated_at ASC"
)

func IsSupportedOrder(order string) bool {
	switch order {
	case createdAtAsc, createdAtDesc, updatedAtAsc, updatedAtDesc:
		return true
	}
	return false
}
