package main

import (
	"fmt"

	"./multimedia"
)

func main() {
	var respuesta int

	for ok := true; ok; ok = (respuesta == 1) {

		var (
			tituloImagen  = ""
			formatoImagen = ""
			canalesImagen = ""
			tituloAudio   = ""
			formatoAudio  = ""
			duracionAudio = ""
			tituloVideo   = ""
			formatoVideo  = ""
			framesVideo   = ""
		)

		fmt.Println("   IMAGEN ")
		fmt.Println("Ingresa titulo: ")
		fmt.Scan(&tituloImagen)

		fmt.Println("ingresa formato")
		fmt.Scan(&formatoImagen)

		fmt.Println("ingresa canales")
		fmt.Scan(&canalesImagen)

		fmt.Println("    AUDIO ")
		fmt.Println("ingresa titulo")
		fmt.Scan(&tituloAudio)

		fmt.Println("ingresa formato")
		fmt.Scan(&formatoAudio)

		fmt.Println("ingresa duracion")
		fmt.Scan(&duracionAudio)

		fmt.Println("    VIDEO ")
		fmt.Println("ingresa titulo")
		fmt.Scan(&tituloVideo)

		fmt.Println("ingresa formato")
		fmt.Scan(&formatoVideo)

		fmt.Println("ingresa frames")
		fmt.Scan(&framesVideo)

		fmt.Println("¿Deseas agregar otro valor? \n 1.SI \n 2.MOSTRAR MULTIMEDIA \n 3.SALIR")
		fmt.Scan(&respuesta)
		if respuesta == 2 {
			fm := multimedia.ContenidoWeb{Multimedias: []multimedia.Multimedia{

				&multimedia.Imagen{tituloImagen, formatoImagen, canalesImagen},
				&multimedia.Audio{tituloAudio, formatoAudio, duracionAudio},
				&multimedia.Video{tituloVideo, formatoVideo, framesVideo},
			}}

			fmt.Println(fm.Mostrar())
		}

	}
}
