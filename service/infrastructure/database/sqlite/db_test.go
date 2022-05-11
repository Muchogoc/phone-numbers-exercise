package sqlite

import (
	"context"
	"os"
	"testing"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
)

func TestMain(m *testing.M) {
	initialValue := os.Getenv(RunningTestsEnvName)
	os.Setenv(RunningTestsEnvName, "true")

	exitVal := m.Run()

	os.Setenv(RunningTestsEnvName, initialValue)

	os.Exit(exitVal)
}

func TestDBImpl_GetCustomers(t *testing.T) {
	type args struct {
		ctx        context.Context
		filters    *domain.FilterInput
		pagination domain.PaginationInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success: list customers",
			args: args{
				ctx: nil,
				filters: &domain.FilterInput{
					Country: nil,
				},
				pagination: domain.PaginationInput{
					Limit:  -1,
					Offset: -1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewSqliteDBImpl()
			got, err := d.GetCustomers(tt.args.ctx, tt.args.filters, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.GetCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("DBImpl.GetCustomers() = %v", got)
			}
		})
	}
}
