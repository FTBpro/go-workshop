package something

import (
	"fmt"
)

type Actioner interface{
	Do(int, int) int
}

func Do(x, y int, action Actioner) {
	result := action.Do(x, y)
	fmt.Printf("x = %d, y = %d, action.Do(x, y) = %d\n", x, y, result)
}