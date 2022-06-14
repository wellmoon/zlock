package main

import (
	"fmt"

	"github.com/wellmoon/go/utils"
	"github.com/wellmoon/zlock/zlock"
)

func main() {
	zlockMap := zlock.New()

	go func() {

		zlockMap.Lock(1)
		fmt.Println("key 1 start lock go routinue 1")
		utils.Sleeps(5)
		zlockMap.Unlock(1)
		fmt.Println("key 1 unlock go routinue 1")
	}()
	go func() {

		zlockMap.Lock(1)
		fmt.Println("key 1 start lock go routinue 2")
		utils.Sleeps(5)
		zlockMap.Unlock(1)
		fmt.Println("key 1 unlock go routinue 2")
	}()
	// go func() {
	// 	fmt.Println("key 2 start go routinue 1")
	// 	zlockMap.Lock(2)
	// 	utils.Sleeps(5)
	// 	zlockMap.Unlock(2)
	// 	fmt.Println("key 2 unlock go routinue 1")
	// }()

	utils.Sleeps(20)
}
