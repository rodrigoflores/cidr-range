package main

import (
	"fmt"
	"os"

	cidr "github.com/rodrigoflores/cidr-range/pkg"
)

func main() {
	// Check if CIDR argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <CIDR>")
		fmt.Println("Example: go run main.go 10.0.0.0/8")
		os.Exit(1)
	}

	// Get CIDR from command line argument
	cidrS := os.Args[1]
	parsedCidr := cidr.NewCidrFromString(cidrS)

	fmt.Printf("CIDR: %s\n", cidrS)
	fmt.Println("First IP as integer:", parsedCidr.FirstIp)
	fmt.Println("First IP:", parsedCidr.FirstIp.IpAsString())
	fmt.Printf("First IP as binary: %s\n", parsedCidr.FirstIp.IpAsBinary())
	fmt.Printf("Last IP as integer: %d\n", parsedCidr.LastIp.AddressAsInt)
	fmt.Printf("Last IP: %s\n", parsedCidr.LastIp.IpAsString())
	fmt.Printf("Last IP as binary: %s\n", parsedCidr.LastIp.IpAsBinary())
	fmt.Println("Size of the network:", parsedCidr.AmountOfAddresses)
	fmt.Printf("Network Mask: %s\n", parsedCidr.NetworkMask)
}
