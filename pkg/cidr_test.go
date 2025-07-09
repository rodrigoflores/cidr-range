package cidr

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cidr", func() {
	Describe("newCidrFromString", func() {
		It("Creates a new Ip instance with the correct addressAsInt", func() {
			testCases := []struct {
				cidrExample       string
				firstIp           string
				lastIp            string
				networkIdentifier int
				amountOfAddresses int64
				networkMask       string
			}{
				{"0.0.0.0/0", "0.0.0.0", "255.255.255.255", 0, 4294967296, "0.0.0.0"},
				{"10.0.0.0/8", "10.0.0.0", "10.255.255.255", 8, 16777216, "255.0.0.0"},
				{"10.0.0.0/16", "10.0.0.0", "10.0.255.255", 16, 65536, "255.255.0.0"},
				{"10.0.1.0/16", "10.0.0.0", "10.0.255.255", 16, 65536, "255.255.0.0"}, // Should normalize to 10.0.0.0
				{"10.0.0.0/20", "10.0.0.0", "10.0.15.255", 20, 4096, "255.255.240.0"},
				{"10.0.0.0/24", "10.0.0.0", "10.0.0.255", 24, 256, "255.255.255.0"},
				{"10.0.0.0/32", "10.0.0.0", "10.0.0.0", 32, 1, "255.255.255.255"},
				{"192.168.1.100/24", "192.168.1.0", "192.168.1.255", 24, 256, "255.255.255.0"}, // Another normalization test
				{"128.0.0.0/1", "128.0.0.0", "255.255.255.255", 1, 2147483648, "128.0.0.0"},
				{"192.0.0.0/2", "192.0.0.0", "255.255.255.255", 2, 1073741824, "192.0.0.0"},
				{"224.0.0.0/3", "224.0.0.0", "255.255.255.255", 3, 536870912, "224.0.0.0"},
				{"240.0.0.0/4", "240.0.0.0", "255.255.255.255", 4, 268435456, "240.0.0.0"},
				{"248.0.0.0/5", "248.0.0.0", "255.255.255.255", 5, 134217728, "248.0.0.0"},
				{"252.0.0.0/6", "252.0.0.0", "255.255.255.255", 6, 67108864, "252.0.0.0"},
				{"254.0.0.0/7", "254.0.0.0", "255.255.255.255", 7, 33554432, "254.0.0.0"},
			}
			for _, tc := range testCases {
				result := NewCidrFromString(tc.cidrExample)
				Expect(result).To(Equal(Cidr{
					FirstIp:           NewIpFromString(tc.firstIp),
					LastIp:            NewIpFromString(tc.lastIp),
					NetworkIdentifier: tc.networkIdentifier,
					AmountOfAddresses: tc.amountOfAddresses,
					NetworkMask:       tc.networkMask,
				}), "CIDR %s failed", tc.cidrExample)
			}
		})
	})
})
