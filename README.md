# Actors in akka-models 

El modelo de actor en Go Lang es un modelo de concurrencia que trata a los actores como entidades independientes, cada una con su propio estado y comportamiento. GoLang proporciona un mecanismo de concurrencia ligero a través de gorutinas y canales, que se pueden utilizar para implementar el modelo de actor.

En este contexto, un actor es una unidad de cálculo independiente y concurrente que se comunica con otros actores mediante el paso de mensajes. Las gorutinas se pueden utilizar para representar actores y los canales facilitan la comunicación entre ellos. Cada actor tiene su propio estado y procesa mensajes de forma asincrónica.

Para implementar el modelo de actor en Go Lang, normalmente se crearía una rutina para cada actor y se utilizarían canales para enviar mensajes entre actores. Los mensajes contendrían información sobre la acción o comportamiento previsto del actor receptor.

Este programa crea un sistema de actores simple con tres actores. Los mensajes se envían simultáneamente a los actores y cada actor actualiza su estado en función de los mensajes recibidos. El programa demuestra los principios del modelo de actor de aislamiento, procesamiento concurrente y paso de mensajes.

A continuación se muestra el resultado generado por el código Go proporcionado.

NOTA: Tenga en cuenta que el resultado específico puede variar según la ejecución de la rutina de ejecución. Este es un ejemplo básico y, en un escenario del mundo real, es posible que desee gestionar la terminación del actor, la supervisión y otros aspectos del modelo de actor. Bibliotecas como github.com/AsynkronIT/protoactor-go proporcionan implementaciones de modelos de actores más avanzadas en Go Lang.

Mensaje de procesamiento del actor 3: 5
Mensaje de procesamiento del actor 8: 2
Mensaje de procesamiento del actor 1: 3
Mensaje de procesamiento del actor 2: 1
Mensaje de procesamiento del actor 3: 4

Exploramos cómo PayPal manejó miles de millones de transacciones diarias utilizando solo ocho máquinas virtuales. Comenzamos con los inicios de PayPal y los desafíos que enfrentaron a medida que su popularidad se disparaba. Para hacer frente al aumento de la demanda, invirtieron en nuevo hardware, pero pronto alcanzaron un millón de transacciones por día. ¿La solución? Ampliaron la escala a más de 1000 máquinas virtuales. Sin embargo, esto trajo nuevos problemas como problemas de red y altos costos.

Luego vino el modelo actor: una forma genial de gestionar las cosas. En palabras simples, es como tener actores (o trabajadores) que hablan entre sí pasándose mensajes. Esto ayudó a PayPal a volverse súper eficiente y explicamos cómo funcionaba con actores livianos y concurrencia inteligente.

Para hacer las cosas más divertidas, incluso escribimos un código Go simple para mostrar cómo funciona el modelo de actor en acción. El código crea actores, les envía mensajes y les permite hacer lo suyo.

Finalmente una idea más clara de cómo sistemas como el de PayPal pueden manejar grandes tareas.
