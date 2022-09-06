'''
★☆☆ Defina una función `esta_presente` que reciba una lista de cadenas y que
devuelva un bool. La función deberá devolver verdadero si alguno de los siguientes
nombres está en la lista:
```python
invitados = ["Cristina", "Mauricio", "Alfonsin", "Alberto"]
```
Utilice la variable `invitados` como variable global.
'''
invitados = ["Cristina", "Mauricio", "Alfonsin", "Alberto"]

def esta_presente(lista):
    for invitado in invitados:
        if invitado in lista:
            return True
    return False
