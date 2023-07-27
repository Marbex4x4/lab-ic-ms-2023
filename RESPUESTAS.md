# Respuestas

Indica tu nombre a continuación: 

Por cada etapa agrega una sección abajo y escribe las respuestas a las preguntas de cada etapa.

## ETAPA 1

Escribe respuestas de la etapa 1 acá

## ETAPA 2

Escribe respuestas de la etapa 2 acá

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