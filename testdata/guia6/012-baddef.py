'''
★☆☆ Dado el siguiente programa modifíquelo para que imprima el resultado de la función `simple`.
Explique porqué razón falla el programa tal como se presenta.
```python
resultado = simple()
print(resultado)

def simple():
    return "buenisimo"
```
'''
# La definición de la función debe estar antes de su uso.
def simple():
    return "buenisimo"

resultado = simple()
print(resultado)