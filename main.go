package main

import (
	"fmt"
	"math"
)


func ROI(ingresos, gastos float64) float64 {
	result := (ingresos - gastos ) / gastos * 100
	return result
}

func ROAS(ingresosVentas, gastos float64) float64 {
	result := ingresosVentas / gastos
	// redondear en dos decimales
	result = math.Round(result * 100) / 100
	return result
}

func main() {

	inversionEmpresa := 1000
	ingresoEmpresa := 3000

	total := ROI(float64(ingresoEmpresa), float64(inversionEmpresa))
	fmt.Printf("Estás ganando un %v porciento \n", total)

	ingresoVentasEmpresa := 12000
	gastosEmpresa := 4500

	resultado := ROAS(float64(ingresoVentasEmpresa), float64(gastosEmpresa))
	fmt.Printf("Estás ganando un %v \n", resultado)
}