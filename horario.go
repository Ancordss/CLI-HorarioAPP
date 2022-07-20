package main

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

var (
	data = [][]interface{}{
		{"Lu 18:00 || Mi 19:30", "Lunes     || Miercoles", "Programacion II", "https://meet.google.com/rop-rppv-mix", 1},
		{"Lu 19:30 || Vi 19:30", "Lunes     || Viernes", "Microeconomia", "Revisar canvas cambia cada semana", 2},
		{"Ma 18:00 || Ju 18:00", "Martes    || Jueves", "Estadistica", "https://bit.ly/3aIi6VO", 0},
		{"Ma 19:30 || Ju 19:30", "Martes    || Jueves", "Calculo II", "https://meet.google.com/eei-hfpf-ihu", 1},
		{"Mi 18:00 || Vi 18:00", "Miercoles || Viernes", "Fisica II", "https://bit.ly/3ANoyp5, Asistencia= https://bit.ly/3RCHNrc", 0},
	}

	styles = map[string]*simpletable.Style{

		"Unicode style": simpletable.StyleUnicode,
	}
)

func main() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#  HORARIO #"},
			{Align: simpletable.AlignCenter, Text: "#   DIAS   #"},
			{Align: simpletable.AlignCenter, Text: "#   CLASE  #"},
			{Align: simpletable.AlignCenter, Text: "#   LINK   #"},
			{Align: simpletable.AlignCenter, Text: "#  TAREAS  #"},
		},
	}

	subtotal := 0
	for _, row := range data {
		r := []*simpletable.Cell{
			//{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", row[0].(int))},
			{Align: simpletable.AlignCenter, Text: row[0].(string)},
			{Text: row[1].(string)},
			{Text: row[2].(string)},
			{Text: row[3].(string)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row[4])},
		}

		table.Body.Cells = append(table.Body.Cells, r)
		subtotal += row[4].(int)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 4, Text: "Created By Ancordss"},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("Total: %d", subtotal)},
		},
	}

	for _, s := range styles {
		//fmt.Println(n)

		table.SetStyle(s)
		table.Println()

		fmt.Println()
	}
}
