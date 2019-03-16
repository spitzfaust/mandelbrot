package main

import "testing"

func Test_normalizeToViewRectangle(t *testing.T) {
	type args struct {
		pX   int
		pY   int
		w    int
		h    int
		minX float64
		minY float64
		maxX float64
		maxY float64
	}
	tests := []struct {
		name  string
		args  args
		wantX float64
		wantY float64
	}{
		{
			name: "1024, 1024 is mapped to max values",
			args: args{
				pX:   0,
				pY:   0,
				w:    1024,
				h:    1024,
				minX: -2,
				minY: -1,
				maxX: 1,
				maxY: 1,
			},
			wantX: -2,
			wantY: -1,
		},
		{
			name: "0, 0 is mapped to min values",
			args: args{
				pX:   1024,
				pY:   1024,
				w:    1024,
				h:    1024,
				minX: -2,
				minY: -1,
				maxX: 1,
				maxY: 1,
			},
			wantX: 1,
			wantY: 1,
		},{
			name: "0, 0 is mapped to min values",
			args: args{
				pX:   1,
				pY:   1,
				w:    1024,
				h:    1024,
				minX: -2,
				minY: -1,
				maxX: 1,
				maxY: 1,
			},
			wantX: -1.997070313,
			wantY: -0.998046875,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := normalizeToViewRectangle(tt.args.pX, tt.args.pY, tt.args.w, tt.args.h, tt.args.minX, tt.args.minY, tt.args.maxX, tt.args.maxY)
			if !inDelta(got, tt.wantX, 0.000000001) {
				t.Errorf("normalizeToViewRectangle() got = %v, wantX %v", got, tt.wantX)
			}
			if !inDelta(got1, tt.wantY, 0.000000001) {
				t.Errorf("normalizeToViewRectangle() got1 = %v, wantY %v", got1, tt.wantY)
			}
		})
	}
}

func inDelta(a, b, delta float64) bool {
	dt := a - b

	return dt >= -delta && dt <= delta
}
