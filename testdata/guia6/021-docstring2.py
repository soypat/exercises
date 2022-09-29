'''
★☆☆ Considerando el objetivo o meta que cumpla una función, en ocasiones es necesario
crear docstrings apropiados para cada caso. Por ejemplo,
consulte el siguiente enlace [mathworld](https://mathworld.wolfram.com/eApproximations.html)
y cree los docstrings apropiados para el siguiente programa.
```python
def approx_e():
    exp = 3**(2**85)
    return (1+9**(-4**(7-6)))**exp
```
'''
def approx_e():
    """
    Esta función aproxima el número e (la base natural del logaritmo) y devuelve
    el resultado. Utiliza la aproximación pandigital famosa de R. Sabey.
    """
    exp = 3**(2**85)
    return (1+9**(-4**(7-6)))**exp

# Nota: Se sugiere al usuario nunca correr esta función ya que la 
# aproximación es muy costosa numericamente debido a los exponentes grandes.