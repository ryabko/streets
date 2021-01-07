# Data Converter
This is a utility to convert Streets Database from TSV files exported from MySQL to JSON files.

### Usage
In the project root directly create a go file with the following code.
```go
package main

import (
	"fmt"
	"ru.kalcho/streets/data/converter"
)

func main() {
	fmt.Println("Converting tsv to json...")
	err := converter.ConvertTsv2Json("data/tsv", "data/json")
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
```
