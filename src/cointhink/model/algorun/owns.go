package algorun

func Owns(algologId string, accountId string) bool {
	log, err := Find(algologId)
	if err != nil {
		return false
	} else {
		return log.AccountId == accountId
	}
	return false
}
