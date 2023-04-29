package main

import (
	MS "GCByGo/GCMarkSweep/achieve/notConsidered/PKG"
	"fmt"
)

func main() {
	MS.Init_data()
	fmt.Println("### init ###")
	MS.Print_data()

	MS.Mark_phase()
	fmt.Println("### mark phase ###")
	MS.Print_data()

	MS.Sweep_phase()
	fmt.Println("### sweep phase ###")
	MS.Print_data()
}
