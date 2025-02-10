# News feeder

API desarrollada para implementar los conceptos de concurrencia y channels própios de Golang. para esta POC se realizan llamados a 2 APIs externas:
- https://open-meteo.com/en/docs#hourly=temperature_2m&timezone=auto
- https://newsapi.org/

## Contexto
Hay escenarios en los cuales existe la necesidad de realizar tareas de forma concurrente, como consultar varios servicios externos y dar una respuesta al usuario pero la misma debe enviarse dentro de un tiempo acotado o se responde al usuario con timeout. Para la ejecución de tareas concurrentes, existe en Go las goroutines que son funciones que se ejecutan independientemente de la función principal (la que dispara a las mismas mediante la palabra reservada "go"). <br>
Al lanzarse estas funciones ocurre un problema, si la goroutine principal finaliza antes que las mismas, el programa termina y las mismas nunca llegarán a cumplir con su objetivo. Es en este momento que llega a nuestro rescate los channels, un tipo de dato que funciona como un pipe, permitiendo a las goroutines comunicarse entre sí. En las funciones que se encuentran en el package client van a encontrar el ejemplo de cómo se pueden usar los channels para este objetivo.
#### Nota
*El nombre del artefacto incluye mi repositorio personal de Github ya que el mismo me lo llevo de ejemplo*