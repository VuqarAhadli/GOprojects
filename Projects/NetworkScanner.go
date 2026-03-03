package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

func scanhost(host string, start int, step int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i <= 65535; i += step {
		address := host + ":" + strconv.Itoa(i)

		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		if err != nil {
			continue
		}
		fmt.Printf("\x1b[0;42mOPEN\x1b[0m %v\n", address)

		conn.Close()
	}
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	var wg sync.WaitGroup
	host := GetOutboundIP().String()
	workers := 20

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go scanhost(host, i+1, workers, &wg)
	}

	wg.Wait()
}

/*

Copyright - Vugar Ahadli 2026

License - BSD clause 3
Copyright 2026 Vugar Ahadli

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS “AS IS” AND ANY EXPRESS OR IMPLIED WARRANTIES,
INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY,
OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR
TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY
OF SUCH DAMAGE.

*/
