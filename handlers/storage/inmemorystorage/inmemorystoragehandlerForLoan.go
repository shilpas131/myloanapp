package inmemorystorage

import (
	"errors"
	"myloanapp/loanapp"
	"time"

	cache "github.com/patrickmn/go-cache"
)

type InMemoryStorageLoanHandler struct {
	cache *cache.Cache
}

func (st *InMemoryStorageLoanHandler) AddNewEntryToStorage(args ...interface{}) (interface{}, error) {
	key := args[0].(int)
	value := args[1].(loanapp.Loan)
	st.cache.Set(string(key), &value, cache.DefaultExpiration)

	return nil, nil
}

func (st *InMemoryStorageLoanHandler) UpdateEntryInStorage(args ...interface{}) (interface{}, error) {
	key := args[0].(int)
	value := args[1].(loanapp.Loan)
	st.cache.Set(string(key), &value, cache.DefaultExpiration)

	return nil, nil
}

func (st *InMemoryStorageLoanHandler) GetEntryFromStorage(args ...interface{}) (interface{}, error) {
	key := args[0].(int)
	if val, found := st.cache.Get(string(key)); found {
		return val, nil
	}

	return nil, errors.New("no loan data found")
}

func GetNewInMemoryStorageLoanHandler() *InMemoryStorageLoanHandler {
	c := cache.New(5*time.Minute, 10*time.Minute)

	return &InMemoryStorageLoanHandler{
		cache: c,
	}
}
