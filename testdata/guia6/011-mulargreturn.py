'''
★☆☆  Defina una función `division_entera` que tome dos números enteros como argumentos y que 
devuelva **dos** números enteros: el primer número entero devuelto debe ser el resultado de la division entera 
y el segundo número devuelto es el resto (módulo) de la división.
'''
def division_entera(divisor, dividendo):
    cociente = divisor // dividendo
    resta = divisor % dividendo
    return cociente, resta
