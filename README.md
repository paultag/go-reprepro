go-reprepro
===========

```
package main

import (
	"fmt"

	"pault.ag/go/reprepro"
)

func main() {
	repo := reprepro.NewRepo("/home/tag/tmp/repo/")
	err := repo.Include("unstable", "/home/tag/dev/debian/golang/golang-pault-go-debian_0.1+git20150726.2.66db6de-1_source.changes")
	fmt.Printf("%s\n", err)
}
```
