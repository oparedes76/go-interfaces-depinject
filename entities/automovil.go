package entities

import "fmt"

type Automovil struct {
	Nombre string
	Marca  string
	Modelo string
	Color  string
}

func NewAutomobil(nombre string, marca string, modelo string, color string) *Automovil {
	return &Automovil{
		Nombre: nombre,
		Marca:  marca,
		Modelo: modelo,
		Color:  color,
	}
}

func (a Automovil) Mover() {
	fmt.Println("El automóvil", a.Nombre, "se está moviendo.")
}

func (a Automovil) Fabricar() {
	fmt.Println("El automobil", a.Nombre, "se está fabricando de color", a.Color)
}
