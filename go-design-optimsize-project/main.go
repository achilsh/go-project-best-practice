package main

import (
	"fmt"
	"go-design-optimsize-project/service"
	"go-design-optimsize-project/service/chain_responsibility"
	"go-design-optimsize-project/service/easy_make_mistake"
)

func main() {
	fmt.Println("call go-design-optimize-project demo")
	fmt.Println("call type 1.")
	service.GetProduction()
	fmt.Println("call type 2.")
	service.GetProduction2(200)
	fmt.Println("call type 3.")
	service.GetProduction3(500, 2)

	fmt.Println("call chain responsibility demo: ")
	chain_responsibility.CallChainResponsibility()

	easy_make_mistake.CallCheckNilInterfaceVar()

	easy_make_mistake.UnmarshalJsonDemo()
	easy_make_mistake.WaitGroupCallWithMistakeAdd()
	easy_make_mistake.CoroutineLeakDemo()

}
