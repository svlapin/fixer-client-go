# fixer-client-go
fixer.io API client in Go

#### Example:
```go
package main

import (
	"fmt"
	"log"

	"github.com/svlapin/fixer-client-go"
)

func queryLatest() {
	resp, err := fixerClient.Latest("EUR")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Latest EUR to USD (%s): %v", resp.Date, resp.Rates["USD"]))
}

func queryForSpecificDate() {
	resp, err := fixerClient.ForDate("2017-06-01", "EUR")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("EUR to USD on 2017-06-01: %v", resp.Rates["USD"]))
}

func main() {
	queryLatest()
	queryForSpecificDate()
}
```

```
2017/09/13 09:58:34 Latest EUR to USD (2017-09-12): 1.1933
2017/09/13 09:58:34 EUR to USD on 2017-06-01: 1.1219
```
