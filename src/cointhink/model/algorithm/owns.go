package algorithm

func Owns(algorithmId string, accountId string) bool {
	item, err := Find(algorithmId)
	if err != nil {
		return false
	} else {
		return item.AccountId == accountId
	}
	return false
}
