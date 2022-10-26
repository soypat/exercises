'''
★☆☆ Defina una función llamada `greet_user` que tome como argumento una cadena y que imprima 
`Bienvenida/o de vuelta, ` seguido por la cadena. Si la función no recibe ningún argumento
debería imprimir `Bienvenida/o de vuelta, Usuaria/o`.
Ejemplo:
```python
> greet_user("Esteban")
> Bienvenida/o de vuelta, Esteban
> greet_user()
> Bienvenida/o de vuelta, Usuaria/o
```
'''
def greet_user(username="Usuaria/o"):
    print(f"Bienvenida/o de vuelta, {username}")