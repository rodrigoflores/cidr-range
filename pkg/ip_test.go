package cidr

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ip", func() {
	Describe("newIpFromInt", func() {
		It("Creates a new Ip instance with the correct addressAsInt", func() {
			testCases := []struct {
				ipAsInt  int
				asBinary string
				asString string
			}{
				{0, "00000000000000000000000000000000", "0.0.0.0"},
				{255, "00000000000000000000000011111111", "0.0.0.255"},
				{4294967295, "11111111111111111111111111111111", "255.255.255.255"},
				{65280, "00000000000000001111111100000000", "0.0.255.0"},
				{16711680, "00000000111111110000000000000000", "0.255.0.0"},
				{4278190080, "11111111000000000000000000000000", "255.0.0.0"},
			}

			for _, tc := range testCases {
				result := NewIpFromInt(tc.ipAsInt)
				Expect(result).To(Equal(Ip{AddressAsInt: tc.ipAsInt}))
				Expect(result.IpAsBinary()).To(Equal(tc.asBinary))
				Expect(result.IpAsString()).To(Equal(tc.asString))
			}
		})
	})
	Describe("newIpFromString", func() {
		It("Creates a new Ip instance with the correct addressAsInt", func() {
			testCases := []struct {
				ip       string
				expected int
				asBinary string
				asString string
			}{
				{"0.0.0.0", 0, "00000000000000000000000000000000", "0.0.0.0"},
				{"10.0.0.0", 167772160, "00001010000000000000000000000000", "10.0.0.0"},
				{"20.0.0.0", 335544320, "00010100000000000000000000000000", "20.0.0.0"},
				{"30.0.0.0", 503316480, "00011110000000000000000000000000", "30.0.0.0"},
				{"50.0.0.0", 838860800, "00110010000000000000000000000000", "50.0.0.0"},
				{"0.0.0.255", 255, "00000000000000000000000011111111", "0.0.0.255"},
				{"0.0.255.0", 65280, "00000000000000001111111100000000", "0.0.255.0"},
				{"0.255.0.0", 16711680, "00000000111111110000000000000000", "0.255.0.0"},
				{"255.0.0.0", 4278190080, "11111111000000000000000000000000", "255.0.0.0"},
			}

			for _, tc := range testCases {
				result := NewIpFromString(tc.ip)

				Expect(result).To(Equal(Ip{AddressAsInt: tc.expected}), "%s: Expected %d to be %d", tc.ip, result.AddressAsInt, tc.expected)
				Expect(result.IpAsBinary()).To(Equal(tc.asBinary), "%s: Expected %s to be %s", tc.ip, result.IpAsBinary(), tc.asBinary)
				Expect(result.IpAsString()).To(Equal(tc.asString), "%s: Expected %s to be %s as %s", tc.ip, result.IpAsString(), tc.asString)
			}
		})
	})
	Describe("IP Conversion", func() {
		Context("IpToInteger", func() {
			It("Convert valid ips correctly", func() {
				testCases := []struct {
					ip       string
					expected int
				}{
					{"0.0.0.0", 0},
					{"192.168.1.1", 3232235777},
					{"10.0.0.0", 167772160},
					{"255.255.255.255", 4294967295},
				}

				for _, tc := range testCases {
					result := IpToInteger(tc.ip)
					Expect(result).To(Equal(tc.expected), "Failed for IP => Int: %s/%d", tc.ip, tc.expected)
				}
			})
		})
	})
})
