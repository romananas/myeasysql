package myeasysql

import (
	"reflect"
	"testing"
)

func Test_SortKeys(t *testing.T) {
	tests := []struct {
		tags  []string
		names []string
		keys  []string
		want  []int
	}{
		{
			tags:  []string{"id", "username", "birth", "password"},
			names: []string{"id", "username", "birth", "password"},
			keys:  []string{"username", "password", "birth"},
			want:  []int{1, 3, 2},
		},
		{
			tags:  []string{},
			names: []string{"id", "username", "birth", "password"},
			keys:  []string{"username", "password", "birth"},
			want:  []int{1, 3, 2},
		},
		{
			tags:  []string{"id", "username", "birth", "password"},
			names: []string{"ID", "USERNAME", "BIRTH", "PASSWORD"},
			keys:  []string{"username", "password", "birth"},
			want:  []int{1, 3, 2},
		},
		{
			tags:  []string{"id", "username", "birth", "password"},
			names: []string{"id", "username", "birth", "password"},
			keys:  []string{"email", "phone"},
			want:  []int{},
		},
		{
			tags:  []string{"id", "username", "birth", "password"},
			names: []string{"id", "username", "birth", "password"},
			keys:  []string{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := sortKeys(tt.tags, tt.names, tt.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_SortKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
