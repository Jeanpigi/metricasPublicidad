package main

import "testing"

func TestTotalROI(t *testing.T)  {
	tables := []struct {
		ingresos float64
		gastos float64
		result float64
	} {
		{2500, 1000, 150},
		{3000, 1000, 200},
	}

	for _, item := range tables {
		total := ROI(item.ingresos, item.gastos)
		if total != item.result {
			t.Errorf("Sum was incorrect, got %v, expected %v", total, item.result)
		}
	}
}

func TestTotalROAS(t *testing.T) {
	tables := []struct {
		ingresos float64
		gastos float64
		result float64
	} {
		{12000, 4500, 2.67},
		{13000, 5000, 2.6},
	}

	for _, item := range tables {
		total := ROAS(item.ingresos, item.gastos)
		if total != item.result {
			t.Errorf("Sum was incorrect, got %v, expected %v", total, item.result)
		}
	}
}