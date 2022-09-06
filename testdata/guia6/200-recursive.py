'''
★★★★ ¿Qué imprime el siguiente programa?
```python
def anidar_parentesis(cadena, n):
    if n <= 0: # Condición de finalización.
        return cadena 
    cadena = "(" + cadena
    cadena = anidar_parentesis(cadena, n-1)
    return cadena + ")"

c = anidar_parentesis("hola", 3)
print(c)
```
'''
# Imprime "(((hola)))"
# Una función que se llama a si misma se le dice "recursiva".
# Además de tener utilidad limitada en Python son dificiles
# de entender al momento de querer seguirles el hilo de ejecución.
# Este ejercicio está aquí solo para cosquillar la curiosidad.