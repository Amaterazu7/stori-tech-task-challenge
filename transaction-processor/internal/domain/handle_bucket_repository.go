package domain

type HandleBucketRepository interface {
	FindFileByName(fileName string) (string, error)
	Find() error
}
