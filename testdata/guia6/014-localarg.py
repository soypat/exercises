'''
★★☆ Se define la variable `frutas` en el siguiente código y se la iguala a 99. 
¿Será que se puede igualar frutas a 0 desde adentro de una función?
¿Que imprime el siguiente código?
```python
frutas = 99

def descartar_frutas():
    frutas = 0

igualar_a_cero(frutas)
print(frutas)
```
'''
# Las funciones son un entorno semi-cerrado. Las asignaciones dentro de la función
# crean una nueva variable! La variable local `frutas` adentro de la función
# es una variable nueva distinta a la variable global `frutas` afuera de la función!
