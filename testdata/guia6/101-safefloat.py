'''
★★★ Examine la función `convertir_a_float` y el uso de la misma.
¿Qué **debería** imprimir el programa cuando es corrido?
¿Qué es lo que **realmente** imprime el programa?
¿Cómo mejoraría la función `convertir_a_float`?
```python
def convertir_a_float(cadena):
    separado = cadena.split(".")
    hay_punto = len(separado) == 2
    es_real = hay_punto and separado[0].isnumeric() and separado[1].isnumeric()
    if es_real:
        return float(cadena)
    else:
        return 0.0

a = convertir_a_float("2.5")
b = convertir_a_float("2")
c = convertir_a_float("1.")
d = convertir_a_float("no soy float")
print(a+b+c+d)
```

Una vez resuelta el ejercicio: ¿En qué casos sería útil esta función?
¿Por qué no querríamos reemplazar la función `convertir_a_float` con `float`?

'''
# Con la función incorrecta imprime 2.5
# "Debería" imprimir la suma de 2.5+2+1+0., `5.5`
# dado que "no soy float" evalua a 0.0
def convertir_a_float(cadena):
    separado = cadena.split(".")
    hay_punto = len(separado) == 2
    es_real = hay_punto and separado[0].isnumeric() and separado[1].isnumeric()
    entero_punto = hay_punto and separado[0].isnumeric() and len(separado[1])==0
    es_entero = len(separado) == 1 and separado[0].isnumeric()
    if es_real or entero_punto or es_entero:
        return float(cadena)
    else:
        return 0.0

a = convertir_a_float("2.5")
b = convertir_a_float("2")
c = convertir_a_float("1.")
d = convertir_a_float("no soy float")
print(a+b+c+d)
