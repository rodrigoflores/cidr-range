package cidr

import (
	"strconv"
	"strings"
)

type Cidr struct {
	FirstIp           Ip
	NetworkIdentifier int
	LastIp            Ip
	AmountOfAddresses int64
	NetworkMask       string
}

func NewCidrFromString(cidrString string) Cidr {
	cidrParts := strings.Split(cidrString, "/")
	baseIp := cidrParts[0]
	networkIdentifier, _ := strconv.Atoi(cidrParts[1])
	amountOfAddresses := 1 << (32 - networkIdentifier)
	zeroBits := 32 - networkIdentifier
	firstIp := NewIpFromInt((NewIpFromString(baseIp).AddressAsInt >> zeroBits) << zeroBits)

	return Cidr{
		FirstIp:           firstIp,
		NetworkIdentifier: networkIdentifier,
		AmountOfAddresses: int64(amountOfAddresses),
		LastIp:            NewIpFromInt(firstIp.AddressAsInt + amountOfAddresses - 1),
		NetworkMask:       IntegerToIp((1 << 32) - amountOfAddresses),
	}
}
