'''
★★☆ Defina una función `local_avg` que tome los siguientes argumentos:
* `numbers`: Una lista de números
* `index`: Un índice que no debe superar la longitud de la lista
* `window` *(Opcional)*: Cantidad **impar** de elementos a tomar en el promedio. Por defecto 3.  

La función debería calcular la [media móvil central](https://es.wikipedia.org/wiki/Media_m%C3%B3vil) de la lista de números
en la ventana (window) alrededor del indice. Idealmente se toman `window` cantidad de elementos para el cálculo.
En el caso en donde la ventana excede los límites de la lista se debe
calcular el promedio correctamente dividiendo la suma de los números adentro de la ventana
por la cantidad de números adentro de la ventana.


La salida de la función debería ser:
* La media movil en el índice. Si hay un argumento inválido debería devolver 0 (float).
* Una cadena que indica si algún argumento era inválido. Si no hay argumentos inválidos se devuelve una cadena vacía.  

Ejemplo:  
```python
# Se promedian los 3 valores centrados en el índice 3:
> local_avg([1, 2, 3, 4, 3], 2) # == (2+3+4)/3
> (3.0, '')
# Se promedian los 3 valores centrados en el índice 0.
# Como el índice -1 cae afuera de la lista no se incluye en el promedio:
> local_avg([1, 2, 3, 4, 3], 0) # == (1+2)/2
> (1.5, '')
# Se promedian los 5 valores centrados en 5, que incluye toda la lista:
> local_avg([1, 2, 3, 4, 3], 2, window=5) # ==(1+2+3+4+3)/5
> (2.6, '')
'''
def local_avg(numbers, index, window=3):
    if window % 2 != 1:
        return 0, "window necesita ser impar"
    elif index >= len(numbers) or index < 0:
        return 0, "índice inválido"
        
    distance = (window-1) // 2
    # Calculamos el trajecto recorrido en la lista y lo limitamos a los extremos.
    minimo = max(0, index-distance)
    maximo = min(len(numbers)-1, index+distance)+1
    # Calculamos la ventana real que puede diferir del argumento cuando el
    # indice esta cerca de los extremos de la lista.
    window = maximo-minimo 
    return sum(numbers[minimo:maximo]) / window, ""