package DatosEntrenamiento

type DatosEntrenamiento struct {
	hr_recuperacion []int
	tiempo          []int
}

func (datos DatosEntrenamiento) calcular_PulsacionesPorMinuto() (media int) {
	for i := 0; i < len(datos.hr_recuperacion); i++ {
		for j := 0; j < len(datos.tiempo); j++ {

		}
	}
}
