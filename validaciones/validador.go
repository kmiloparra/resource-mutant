package validaciones

import (
	"regexp"
	"strings"
)

func ValidacionNxN(dna []string) bool {
	var numeroFilas int = len(dna)
	for _, v := range dna {
		if len(v) != numeroFilas {
			return false
		}
	}
	return true
}

func ValidacionFilaVacia(dna []string) bool {
	for _, v := range dna {
		if len(strings.TrimSpace(v)) == 0 {
			return false
		}
	}
	return true
}

func ValidacionDominio(dna []string) bool {
	for _, v := range dna {
		matched, err := regexp.MatchString("[^ACGT]", strings.ToUpper(v))
		if err != nil || matched {
			return false
		}
	}
	return true
}


