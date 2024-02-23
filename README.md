# ColorGradient

## Exemple code:
```Go
package main

import (
	"fmt"
	"github.com/edenbwt/ColorGradient"
)
	// using the ansi colors code to make a gradient containing 6 colors value
var Orange_Purple []int = []int{166, 167, 168, 169, 170, 171}

func main() {
		//using the ColorBACKGradient to return a text wis a colored gradient background 
	fmt.Println(ColorGradient.ColorBACKGradient("colors gradient as background", Orange_Purple))
		//using the ColorGradient to return a text wis a color gradient
	fmt.Println(ColorGradient.ColorGradient("colors gradient as text", Orange_Purple))

}
```
![](https://i.ibb.co/8DCbPwp/Capture-d-cran-2024-02-23-131706.png)
