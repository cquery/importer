package lib

import (
	"fmt"
)

type Updater interface {
	Set(name string) *UpdateSet
	Update(set ...*UpdateSet) error
}

type UpdateSet struct {
	Name   string
	Keys   []string
	Values []string
}

func (s *UpdateSet) AddString(key, value string) {
	s.Keys = append(s.Keys, key)
	s.Values = append(s.Values, value)
}

func (s UpdateSet) String() string {
	return fmt.Sprintf("%s:%s", s.Keys, s.Values)
}
