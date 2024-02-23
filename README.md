# ColorGradient

pakage to creat colors gradient to make your code pretty :)
<img width="187" alt="Capture d'écran 2024-02-23 123000" src="https://github.com/edenbwt/ColorGradient/assets/94796854/1a1e14d8-9607-4d0c-8a42-026275cb9086">
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
<img width="300" alt="Capture d'écran 2024-02-23 131706" src="https://github.com/edenbwt/ColorGradient/assets/94796854/5093e311-8a09-40eb-ae43-cd6fa3b50605">

