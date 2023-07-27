# Respuestas

Indica tu nombre a continuación: 

Por cada etapa agrega una sección abajo y escribe las respuestas a las preguntas de cada etapa.

## ETAPA 1

Escribe respuestas de la etapa 1 acá

¿Cuál es la diferencia entre los archivos con el verbo `Create` con los archivos con el verbo `Add`?
R= Los create crean objetos en la db, y los add agregan nuevos registros

¿Cómo se llama el servicio que se declara en el archivo docker-compose.yml?
R= servicio flyway

¿Cuál es el comando que se ejecuta en el servicio declarado?
R= -locations=filesystem:/flyway/sql -connectRetries=60 migrate

## ETAPA 2

Escribe respuestas de la etapa 2 acá

¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
R= Se debe actualizar la dependencia que tiene el servicio flyway y la variable de entorno POSTGRES_SERVER

## ETAPA 3

Revisa el archivo movies-api/Dockerfile.

¿Qué te llama la atención? 
R= Tiene 2 etapas, en la primera crea el build en base a una imagen de go, y en la segunda etapa genera la imagen final para el deploy con solo los binarios

Revisa el archivo docker-compose.yml.

¿Cómo se relacionan el archivo docker-compose.yml y el archivo movies-api/Dockerfile?
R= El docker-compose.yml construye el contenedor en base a la imagen generada por el  movies-api/Dockerfile

¿Qué crees que hace el atributo context debajo de build (está en la linea 6 del archivo docker-compose.yml)? 
R= define la ruta al directorio que contiene el Dockerfile 

## ETAPA 4

Compara los archivos Dockerfile de movies-api y movies-front. R= El Dockerfile del front no crea artefacto(binario), ya que regularmente para Node el despliegue se realiza enviando el código fuente y los archivos de configuración, instalando las dependencias para ejecutar el programa, a diferencia del Dockerfile de la api que por ser para go si tiene el build.

Compara el atributo build del servicio movies-api con el de movies-front. ¿Cuál es la diferencia? ¿Qué pasa si los dejas iguales?R= Las 2 maneras son correctas ya que en el atributo buil se debe especificar la ruta al archivo Dockerfile, y esto puede ser indicando directamente la ruta(movies-front) o haciendo uso del atributo context con el que tambien se señala la ruta del dockerfile(movies-api), se pueden dejar iguales usando una u otra manera y funcionará igual.

