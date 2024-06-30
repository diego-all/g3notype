

prompt

Requiero un programa de tipo CLI en golang utiliando la libreria spf13/cobra que permita capturar comandos para ejecutar ciertas funciones. Requiero que en el programa main exista la forma de recibir el parametro init en conjunto con el nombre del proyecto (my-ska) y el flag -db con valor ingresado por el usuario. Este programa debera permitir ser ejecutado haciendo un llamado al repositorio de github donde se encuentre la aplicacion. go run github.com/diego-all/run-from-github@latest init --db postgres my-ska Este comando init llamara la funcion generate() que estará en el paquete generator del programa. Podrias darme la respuesta en español por favor.

Podrias darme el codigo fuente por favor.


docker run --rm \
-it -p 8400-8500:8400-8500 \
-v ~/.msf4:/root/.msf4 \
-v /tmp/msf:/tmp/data \
phocean/msf



## spf13/cobra

Duda orden en los parametros.

**Argumentos Posicionales vs. Flags**

**Posicionales** son aquellos que no llevan un prefijo con guiones (--). Simplemente se pasan en el orden esperado por el comando. En tu caso, init espera un argumento posicional para el nombre del proyecto.

**Flags** son argumentos que tienen un nombre precedido por uno o dos guiones (- o --) y generalmente se usan para opciones que pueden o no ser proporcionadas. En tu comando, --db y --config son flags.




## Orden de envio de parametros de ejecucion

    go run main.go init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json projectTest

Recordar si se usa el de GitHub se debe de tener el repo actualizado.

    go run github.com/diego-all/run-from-gh@latest init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json projectTest


