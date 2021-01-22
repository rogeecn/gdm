package utils

func IsOK(ret int64) bool {
	return ret == 1
}

func IsColorOK(ret int64) bool {
	return ret == 0
}

func IsFindStrOK(ret int64) bool {
	return ret != -1
}
