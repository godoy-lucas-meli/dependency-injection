package business

import "testing"

func TestBeerPacksQuantity(t *testing.T) {
	type args struct {
		attendees      uint
		minBeersPerBox uint
		temperature    float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "< 20 degrees",
			args: args{
				attendees:      2,
				minBeersPerBox: 6,
				temperature:    19.99,
			},
			want:    0.25,
			wantErr: false,
		},
		{
			name: "20 degrees",
			args: args{
				attendees:      6,
				minBeersPerBox: 6,
				temperature:    22,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "too hot, I'm thirsty",
			args: args{
				attendees:      6,
				minBeersPerBox: 6,
				temperature:    30,
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "no attendees",
			args: args{
				attendees:      0,
				minBeersPerBox: 6,
				temperature:    30,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "empty beer packs",
			args: args{
				attendees:      99,
				minBeersPerBox: 0,
				temperature:    30,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := beerPacksQuantity(tt.args.attendees, tt.args.minBeersPerBox, tt.args.temperature)
			if (err != nil) != tt.wantErr {
				t.Errorf("beerPacksQuantity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("beerPacksQuantity() got = %v, want %v", got, tt.want)
			}
		})
	}
}
