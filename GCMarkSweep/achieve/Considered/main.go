package main

import (
	MSC "GCByGo/GCMarkSweep/achieve/Considered/PKG"
	"fmt"
)

func main() {
	MSC.Init_data(7, 20)
	fmt.Println("### init ###")
	MSC.Print_data()

	MSC.Mark_phase()
	fmt.Println("### mark phase ###")
	MSC.Print_data()

	MSC.Sweep_phase()
	fmt.Println("### sweep phase ###")
	MSC.Print_data()
	//测试分配
	_ = MSC.NewObject("Data", []int{22}, 6)
}
