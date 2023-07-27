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

...