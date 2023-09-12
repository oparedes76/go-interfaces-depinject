package entities

import "fmt"

type Avion struct {
	Nombre    string
	Marca     string
	Modelo    string
	Capacidad int
}

func NewAvion(nombre string, marca string, modelo string, capacidad int) *Avion {
	return &Avion{
		Nombre:    nombre,
		Marca:     marca,
		Modelo:    modelo,
		Capacidad: capacidad,
	}
}

func (a Avion) Mover() {
	fmt.Println("El avión", a.Nombre, "se está moviendo.")
}

func (a Avion) Fabricar() {
	fmt.Println("El avión", a.Nombre, "se está fabricando con capacidad de", a.Capacidad, "pasajeros")
}
