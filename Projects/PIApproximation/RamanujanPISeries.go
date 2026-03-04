package main

import (
	"fmt"
	"math"
)

func factorialfloat(k int) float64 {
	result := 1.0
	for i := 1; i <= k; i++ {
		result *= float64(i)
	}
	return result
}

func ram(k int) float64 {
	numerator := factorialfloat(4*k) * float64(1103+26390*k)
	denominator := math.Pow(factorialfloat(k), 4) * math.Pow(396, float64(4*k))
	return numerator / denominator
}

func main() {
	sum := 0.0
	for k := 0; k < 4; k++ {
		sum += ram(k)
		pi := 1 / ((2 * math.Sqrt(2) / 9801) * sum)
		fmt.Printf("Term: %d  π = %.15f\n", k+1, pi)
	}
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
