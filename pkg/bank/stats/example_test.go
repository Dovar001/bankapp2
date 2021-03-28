package stats

import (
	"fmt"

	"github.com/Dovar001/bank/pkg/bank/types"
)

func ExampleAvg() {

	cards := []types.Payment{

		{
		ID: 4444,
		Amount: 10_000,
		Category: "Машина",
		},
		{
		ID: 3333,
		Amount: 30_000,
		Category: "Ресторан",
		},
		
		{
		ID: 1111,
		Amount: 30_000,
		Category: "Магазин",
		},
		
	}

	 mid := Avg(cards)
	 fmt.Println(mid)

	//Output:23333

	
}