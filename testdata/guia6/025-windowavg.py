'''
★★☆ Defina una función `moving_avg` que tome los siguientes argumentos:
* `numbers`: Una lista de números
* `window` *(Opcional)*: Cantidad **impar** de elementos a tomar en el promedio. Por defecto 3.

Calcule la [media móvil central](https://es.wikipedia.org/wiki/Media_m%C3%B3vil) de la lista de números.

* La media movil. Es una **lista** de los promedios de la lista números.
* Una cadena que indica si algún argumento era inválido. Si no hay argumentos inválidos se devuelve una cadena vacía. 
'''
def local_avg(numbers, index, window):      
    distance = (window-1) // 2
    # Calculamos el trajecto recorrido en la lista y lo limitamos a los extremos.
    minimo = max(0, index-distance)
    maximo = min(len(numbers)-1, index+distance)+1
    # Calculamos la ventana real que puede diferir del argumento cuando el
    # indice esta cerca de los extremos de la lista.
    window = maximo-minimo 
    return sum(numbers[minimo:maximo]) / window, ""

def moving_avg(numbers, window=3):
    if window % 2 != 1:
        return 0, "window necesita ser impar"
    avg = []
    for i,n in enumerate(numbers):
        avg.append(local_avg(numbers, i, window))
    
    return avg