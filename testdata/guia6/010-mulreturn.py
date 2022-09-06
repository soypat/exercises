'''
★☆☆ Defina una función `susurro_grito` que tome como argumento una cadena y que 
devuelva dos cadenas: La cadena en minúsculas y la misma cadena en mayúsculas.
'''
def susurro_grito(se_dice):
    susurro = se_dice.tolower()
    grito = se_dice.toupper()
    return susurro, grito