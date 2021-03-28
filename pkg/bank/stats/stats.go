package stats

import "github.com/Dovar001/bank/pkg/bank/types"

//Avg расчитивает средную сумму платежа

func Avg(payments [] types.Payment)  types.Money{

	var mid types.Money
	var a types.Money
	for _, operation := range payments {

		mid+=operation.Amount
		a++

	}
	mid=mid/a
	return mid
}