package usecases

import (
	"context"
	"reflect"
	"testing"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
	"github.com/Muchogoc/phone-numbers-exercise/service/infrastructure"
)

func TestUsecasesImpl_ListCustomers(t *testing.T) {
	type args struct {
		ctx        context.Context
		filters    *domain.FilterInput
		pagination domain.PaginationInput
	}
	tests := []struct {
		name    string
		args    args
		want    []*domain.Customer
		wantErr bool
	}{
		{
			name: "success: list customers",
			args: args{
				ctx: nil,
				filters: &domain.FilterInput{
					Country: nil,
					State:   nil,
				},
				pagination: domain.PaginationInput{
					Limit:  -1,
					Offset: -1,
				},
			},
			want:    []*domain.Customer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			infra := infrastructure.NewInfrastructure()
			u := NewUsecasesImpl(infra)
			got, err := u.ListCustomers(tt.args.ctx, tt.args.filters, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecasesImpl.ListCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecasesImpl.ListCustomers() = %v, want %v", got, tt.want)
			}
		})
	}
}
