package promotions

import "testing"

func Test_bootsCategoryDiscount_Discount(t *testing.T) {
	type args struct {
		p *product
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Given a product with {boots} category, It will apply 30% discount",
			args: args{
				p: &product{
					Sku:      "test",
					Name:     "boots-test",
					Category: "boots",
					Price:    Price{},
				},
			},
			want: 30,
		},
		{
			name: "Given a product with {others} category, It will apply 0% discount",
			args: args{
				p: &product{
					Sku:      "test",
					Name:     "others-test",
					Category: "others",
					Price:    Price{},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bcd := &bootsCategoryDiscount{}
			if got := bcd.Discount(tt.args.p); got != tt.want {
				t.Errorf("Discount() = %v, want %v", got, tt.want)
			}
		})
	}
}
