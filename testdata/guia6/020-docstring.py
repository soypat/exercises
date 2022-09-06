'''
★☆☆ A continuación se tiene una función con un **docstring**. Los docstrings ayudan
a documentar lo que hace una función y muestran cuando se llama help.
```python
def suma(a, b):
    """Devuelve la suma de a y b."""
    return a + b
```
Reescriba la función asi el docstring ocupa multiples líneas según [PEP-257](https://peps.python.org/pep-0257/).
'''
# Dos resultados válidos:
def suma(a, b):
    """Devuelve la suma de a y b.
    """
    return a + b

def suma(a, b):
    """
    Devuelve la suma de a y b.
    """
    return a + b

# E incluso se pueden usar comillas simples
def suma(a, b):
    '''
    Devuelve la suma de a y b.
    '''
    return a + b

print(help(suma))