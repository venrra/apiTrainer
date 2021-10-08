# 🏃‍♀️apiTrainer:running_man:

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

El estado físico y la vida saludable siempre son temas que nos interesan entre la mayoria de las personas, con lo que muchas muchas veces practicamos algun deporte o llevamos a cabo un entrenamiento. Siempre queremos mejorar y algunas veces puede llegar a ser frustrante o tedioso saber si nuestro entrenamiento es efectivo y da resultados o si debemos modificarlo.

## 📝🆕 Sobre el proyecto

__apiTrainer__ pretende convertirse en una herramienta para entrenadores y usuarios habituales, para facilitar la monitorización de sus entrenamientos, que quiere resolver y ofrecer:

- Asistencia en entrenamiento personal calculando estado fisico de un usuario en un punto concreto de su entrenamiento.

- Calcular el [valor de recuperación](./docs/terminologia.md#Calculo-del-valor/tiempo-de-recuperación-HR) de un usuario.

- Ofrecer estadísticas de entrenamiento por usuario.

- Calcular mejora fisica de un usuario.

- Análisis de los resultados de un conjunto de usuarios que realizan el mismo entrenamiento.

En concreto, se quiere ofrecer al usuario el cálculo del valor de recuperación, estadísticas de su entrenamiento y estadísticas sobre grupos de usuarios. Para ello se usarán valores de frecuencia cardiaca aportados por el usuario, que junto a caracteristicas como altura, peso o el sexo, dará un valor como el _tiempo de recuperación_ (entre otros). Con estos resultados se podrán construir estadisticas sobre el entrenamiento de un usuario.

Además se pretende NO anclar esta herramienta a una aplicación, es decir no solo un tipo de aplicación podrá usar la herramienta (app móvil, web...). Claramente un ejemplo de uso sería una aplicación móvil que a través de un sensor de frecuencia cardiaca podrá leer valores de HR y los enviase a __apiTrainer__, pudiendo calcular una sesión de entrenamiento, pero además generando estadísticas de todos los entrenamientos que se han realizado, no solo de un usuario sino de varios; a su vez, un entrenador podrá usar __apiTrainer__ para calcular estadísticas de sus clientes ya sea con aplicación móvil, web o lo que proceda.
 
Otro ejemplo sería el uso de __apiTrainer__ en dispositivos específicos de entrenamiento diseñados específicamente para el deporte pero con capacidad de comunicación.
 
La funcionalidad de __apiTrainer__ estaría descentralizada de un dispositivo pudiendo, crear entornos de entrenamiento más dinámicos.

## 🛠️ Estado del desarrollo.

Ahora mismo se están dando los primeros pasos y concretando el rumbo a seguir. Se están definiendo las primeras tareas y analizando las funcionalidades básicas [M1](https://github.com/venrra/apiTrainer/milestones/1) y [M2](https://github.com/venrra/apiTrainer/milestones/2).

Aún se está dando forma a la idea base. Se prevé un primer Releases en 3 semanas, con las primeras funcionalidades.

[Aqui](https://github.com/venrra/apiTrainer/milestones) se pueden ver los primeros objetivos y su estado. 

## 🆕Tareas pendientes

- [ ] [[HU2]](https://github.com/venrra/apiTrainer/issues/4) Como Laura quiero saber como representar los valores de frecuencia cardíaca para poder operar con ellos

- [ ] [[HU1]](https://github.com/venrra/apiTrainer/issues/3) Como Laura quiero saber mi valor de recuperación.

- [x] Crear docuementacion de personas, user-journey y definiciones o palabras clave. debe ser lo minimo para cubrir [HU1] y [H1]

## 📄Documentación

Podéis encontrar toda la documentación, terminología, biografías de usuarios en [/docs](./docs) y en la [wiki](https://github.com/venrra/apiTrainer/wiki). Lo más destacable:

- [Personas](./docs/personas.md)
- [Terminología](./docs/terminologia.md)
- [user journey](./docs/user-journey.md)

## Autor

- [Venrra](https://github.com/venrra/)
