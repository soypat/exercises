'''
★★★ *Divide y vencerás*. "Chaos es caos en inglés" te diría Dr. Chaos,
charlando con una taza de té Chai en la mano. 
En verdad no es tan malo como su nombre lo hace aparentar...
si es que tenés un buen manejo de la semiótica.

Dr. Chaos ha descubierto un proceso para [transmogrificar](https://www.merriam-webster.com/dictionary/transmogrify#:~:text=Definition%20of%20transmogrify,with%20grotesque%20or%20humorous%20effect)
palabras. El doctor asegura que el proceso no altera cómo los humanos
perciben la palabra y que podría usarse para innumerables aplicaciones cientificas.
En este momento tu solo quieres escribir el programa que implemente
este transmogrificador asi te deja en paz de una vez y por todas.

Dr. Chaos te ha entregado un fax con el proceso del transmogrificador.
Estás leyendo el proceso cuando te deja a solas mientras va al cuarto de al lado.
Al rato comienza a sonar *"Todo un palo"* a traves de la pared.
Te dices a ti mismo "pues claro, ricotero tenía que ser" mientras
comienzas a escribir las primeras líneas del programa.

### Fax borroso del Dr. Chaos, el malevolo semiótico
```
PRIMER PASO: REDUCIR RUIDO 
1. Ruido se considera agrupaciones de caractéres inecesarias.
2. Dos o más vocales juntas son ruido. Grupos como estos deben
reducirse a meramente una vocal representativa. Se conserva la vocal
con mayor prioridad. La prioridad de las vocales en orden descendente es
Y,U,I,A,E,O

SEGUNDO PASO: ELIMINAR RUIDO
1. Reemplazar agrupaciones de caracteres ""disonantes"" por agrupaciones ""puras""
Lista de reemplazos. Los caracteres ""disonantes"" a la izquierda,
los ""puros"" a la derecha
{
    "ts":"ps",
    "mi":"ji",
    "fi":"gi",
}

TERCER PASO: ASIMILAR
1. Si y solo si la palabra tuvo 0 o 1 agrupaciones disonantes se invierte el
orden de los caracteres de la palabra EXCLUYENDO el primer y último caracter.

ULTIMO PASO: AUMENTAR EN-FA-SIS, ENFASIS, SI?
1. Hay letras que deben tener mas enfasis que otras.
Estas letras deben quedar en mayusculas en el resultado final.
2. Las vocales del castellano deben ser enfáticas. La H también.
```

**Sugerencia:** Se sugiere usar **funciones**.
'''
def transmogrificar(cadena):
    pass # La solución es dejada como ejercicio para el alumno.