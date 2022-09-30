'''
★☆☆ Defina una función `expv` que tome dos argumentos:

* El exponente como argumento mandatorio.
* La base como argumento opcional. Por defecto es el número [*e*](https://en.wikipedia.org/wiki/E_(mathematical_constant)).

El resultado de la función debería ser el resultado de la base elevado al exponente (float).
'''
def expv(exponente, base=2.718281828459045):
    return base**exponente