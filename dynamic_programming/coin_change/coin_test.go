package coin_change

import "testing"

func TestCoinChange(t *testing.T) {
	for _, tt := range []struct {
		givenCoins  []int
		givenAmount int
		expectedNum int
	}{
		{givenCoins: []int{1, 2, 5}, givenAmount: 11, expectedNum: 3},
	} {
		if res := coinChange(tt.givenCoins, tt.givenAmount); res != tt.expectedNum {
			t.Fatalf("expected %d number of coins for coins %v and amount %d. Got %d coins instead.", tt.expectedNum, tt.givenCoins, tt.givenAmount, res)
		}
	}
}
