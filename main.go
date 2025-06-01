package main

import (
	"fmt"
	"math/big"
)

const maxi = 1000

func main() {
	f := make([]*big.Int, maxi)
	f[0] = big.NewInt(1)
	f[1] = big.NewInt(3)

	for i := 2; i < maxi; i++ {
		f[i] = new(big.Int).Add(
			new(big.Int).Mul(big.NewInt(5), f[i-1]),
			f[i-2],
		)
	}

	var arr []*big.Int
	for _, val := range f {
		if new(big.Int).Mod(val, big.NewInt(2)).Int64() == 1 {
			arr = append(arr, val)
		}

		if len(arr) > 40 {
			break
		}
	}

	if len(arr) >= 40 {
		fmt.Println("A[40]:", arr[39])
	} else {
		fmt.Println("мало нечётных чисел")
	}
}
