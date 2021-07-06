package usecase

import (
	"reflect"
	"sushiApi/internal/db/repository"
	"testing"
)

func TestNewSushi(t *testing.T) {
	type args struct {
		db repository.SushiDataInterface
	}

	d := repository.NewSushiData(repository.NewBaseRepository())
	tests := []struct {
		name string
		args args
		want *Sushi
	}{
		{
			name: "正常",
			args: args{
				db: d,
			},
			want: &Sushi{sd: d},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSushi(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSushi() = %v, want %v", got, tt.want)
			}
		})
	}
}
