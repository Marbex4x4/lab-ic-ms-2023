# Respuestas

Indica tu nombre a continuación: 

Por cada etapa agrega una sección abajo y escribe las respuestas a las preguntas de cada etapa.

## ETAPA 1

Escribe respuestas de la etapa 1 acá

## ETAPA 2

Escribe respuestas de la etapa 2 acá

## ETAPA 4

Compara los archivos Dockerfile de movies-api y movies-front. R= El Dockerfile del front no crea artefacto(binario), ya que regularmente para Node el despliegue se realiza enviando el código fuente y los archivos de configuración, instalando las dependencias para ejecutar el programa, a diferencia del Dockerfile de la api que por ser para go si tiene el build.

Compara el atributo build del servicio movies-api con el de movies-front. ¿Cuál es la diferencia? ¿Qué pasa si los dejas iguales?R= Las 2 maneras son correctas ya que en el atributo buil se debe especificar la ruta al archivo Dockerfile, y esto puede ser indicando directamente la ruta(movies-front) o haciendo uso del atributo context con el que tambien se señala la ruta del dockerfile(movies-api), se pueden dejar iguales usando una u otra manera y funcionará igual.