'''
★★☆ En el siguiente programa la función `igualar_a_cero` recibe un número y lo iguala
a cero. Luego se invoca dicha función con la variable `frutas` que vale 99.

Esto significa que corre la función y adentro de `igualar_a_cero` se iguala el
argumento (en este caso, `frutas`) a cero. Por qué entonces el valor de `frutas` no es 0
una vez terminado el programa?
```python
def igualar_a_cero(numero):
    numero = 0

frutas = 99
igualar_a_cero(frutas)
print(frutas)
```
'''
# Las funciones reciben una copia del argumento, no el argumento en si.
# Se está igualando la copia de `frutas` a cero mientras que el valor
# de `frutas` queda igual.
# Bonus: El siguiente código hace algo diferente al del problema?
def igualar_a_cero(frutas):
    frutas = 0

frutas = 99
igualar_a_cero(frutas)
print(frutas)
