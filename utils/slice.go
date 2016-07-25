package utils

import (
	"sort"
)

type ByteSlice []byte
type Flags map[byte]bool

func (s ByteSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s ByteSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByteSlice) Len() int {
	return len(s)
}

func (flags Flags) GetFlags() []byte {
	res := make(ByteSlice, 0, len(flags))
	for flag, _ := range flags {
		res = append(res, flag)
	}

	sort.Sort(res)
	return res
}

func (flags Flags) GetFlag() byte {
	var val byte
	for flag, ok := range flags {
		if ok {
			val = val | flag
		}
	}

	return val
}

func (flags Flags) SetFlag(val byte) {
	fs := GetFlags8(val)
	newFlags := make(map[byte]bool)
	for _, flag := range fs {
		newFlags[flag] = true
	}

	for flag, _ := range flags {
		_, flags[flag] = newFlags[flag]
	}
}
