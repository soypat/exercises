'''
★★☆ Es importante reconocer la diferencia entre una variable local a una función 
y una variable global.

¿Por qué el siguiente programa resulta en un error?  
¿Cómo se puede corregir?
```python
frutas = 99

def imprimir_frutas():
    print(f"Tengo", frutas, "frutas!")
    frutas = 0

imprimir_frutas()
```
'''
# La función `imprimir_frutas` tiene dos lineas.
# Cómo hemos visto, la segunda línea declara la variable local `frutas`.
# Sin embargo, hay un uso de la variable local `frutas` antes de que sea declarada
# y por ende el programa tira un error.

# Cabe destacar que mientras hay una variable global que  también se llama `frutas`,
# las variables locales toman precedencia en una función. No se puede usar 
# una variable local y global del mismo nombre en una misma función en Python.
