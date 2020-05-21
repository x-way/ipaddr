package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"strings"
)

func filterIPv6(addr string) bool {
	return !strings.Contains(addr, ":")
}

func filterLinkLocal(addr string) bool {
	return !strings.HasPrefix(addr, "fe80:")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-6] [-l] [<interface>]\n\nParameters:\n", os.Args[0])
		flag.PrintDefaults()
	}
	showIPv6 := flag.Bool("6", false, "also show IPv6 addresses")
	showLinkLocal := flag.Bool("l", false, "also show IPv6 link-local addresses")
	flag.Parse()

	filter := func(_ string) bool {
		return true
	}
	if !*showIPv6 {
		filter = filterIPv6
	} else if !*showLinkLocal {
		filter = filterLinkLocal
	}

	ifaces := getInterfaces(flag.Args())
	maxlen := longestName(ifaces)
	for _, iface := range ifaces {
		printIface(maxlen, iface, filter)
	}
}

func getInterfaces(args []string) []net.Interface {
	var ifaces []net.Interface
	var err error
	if len(args) == 0 {
		ifaces, err = net.Interfaces()
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, ifname := range args {
		iface, err := net.InterfaceByName(ifname)
		if err != nil {
			log.Fatal(err)
		}
		ifaces = append(ifaces, *iface)
	}
	return ifaces
}

func longestName(ifaces []net.Interface) int {
	length := 0
	for _, iface := range ifaces {
		if len(iface.Name) > length {
			length = len(iface.Name)
		}
	}

	return length
}

func printIface(namelen int, iface net.Interface, filter func(string) bool) {
	addrs, err := iface.Addrs()
	if err != nil {
		log.Fatal(err)
	}
	var strAddrs []string
	for _, addr := range addrs {
		strAddr := addr.String()
		if filter(strAddr) {
			strAddrs = append(strAddrs, strAddr)
		}
	}
	if len(strAddrs) == 0 {
		return
	}
	sort.Strings(strAddrs)
	fmtStr := fmt.Sprintf("%%-%ds %%s\n", namelen)
	ifname := iface.Name
	for i, addr := range strAddrs {
		if i > 0 {
			ifname = ""
		}
		fmt.Printf(fmtStr, ifname, addr)
	}
}
