package main

import (
	"fmt"
	"github.com/kmiloparra/resource-mutant/constantes"
	"github.com/kmiloparra/resource-mutant/utilidades"
)

func main(){
	fmt.Println(utilidades.EncontrarIncidenciasHash(utilidades.StringToMap(constantes.SECUENCIAS_GEN_MUTANTE_HASHMAP),utilidades.StringToMap(constantes.SECUENCIAS_INVALIDAS),"GGGGAGTA"))
	fmt.Println(utilidades.EncontrarIncidenciasHash(utilidades.StringToMap(constantes.SECUENCIAS_GEN_MUTANTE_HASHMAP),utilidades.StringToMap(constantes.SECUENCIAS_INVALIDAS),"TCGAGGGG"))
}



