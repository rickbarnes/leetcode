package maxprofit

import (
	"fmt"
	"testing"
)

func TestMaxProfitShouldReturnHighestProfit(t *testing.T) {
	profit := maxProfit([]int{7, 1, 5, 3, 6, 4})
	ans := 5
	if profit != ans {
		t.Errorf("Received: %d, Expected: %d", profit, ans)
	}
	fmt.Printf("Received: %d, Expected: %d \n", profit, ans)
}

func TestMaxProfitShouldReturnHighestProfitTwo(t *testing.T) {
	profit := maxProfit([]int{7, 1, 5, 3, 6, 4, 11})
	ans := 10
	if profit != ans {
		t.Errorf("Received: %d, Expected: %d", profit, ans)
	}
	fmt.Printf("Received: %d, Expected: %d \n", profit, ans)
}

func TestMaxProfitShouldReturn0IfNoProfitPossible(t *testing.T) {
	profit := maxProfit([]int{7, 6, 4, 3, 1})
	ans := 0
	if profit != ans {
		t.Errorf("Received: %d, Expected: %d", profit, ans)
	}
}
