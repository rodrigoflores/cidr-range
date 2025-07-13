# cidr-range

Given a [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) (such as 10.0.0.0/8), generates which is the first and the last IP, the Integer and Binary representations for both, the network size and the network mask.

```
$ cidr-range 10.0.0.0/8
CIDR: 10.0.0.0/8
First IP as integer: {167772160}
First IP: 10.0.0.0
First IP as binary: 00001010000000000000000000000000
Last IP as integer: 184549375
Last IP: 10.255.255.255
Last IP as binary: 00001010111111111111111111111111
Size of the network: 16777216
Network Mask: 255.0.0.0
```

Only supports IPv4 for now.