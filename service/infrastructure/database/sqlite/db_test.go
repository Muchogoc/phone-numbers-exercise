package sqlite

import (
	"context"
	"reflect"
	"testing"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
	"gorm.io/gorm"
)

func TestDBImpl_GetCustomers(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx        context.Context
		filters    *domain.FilterInput
		pagination domain.PaginationInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Customer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewSqliteDBImpl()
			got, err := d.GetCustomers(tt.args.ctx, tt.args.filters, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.GetCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBImpl.GetCustomers() = %v, want %v", got, tt.want)
			}
		})
	}
}
