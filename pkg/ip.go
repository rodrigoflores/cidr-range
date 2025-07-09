package cidr

import (
	"fmt"
	"strconv"
	"strings"
)

type Ip struct {
	AddressAsInt int
}

func IpToInteger(ip string) int {
	parts := strings.Split(ip, ".")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	w, _ := strconv.Atoi(parts[3])
	return x*256*256*256 +
		y*256*256 +
		z*256 +
		w
}

func IntegerToIp(ipInt int) string {
	firstOctet := ipInt / (256 * 256 * 256)
	secondOctet := (ipInt % (256 * 256 * 256)) / (256 * 256)
	thirdOctet := (ipInt % (256 * 256)) / 256
	fourthOctet := ipInt % 256

	return fmt.Sprintf("%d.%d.%d.%d", firstOctet, secondOctet, thirdOctet, fourthOctet)
}

func (i Ip) IpAsString() string {
	return IntegerToIp(i.AddressAsInt)
}

func (i Ip) IpAsBinary() string {
	return fmt.Sprintf("%032b", i.AddressAsInt)
}

func NewIpFromString(ipString string) Ip {
	return Ip{
		AddressAsInt: IpToInteger(ipString),
	}
}

func NewIpFromInt(ipAsInt int) Ip {
	return Ip{
		AddressAsInt: ipAsInt,
	}
}
