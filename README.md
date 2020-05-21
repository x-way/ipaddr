# ipaddr - human friendly list of IP addresses and network interfaces
[![CircleCI](https://circleci.com/gh/x-way/ipaddr.svg?style=svg)](https://circleci.com/gh/x-way/ipaddr)
[![Go Report Card](https://goreportcard.com/badge/github.com/x-way/ipaddr)](https://goreportcard.com/report/github.com/x-way/ipaddr)

ipaddr - human friendly list of IP addresses and network interfaces

ipaddr is a convenience tool to quickly show the configured IP addresses with their interfaces on the current system.
It's mostly helpful for people working with systems where the output of `ip addr show` or `ifconfig` is too big to be human readable.

## Installation
```
# go get -u github.com/x-way/ipaddr
```

## Usage
```
Usage: ipaddr [-6] [-l] [<interface>]

Parameters:
  -6	also show IPv6 addresses
  -l	also show IPv6 link-local addresses
```

By default IPv4 addresses of all interfaces are shown
```
# ipaddr
lo          127.0.0.1/8
ens5        198.51.100.160/24
tun24008    10.123.199.78/32
tun71991639 10.200.123.5/32
tun26724    10.100.100.235/32
tun3883710  10.123.111.7/32
```

With the name as parameter only this interface is shown
```
# ipaddr ens5
ens5 198.51.100.160/24
```

Use the `-6` flag to also show IPv6 addresses
```
# ipaddr -6
lo          127.0.0.1/8
            ::1/128
ens5        198.51.100.160/24
            2a001:db0:0:1234::1/64
tun24008    10.123.199.78/32
            fdd9:349c:e7dd:248::78/64
tun71991639 10.200.123.5/32
            fdd9:349c:e7dd:719::5/64
tun26724    10.100.100.235/32
            fdd9:349c:e7dd:267::235/64
tun3883710  10.123.111.7/32
            fdd9:349c:e7dd:388::7/64
```

Use the `-l` flag to also show IPv6 link-local addresses
```
# ipaddr -6 -l
lo          127.0.0.1/8
            ::1/128
ens5        198.51.100.160/24
            2a001:db0:0:1234::1/64
            fe80::a9cd:fcff:fc68:fcb/64
tun24008    10.123.199.78/32
            fdd9:349c:e7dd:248::78/64
            fe80::de9:248f/64
tun71991639 10.200.123.5/32
            fdd9:349c:e7dd:719::5/64
            fe80::de9:719f/64
tun26724    10.100.100.235/32
            fdd9:349c:e7dd:267::235/64
            fe80::de9:267f/64
tun3883710  10.123.111.7/32
            fdd9:349c:e7dd:388::7/64
            fe80::de9:388f/64
veth4020119 fe80::8b4b:8ff:fd1c:59ca/64
veth8ba0ee1 fe80::d5cb:c7ff:fe08:3f/64
veth57f5701 fe80::8df8:a5ff:fdb6:42f6/64
veth4bb8ef3 fe80::db6e:bcff:fe48:4b71/64
vethb812901 fe80::69f6:8cff:fdde:67a3/64
veth00898c8 fe80::ed12:59ff:fe68:ebd7/64
```
