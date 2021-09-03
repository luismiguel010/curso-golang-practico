package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tNew *task){
	t.tasks = append(t.tasks, tNew)
}

func (t *taskList) eliminarDeLista(index int){
	t.tasks = append(t.tasks[:index], t.tasks[index + 1:]...)
}

func (t *taskList) imprimirLista(){
	for _, tarea := range t.tasks {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Tarea", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletado(){
	for _, tarea := range t.tasks {
		if tarea.completado{
			fmt.Println("Nombre", tarea.nombre)
			fmt.Println("Tarea", tarea.descripcion)
		}
	}
}

type task struct {
	nombre string
	descripcion string
	completado bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre: "Completar mi curso de Go",
		descripcion: "Completar el curso de Go esta semana",
	}

	t2 := &task{
		nombre: "Completar mi curso de Python",
		descripcion: "Completar el curso de Python esta semana",
	}

	t3 := &task{
		nombre: "Completar mi curso de NodeJS",
		descripcion: "Completar el curso de NodeJS esta semana",
	}
	
	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	lista.agregarALista(t3)
	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	//fmt.Println("\nTareas completadas")
	lista.imprimirListaCompletado()

	mapaTareas := make(map[string]*taskList)
	mapaTareas["Luis"] = lista

	t4 := &task{
		nombre: "Completar mi curso de C#",
		descripcion: "Completar el curso de C# esta semana",
	}

	t5 := &task{
		nombre: "Completar mi curso de Java",
		descripcion: "Completar el curso de Java esta semana",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Miguel"] = lista2

	fmt.Println("\nTareas de Luis")
	mapaTareas["Luis"].imprimirLista()

	fmt.Println("\nTareas de Miguel")
	mapaTareas["Miguel"].imprimirLista()


}