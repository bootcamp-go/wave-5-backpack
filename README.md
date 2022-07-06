# Bienvenidos al repositorio de la Wave 5 - Go
## A continuación te explicamos como vamos a utilizar este repo
---

Cada uno de ustedes deberá crear su rama de la siguiente manera _**apellido_nombre**_ con la misma estructura de master.

### _Recuerden remplazar **apellido_nombre** con su **apellido** y **nombre**._

Esto nos permitira un espacio donde puedan ir guardando sus avances en las practicas.

Dentro de cada modulo deben ir colocando en carpetas las practicas correspondientes a esa clase (clase**X**_parte**X**)

> Ejemplo: La practica de la clase 1 parte 1 (o mañana), se guarda en el directorio:
    **gobases/clase1_parte1**

Los pasos para empezar a utilizar la mochila son los siguientes:

1. Clonar el repositorio individual y su estructura.
    <pre><code>git clone git@github.com:bootcamp-go/wave-5-backpack.git</pre></code>
2. Cambiar al directorio del repositorio.
    <pre><code>cd wave-5-backpack</pre></code>
3. Crear una rama con la nombre apellido_nombre.
   
    *este comando crea un rama y cambia a esa*
    <pre><code>git checkout -b apellido_nombre</pre></code>
    *agregamos una linea al README para que git considere cambios*
    <pre><code>echo "## Esta es la mochila de apellido_nombre" >> README.md</pre></code>
    *añadimos los cambios al stash*
    <pre><code>git add . </pre></code>
    <pre><code>git commit -m "inital" </pre></code>
    <pre><code>git push origin apellido_nombre</pre></code>
Luego se deberá trabajar con el flujo habitual, por cada clase puede hacer los commits que necesiten:
    <pre><code> git pull origin apellido_nombre</pre></code>
    <pre><code> git add .</pre></code>
    <pre><code> git commit -m "mensaje_explicativo"</pre></code>
    <pre><code> git push origin apellido_nombre</pre></code>

    github.com/bootcamp-go/wave-5-backpack/tree/su_rama/goweb

## Esta es la mochila de Monay_Francisco