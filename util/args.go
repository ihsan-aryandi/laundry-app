package util

import "os"

type args struct {
	arguments []string
}

func NewArgsFinder() *args {
	return &args{arguments: os.Args}
}

func (a *args) Args() []string {
	return a.arguments
}

func (a *args) Has(search string) bool {
	for _, arg := range a.arguments {
		if arg == search {
			return true
		}
	}
	return false
}