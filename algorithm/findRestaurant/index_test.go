package findRestaurant

import (
	"fmt"
	"reflect"
	"testing"
)

type RestaurantTest struct {
	List1 []string
	List2 []string
	Want []string
}

func TestFindRestaurant(t *testing.T) {
	tests := []RestaurantTest{
		{[]string{"Shogun","Tapioca Express","Burger King","KFC"},
			[]string{"KFC","Shogun","Burger King"}, []string{"Shogun"}},
		{[]string{"Shogun","Tapioca Express","Burger King","KFC"},
			[]string{"KFC","Burger King","Tapioca Express","Shogun"},
			[]string{"KFC","Burger King","Tapioca Express","Shogun"}},
	}

	for _, tt := range tests {
		ret := findRestaurant(tt.List1, tt.List2)
		fmt.Println(ret)
		if !reflect.DeepEqual(ret, tt.Want){
			t.Errorf("error!")
		}
	}
}