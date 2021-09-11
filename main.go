package main

import (
	"fmt"
)


func ROI(ingresos, gastos float64) float64 {
	result := (ingresos - gastos ) / gastos * 100
	return result
}

func ROAS(ingresosVentas, gastos float64) float64 {
	result := ingresosVentas / gastos
	return result
}

func main() {
	// si mi inversión fue de 1000 dolares y obtuve 3000 de ganancia.

	inversionEmpresa := 1000
	ingresoEmpresa := 3000

	total := ROI(float64(ingresoEmpresa), float64(inversionEmpresa))
	fmt.Printf("Estás ganando un %v porciento \n", total)

	ingresoVentasEmpresa := 12000
	gastosEmpresa := 4500

	resultado := ROAS(float64(ingresoVentasEmpresa), float64(gastosEmpresa))
	fmt.Printf("Estás ganando un %v \n", resultado)
}