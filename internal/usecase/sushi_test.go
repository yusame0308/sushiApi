package usecase

import (
	"reflect"
	"sushiApi/internal/db/repository"
	"sushiApi/internal/http/gen"
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

func TestSushi_FindSushiById(t *testing.T) {
	type fields struct {
		sd repository.SushiDataInterface
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *gen.Sushi
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				sd: repository.NewMockSushiData(),
			},
			args: args{
				id: 1,
			},
			want: &gen.Sushi{
				NewSushi: gen.NewSushi{
					Name:  "mock",
					Price: 99,
					Sozai: []string{"納豆", "のり", "しゃり"},
				},
				Id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Sushi{
				sd: tt.fields.sd,
			}
			got, err := p.FindSushiById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindSushiById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSushiById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSushi_AddSushi(t *testing.T) {
	type fields struct {
		sd repository.SushiDataInterface
	}
	type args struct {
		param *gen.NewSushi
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *gen.Sushi
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				sd: repository.NewMockSushiData(),
			},
			args: args{
				param: &gen.NewSushi{
					Name:  "mock",
					Price: 99,
					Sozai: []string{"納豆", "のり", "しゃり"},
				},
			},
			want: &gen.Sushi{
				NewSushi: gen.NewSushi{
					Name:  "mock",
					Price: 99,
					Sozai: []string{"納豆", "のり", "しゃり"},
				},
				Id: 99,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Sushi{
				sd: tt.fields.sd,
			}
			got, err := p.AddSushi(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddSushi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSushi() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSushi_FindSushis(t *testing.T) {
	type fields struct {
		sd repository.SushiDataInterface
	}
	type args struct {
		params gen.FindSushisParams
	}
	isAsc := false
	limit := int32(99)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]gen.Sushi
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				sd: repository.NewMockSushiData(),
			},
			args: args{
				params: gen.FindSushisParams{
					Asc:   &isAsc,
					Limit: &limit,
				},
			},
			want: &[]gen.Sushi{
				{
					NewSushi: gen.NewSushi{
						Name:  "mock",
						Price: 99,
						Sozai: []string{"納豆", "のり", "しゃり"},
					},
					Id: 99,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Sushi{
				sd: tt.fields.sd,
			}
			got, err := p.FindSushis(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindSushis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSushis() got = %v, want %v", got, tt.want)
			}
		})
	}
}
