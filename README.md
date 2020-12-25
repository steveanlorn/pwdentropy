# pwdentropy
Responsible to measure password strength using entropy value.

## How to Use
Example:  
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/steveanlorn/pwdentropy"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "password is required")
		os.Exit(2)
	}

	for _, password := range os.Args[1:] {
		entropy := pwdentropy.Calculate(password)
		log.Printf("Entropy %s: %.2f\n", password, entropy)
	}
}
```

### Entropy Threshold Value
![Entropy benchmark table](https://64.media.tumblr.com/05586038c5d783ad973c8f81234eb6f4/tumblr_nmktshSzuq1sq2igro1_1280.png)
You can refer to that table above for the entropy threshold value.

## Resource
- https://generatepasswords.org/how-to-calculate-entropy/
- https://en.wikipedia.org/wiki/List_of_Unicode_characters#Latin_script
- https://datarep.tumblr.com/post/116916012426/time-required-to-brute-force-crack-a-password
