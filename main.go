package main

import (
	"go-interfaces/entities"
)

func main() {
	// Creamos un automóvil y un avión
	automovil := entities.Automovil{
		Nombre: "Ford Mustang",
		Marca:  "Ford",
		Modelo: "2023",
		Color:  "Rojo",
	}
	avion := entities.Avion{
		Nombre:    "Boeing 747",
		Marca:     "Boeing",
		Modelo:    "2022",
		Capacidad: 500,
	}

	// Usamos la interfaz Mover para mover los vehículos
	automovil.Mover()
	avion.Mover()
	automovil.Fabricar()
	avion.Fabricar()

	automovil2 := entities.NewAutomobil("Camaro", "Chevrolet", "2023", "Rojo")
	automovil2.Fabricar()
	automovil2.Mover()

	avion2 := entities.NewAvion("Airbus A310", "Airbus", "2023", 380)
	avion2.Fabricar()
	avion2.Mover()
}
