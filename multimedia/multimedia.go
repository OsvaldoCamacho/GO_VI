package multimedia

type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion string
}

type Video struct {
	Titulo  string
	Formato string
	Frames  string
}

func (i *Imagen) Mostrar() string {

	return "   Imagen \n" + "Titulo: " + i.Titulo + "\n" +
		"Formato: " + i.Formato + "\n" +
		"Canales: " + i.Canales + "\n"
}

func (a *Audio) Mostrar() string {

	return "   Audio \n" + "Titulo: " + a.Titulo + "\n" +
		"Formato: " + a.Formato + "\n" +
		"Duracion: " + a.Duracion + "\n"
}

func (v *Video) Mostrar() string {

	return "   Video \n" + "Titulo: " + v.Titulo + "\n" +
		"Formato: " + v.Formato + "\n" +
		"Frames: " + v.Frames + "\n"
}

type Multimedia interface {
	Mostrar() string
}

type ContenidoWeb struct {
	Multimedias []Multimedia
}

func (mt *ContenidoWeb) Mostrar() string {

	var mostrarTodo string

	for _, m := range mt.Multimedias {
		mostrarTodo += m.Mostrar()
	}
	return mostrarTodo
}
