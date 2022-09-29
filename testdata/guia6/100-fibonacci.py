'''
★★☆ Defina una función llamada `lista_fibonacci` que tome un número entero `n` como
argumento y devuelva una lista que contenga los [números de Fibonacci](https://es.wikipedia.org/wiki/Sucesi%C3%B3n_de_Fibonacci)
desde `0` hasta `n` inclusive.
'''
# Este problema no tiene un contador, es solo con acumuladores!
def lista_fibonacci(n):
    """
    Devuelve los números de fibonacci de 0 hasta n en una lista.
    """
    fibonacci_anterior = 0
    fibonacci_actual = 1
    lista = []
    while fibonacci_anterior <= n: 
        lista.append(fibonacci_anterior)
        nuevo_fibonacci = fibonacci_anterior + fibonacci_actual
        fibonacci_anterior = fibonacci_actual
        fibonacci_actual = nuevo_fibonacci
    return lista