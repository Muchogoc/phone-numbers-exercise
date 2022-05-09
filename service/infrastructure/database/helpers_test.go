package database

import (
	"reflect"
	"testing"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
)

func Test_phoneCountry(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want domain.Country
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := phoneCountry(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("phoneCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}
