package promotions

import "testing"

func Test_bySKUDiscount_Discount(t *testing.T) {
	type args struct {
		p *product
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Given a product with 000003 sku, it will apply 15% discount",
			args: args{
				p: &product{
					Sku:      "000003",
					Name:     "test",
					Category: "others",
					Price:    Price{},
				},
			},
			want: 15,
		},
		{
			name: "Given a product with 000004 sku, it will apply no discount",
			args: args{
				p: &product{
					Sku:      "000004",
					Name:     "test",
					Category: "others",
					Price:    Price{},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bsd := &bySKUDiscount{}
			if got := bsd.Discount(tt.args.p); got != tt.want {
				t.Errorf("Discount() = %v, want %v", got, tt.want)
			}
		})
	}
}
