# go-homebrew-json

Unmarshal JSON output from homebrew to golang struct

## Usage

1. Get the output from Homebrew:

```bash
brew info --json=v1 <formula name>
```

The output will be a big blob of json

2. Parse json:

```go
packages, err := homebrew.Unmarshal(jsonData)
```

---

The output from homebrew can be tweaked a bit. You can request a specific formula, installed formula etc. Take a look at [the homebrew documentation][querying-homebrew] for details

## Example

```go
package main

import (
	"log"
	"os/exec"

	"github.com/kylelemons/godebug/pretty"
	"github.com/wercker/go-homebrew-json/homebrew"
)

func main() {
	cmd := exec.Command("brew", "info", "--json=v1", "wercker-cli")
	json, err := cmd.Output()
	if err != nil {
		log.Panic(err)
	}

	data, err := homebrew.Unmarshal(json)
	if err != nil {
		log.Panic(err)
	}

	pretty.Print(data)
}
```


[querying-homebrew]: https://github.com/Homebrew/brew/blob/master/share/doc/homebrew/Querying-Brew.md
