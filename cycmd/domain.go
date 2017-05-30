package main

import (
	"fmt"
)

func processDomain(target string, section string) {
	data, err := client.GetDomainLookup(target)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Created: %s\n", data.Created)
	fmt.Printf("Updated: %s\n", data.Updated)

	fmt.Printf("Sources\n")
	for _, x := range data.Sources {
		fmt.Printf("  - %s\n", x)
	}

	fmt.Printf("IPs:\n")
	for _, x := range data.Ips {
		fmt.Printf("  - %s\n", x)
	}

	fmt.Printf("URLs:\n")
	for _, x := range data.Urls {
		fmt.Printf("  - %s\n", x)
	}
}
