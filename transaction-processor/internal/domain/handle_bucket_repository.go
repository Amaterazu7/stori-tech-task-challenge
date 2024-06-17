package domain

type HandleBucketRepository interface {
	FindFileByName(fileName string) error
}
