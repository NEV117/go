package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

/* func main() {
	// Escanear cada puerto y hacer una conexión
	for i := 0; i < 100; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Println("Port", i, "is open")
	}
} */

var site = flag.String("site", "scanme.nmap.org", "url to scan")

// EXECUTE WITH:
// go run net/port.go --site=scanme.webscantest.com
// go run file --site=site to scan the ports
func main() {
	flag.Parse()
	var wg sync.WaitGroup
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}

			conn.Close()
			fmt.Println("Port", port, "is open")
		}(i)
	}
	wg.Wait()
}
