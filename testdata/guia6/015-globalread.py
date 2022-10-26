'''
★☆☆ Del ejercicio anterior sabemos que podemos crear variables nuevas adentro de
una función que comparten el mismo nombre con variables globales afuera de la función.

A pesar que las funciones tienen un espacio de trabajo aparte para variables
locales, no todo es tan "local". ¿Qué imprime el siguiente código? ¿Es válido?
```python
frutas = 99

def imprimir_frutas():
    print(f"Tengo {frutas} frutas!")

imprimir_frutas()
```
'''
# El código es valido e imprime 99.
# Las funciones pueden acceder a variables globales.