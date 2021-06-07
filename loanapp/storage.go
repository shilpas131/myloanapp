package loanapp

type StorageService interface {
	AddNewEntryToStorage(args ...interface{}) (interface{}, error)
	UpdateEntryInStorage(args ...interface{}) (interface{}, error)
	GetEntryFromStorage(args ...interface{}) (interface{}, error)
}
