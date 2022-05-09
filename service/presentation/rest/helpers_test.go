package rest

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
)

func Test_paginationParams(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.PaginationInput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := paginationParams(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("paginationParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("paginationParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterParams(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.FilterInput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filterParams(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
