#Go-Cymon
---

Golang api client for Cymon

---

Example

```
package main

import (

"github.com/go-cymon/api"

"fmt"

)

func main() {

    client := api.NewClient("YOUAPIKEYGOESHERE")

    fmt.Println(client.GetIPLookup("4.2.2.1")
}
```
