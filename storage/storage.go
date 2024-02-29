package storage

//go:generate mockgen -destination=../mock/storage_mock.go -package=mock github.com/timdin/wdipms/storage Storage
type Storage interface {
}
