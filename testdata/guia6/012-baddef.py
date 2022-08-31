"""
Dado el siguiente programa corrijalo para que imprima el resultado de la función `simple`.
Elabore por qué razón falla el programa.
```python
resultado = simple()
print(resultado)

def simple():
    return "buenisimo"
```
"""
# La definición de la función debe estar antes de su uso.
def simple():
    return "buenisimo"

resultado = simple()
print(resultado)