package utilidades

import (
	"fmt"
	"github.com/kmiloparra/resource-mutant/constantes"
	"strings"
)

func Version() string  {

	return "v1.1.2"
}

// pivotea una matrix cuadrada
func PivotearMatrix(dna []string) []string {
	var numeroFilas int = len(dna)
	var dnaTranspuesta []string

	for i := 0; i < numeroFilas; i++ {
		var columna strings.Builder
		columna.Grow(numeroFilas)
		for _, f := range dna {
			columna.WriteString(string(f[i]))
		}
		dnaTranspuesta = append(dnaTranspuesta, columna.String())
	}
	return dnaTranspuesta
}

// Metodo para sacar las diagonales validas para evaluar de la matriz
func ObtenerDiagonales(dna []string) []string {
	var numeroFilas int = len(dna)
	var diagonales []string
	var puntoPartida = numeroFilas - constantes.TAMANIO_SECUENCIA_MUTANTE
	posicionMaxima := numeroFilas - 1

	for y := 0; y <= puntoPartida; y++ {
		var diagonalInferiorDescendente strings.Builder
		diagonalInferiorDescendente.Grow(numeroFilas)
		var diagonalSuperiorDescendente strings.Builder
		diagonalSuperiorDescendente.Grow(numeroFilas)
		var diagonalInferiorAscendente strings.Builder
		diagonalInferiorAscendente.Grow(numeroFilas)
		var diagonalSuperiorAscendente strings.Builder
		diagonalSuperiorAscendente.Grow(numeroFilas)
		xIterador := 0
		yIterador := y
		for x := 0; x < numeroFilas; x++ {
			if x+y < numeroFilas {
				diagonalInferiorDescendente.WriteString(string(dna[yIterador][xIterador]))
				diagonalSuperiorDescendente.WriteString(string(dna[xIterador][yIterador]))
				diagonalInferiorAscendente.WriteString(string(dna[yIterador][posicionMaxima-xIterador]))
				diagonalSuperiorAscendente.WriteString(string(dna[posicionMaxima-yIterador][xIterador]))
			} else {
				break
			}
			yIterador++
			xIterador++
		}
		diagonales = append(diagonales,
			diagonalInferiorDescendente.String(),
			diagonalInferiorAscendente.String(),
			diagonalSuperiorDescendente.String(),
			diagonalSuperiorAscendente.String())
	}
	// devuelvo el slice pero eliminando las diagonales repetidas
	return diagonales[2:]
}

func EncontrarIncidenciasHash(secuenciasValidas, secuenciasinvalidas map[string]string, cadena string) uint8 {
	var tamanioCadena int = len(cadena)
	var numeroIncidencias uint8 = 0

	for i := 0; i < tamanioCadena-(constantes.TAMANIO_SECUENCIA_MUTANTE-1); i++ {
		subCadena := cadena[i : i+constantes.TAMANIO_SECUENCIA_MUTANTE]

		resultado, noExiste := secuenciasValidas[subCadena]
		if noExiste {
			fmt.Println("algo", resultado)
			numeroIncidencias++
			if i < tamanioCadena-constantes.TAMANIO_SECUENCIA_MUTANTE && string(cadena[i+constantes.TAMANIO_SECUENCIA_MUTANTE]) == resultado {
				return constantes.CANTIDAD_SECUENCIA_MUTANTE
			}
			i += constantes.TAMANIO_SECUENCIA_MUTANTE
		} else {
			_, noExiste = secuenciasinvalidas[subCadena[2:]]
			if noExiste {
				i += constantes.CANTIDAD_SECUENCIA_MUTANTE
			}
		}
		fmt.Println("i",i)
	}

	return numeroIncidencias
}

func StringToMap(keyValues string) map[string]string {
	mapResultante := make(map[string]string)
	tuplas := strings.Split(keyValues, ",")
	tupla := make([]string, 2)
	for _, v := range tuplas {
		tupla = strings.Split(v, ":")
		mapResultante[tupla[0]] = tupla[1]
	}
	return mapResultante
}

