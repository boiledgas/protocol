package utils

import (
	"container/list"
)

type Equality interface {
	Equal(interface{}) bool
}

type EqualityFunc func(v1 interface{}, v2 interface{}) bool

func Find(l *list.List, el Equality) {

}

func FindFunc(l *list.List, s interface{}, f EqualityFunc) {
	for e := l.Front(); e != nil; e = e.Next() {
		if f(s, e.Value) {
			break
		}
	}
}
