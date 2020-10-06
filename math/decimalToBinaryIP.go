// Dado um ip em decimal, retorna o mesmo ip com os seus respectivos
// números em binário com 8bits
//
// Exemplo:
// DecimalToBinaryIP("192.168.0.1") == "11000000.10101000.00000000.00000001"
package math

import (
	"strconv"
	"strings"
)

// DecimalToBinary é uma função que recebe um inteiro x onde 0 <= x <= 2^8
// retorna o valor em binário em uma string
func DecimalToBinary(value int) string {
	decimals := [8]int{128, 64, 32, 16, 8, 4, 2, 1}
	var num []string
	for _, v := range decimals {
		if value-v >= 0 {
			value -= v
			num = append(num, "1")
		} else {
			num = append(num, "0")
		}
	}

	return strings.Join(num, "")
}

// DecimalToBinaryIP é uma função que recebe um IP como string e retorna o IP
// com os valores em binário
func DecimalToBinaryIP(ip string) string {
	numIP := strings.Split(ip, ".")
	var binIP []string
	for i := 0; i < 4; i++ {
		v, _ := strconv.Atoi(numIP[i])
		binIP = append(binIP, DecimalToBinary(v))
	}

	return strings.Join(binIP, ".")
}
