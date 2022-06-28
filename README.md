# Bienvenidos al repositorio de la Wave 5 - Go
## A continuación te explicamos como vamos a utilizar este repo
---

Cada uno de ustedes deberá crear su rama de la siguiente manera _**apellido_nombre**_ con la misma estructura de master.

### _Recuerden remplazar **apellido_nombre** con su **apellido** y **nombre**._

Esto nos permitira un espacio donde puedan ir guardando sus avances en las practicas.

Los pasos para empezar a utilizar la mochila son los siguientes:

1. Clonar el repositorio individual y su estructura.
    <pre><code>git clone git@github.com:bootcamp-go/wave-5-backpack.git</pre></code>
2. Crear una rama con la nombre apellido_nombre.
   
    *este comando crea un rama y cambia a esa*
    <pre><code>git checkout -b apellido_nombre</pre></code>
    *agregamos una linea al README para que git considere cambios*
    <pre><code>echo "## Esta es la mochila de apellido_nombre" >> README.md</pre></code>
    *añadimos los cambios al stash*
    <pre><code>git add . </pre></code>
    <pre><code>git commit -m "inital" </pre></code>
    <pre><code>git push origin apellido_nombre</pre></code>
Luego se deberá trabajar con el flujo habitual.
    <pre><code> git pull origin apellido_nombre</pre></code>
    <pre><code> git add .</pre></code>
    <pre><code> git commit -m "mensaje"</pre></code>
    <pre><code> git push origin apellido_nombre</pre></code>
