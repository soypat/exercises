'''
★★★★ ¿Que imprime el siguiente programa?
```python
def hola(uno, dos):
    print("hola",uno, dos)

def soy(cadena):
    print("devolviendo soy")
    return "soy " + cadena

def fulano():
    print("devolviendo fulano")
    return "fulano"

def mirame():
    print("devolviendo mirame")
    return "mirame"

hola(soy(fulano()), mirame())
```
'''
# Imprime lo siguiente:
"""
devolviendo fulano
devolviendo soy
devolviendo mirame
hola soy fulano mirame
"""
# Esto se debe al ordén de evaluación de los argumentos en los lenguajes
# imperativos como Python. Estos siguen una regla simple: 
# Cuando se invoca una función siempre primero se evalúan
# los argumentos de la misma, de izquierda a derecha y finalmente se invoca la función. 
# Por ende, cuando se invoque `hola` primero se va intentar evaluar `soy`, que a su vez
# es una función que también tiene un argumento, entonces no podemos evaluarla sin 
# primero evaluar su primer argumento `fulano`. `fulano` no tiene argumentos y
# por lo tanto es la primera expresión en evaluarse en este ejemplo. Luego de 
# evaluar `fulano` podemos evaluar `soy` y seguir al segundo argumento de `hola`
# e invocar `mirame`, finalmente invocando `hola`.
#
# Cabe destacar que en el día a día no suele importar el orden de evaluación
# de los argumentos de una función, este ejercicio está aquí principalmente
# para cosquillar la curiosidad. Si llega servir para algo este ejercicio que sea
# como una buena contraindicación del uso de `print` adentro de una función.