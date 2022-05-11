package rest

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
)

func Test_paginationParams(t *testing.T) {
	r1 := httptest.NewRequest(http.MethodGet, "/", nil)
	q1 := r1.URL.Query()
	q1.Add("offset", "0")
	q1.Add("limit", "5")
	r1.URL.RawQuery = q1.Encode()

	r2 := httptest.NewRequest(http.MethodGet, "/", nil)

	r3 := httptest.NewRequest(http.MethodGet, "/", nil)
	q3 := r3.URL.Query()
	q3.Add("offset", "-10")
	q3.Add("limit", "10")
	r3.URL.RawQuery = q3.Encode()

	r4 := httptest.NewRequest(http.MethodGet, "/", nil)
	q4 := r4.URL.Query()
	q4.Add("offset", "0")
	q4.Add("limit", "invalid")
	r4.URL.RawQuery = q4.Encode()

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.PaginationInput
		wantErr bool
	}{
		{
			name: "happy case: valid offset and limit",
			args: args{
				r: r1,
			},
			want: &domain.PaginationInput{
				Limit:  5,
				Offset: 0,
			},
			wantErr: false,
		},
		{
			name: "happy case: default offset and limit",
			args: args{
				r: r2,
			},
			want: &domain.PaginationInput{
				Limit:  -1,
				Offset: -1,
			},
			wantErr: false,
		},
		{
			name: "sad case: invalid offset",
			args: args{
				r: r3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "sad case: invalid limit",
			args: args{
				r: r4,
			},
			want:    nil,
			wantErr: true,
		},
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

	r1 := httptest.NewRequest(http.MethodGet, "/", nil)
	q1 := r1.URL.Query()
	q1.Add("country", "Cameroon")
	q1.Add("state", "true")
	r1.URL.RawQuery = q1.Encode()

	r2 := httptest.NewRequest(http.MethodGet, "/", nil)

	r3 := httptest.NewRequest(http.MethodGet, "/", nil)
	q3 := r3.URL.Query()
	q3.Add("country", "kenya")
	q3.Add("state", "true")
	r3.URL.RawQuery = q3.Encode()

	isValid := true

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.FilterInput
		wantErr bool
	}{
		{
			name: "happy case: filters not provided",
			args: args{
				r: r2,
			},
			want: &domain.FilterInput{
				Country: nil,
				State:   nil,
			},
			wantErr: false,
		},
		{
			name: "happy case: valid filters",
			args: args{
				r: r1,
			},
			want: &domain.FilterInput{
				Country: &domain.CountryCameroon,
				State:   &isValid,
			},
			wantErr: false,
		},
		{
			name: "sad case: invalid country",
			args: args{
				r: r3,
			},
			want:    nil,
			wantErr: true,
		},
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
