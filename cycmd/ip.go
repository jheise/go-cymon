package main

import (
	"fmt"
)

func processIP(target string, section string) {
	switch section {
	case "lookup":
		handleIPLookup(target)
	case "events":
		handleIPEvents(target)
	case "domains":
		handleIPDomains(target)
	case "urls":
		handleIPUrls(target)
	case "blacklist":
		handleIPBlacklist(target)
	default:
		fmt.Println("IP doesnt have report: %s\n", section)
	}
}

func handleIPLookup(target string) {
	data, err := client.GetIPAddrLookup(target)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Address: %s\n", data.Addr)
	fmt.Printf("Created: %s\n", data.Created)
	fmt.Printf("Updated: %s\n", data.Updated)
	fmt.Printf("Events: %s\n", data.Events)
	fmt.Printf("Domains: %s\n", data.Domains)
	fmt.Printf("Urls: %s\n", data.Urls)
	fmt.Printf("Sources:\n")
	for _, x := range data.Sources {
		fmt.Printf("  - %s\n", x)
	}

}

func handleIPEvents(target string) {
	fmt.Println(client.GetIPAddrEvents(target))
}

func handleIPUrls(target string) {
	fmt.Println(client.GetIPAddrUrls(target))
}

func handleIPDomains(target string) {
	fmt.Println(client.GetIPAddrDomains(target))
}

func handleIPBlacklist(target string) {
	fmt.Println(client.GetIPAddrBlacklist(target, 1))
}
