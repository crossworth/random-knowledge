package main

import (
	"fmt"
)

// sources:
// https://floating-point-gui.de
// http://fabiensanglard.net/floating_point_visually_explained/
// https://float.exposed/0x4048f5c3
func main() {
	accuracy()
}

func accuracy() {
	fmt.Println(0.1 + 0.2)                                                       // > 0.3
	fmt.Println(1 / 3.0)                                                         // > 0.3333333333333333
	fmt.Println(0.33 + 0.33 + 0.33)                                              // > 0.99
	fmt.Println(0.333 + 0.333 + 0.333)                                           // > 0.999
	fmt.Println(0.3333 + 0.3333 + 0.3333)                                        // > 0.9999
	fmt.Println(0.33333 + 0.33333 + 0.33333)                                     // > 0.99999
	fmt.Println(0.333333 + 0.333333 + 0.333333)                                  // > 0.999999
	fmt.Println(0.3333333 + 0.3333333 + 0.3333333)                               // > 0.9999999
	fmt.Println(0.33333333 + 0.33333333 + 0.33333333)                            // > 0.99999999
	fmt.Println(0.3333333333333333 + 0.3333333333333333 + 0.3333333333333333)    // > 0.9999999999999999
	fmt.Println(0.33333333333333333 + 0.33333333333333333 + 0.33333333333333333) // > 1
}
