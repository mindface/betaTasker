package utils

var seedPath = "data"

func SetSeedPath(p string) {
	seedPath = p
}

func GetSeedPath() string {
	return seedPath
}