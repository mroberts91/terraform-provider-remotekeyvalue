package exampleserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type KeyValueStore struct {
	Pairs []KeyValuePair
	Lock  sync.Mutex
}

var _ ServerInterface = (*KeyValueStore)(nil)

func NewStore(pairs []KeyValuePair) *KeyValueStore {
	return &KeyValueStore{
		Pairs: pairs,
	}
}

func (s *KeyValueStore) GetKeyValuePair(w http.ResponseWriter, r *http.Request, key string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	var result KeyValuePair

	if key == "" {
		w.WriteHeader(http.StatusNotFound)
		err := &Error{
			Code:    http.StatusNotFound,
			Message: "Key must be supplied to lookup Key Value Pair",
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	for _, pair := range s.Pairs {
		lowerKey := strings.ToLower(pair.Key)
		if lowerKey == strings.ToLower(key) {
			result = pair
			break
		}
	}

	if result == (KeyValuePair{}) {
		w.WriteHeader(http.StatusNotFound)
		err := &Error{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Key Value Pair with key %s was not found...", key),
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetDummyData() [2]KeyValuePair {
	pairs := [...]KeyValuePair{
		CreatePair(1, "Foo", "Bar", false),
		CreatePair(2, "Biz", "Baz", true),
	}

	return pairs
}

func CreatePair(id int64, key string, value string, isSensitive bool) KeyValuePair {
	return KeyValuePair{
		Id:          id,
		Key:         key,
		Value:       value,
		IsSensitive: isSensitive,
	}
}
