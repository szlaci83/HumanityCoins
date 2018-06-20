package test

import (
    "fmt"
    "github.com/szlaci83/HumanityCoins/web/controllers"
)

func Serve(app *controllers.Application) {
	result, err := app.Fabric.InvokeCreateEntry("1213", "hashTHIS_IS_A_TEST", "101", "101", "15625626")
	if err != nil {
		fmt.Println("ERROR!!!!!!!!!!!!!!" + result)
	}
	fmt.Println(result)
	result2, err := app.Fabric.InvokeAddTx("1213", "101", "ACCESS", "15625626")
	if err != nil {
		fmt.Println("ERROR!!!!!!!!!!!!!!" + result2)
	}
	fmt.Println(result2)
	result3, err := app.Fabric.InvokeGetByUUID("1213")
	if err != nil {
		fmt.Println("ERROR!!!!!!!!!!!!!!" + result2)
	}
	fmt.Println(result3)
}