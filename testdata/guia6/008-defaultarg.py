"""
Defina una función llamada `greet_user` que tome como argumento una cadena y que imprima 
`Bienvenida/o de vuelta, ` seguido por la cadena. Si la función no recibe ningún argumento
debería imprimir `Bienvenida/o de vuelta, Usuaria/o`.
Ejemplo:
```python
> greet_user("esteban")
> Bienvenida/o de vuelta, esteban
```
"""
def greet_user(username="Usuaria/o"):
    print("Bienvenida/o de vuelta, {}".format(username))