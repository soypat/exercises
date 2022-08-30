"""
Defina una función `division_entera` que tome dos enteros como argumentos y que 
devuelva **dos** enteros: el resultado de la division entera y el resultado del módulo
del primer número dividido por el segundo número.
"""
def division_entera(divisor, dividendo):
    cociente = divisor // dividendo
    resta = divisor % dividendo
    return cociente, resta
