# Divide and conquer
'''
★★★ *Divide y vencerás*. Miguela vive en un pueblo frutero en el valle de Oz.
Todos los días va a trabajar a la fabrica donde hay una cinta transportadora por
donde van llegando las naranjas y bananas cosechadas en el dia. Le han encargado
a Miguela que desarrolle un programa que haga los calculos para separar la fruta
según las reglas a continuación. Miguela solo ha trabajado con lenguajes de 
programación de bajo nivel como Motorola 68000 Assembly y cree que es mejor resolver
el programa en Python, y así ha venido a ti a pedirte ayuda

#### Entrada
Se espera una entrada ingresada por usuario con el siguiente formato `"011000110101F"` donde
cada `0` representa una naranja y cada `1` representa una banana. `F` siempre
esta presente al final de cada entrada. Las frutas llegan a la cinta transportadora
según el orden que especifica la entrada ingresada, de izquierda a derecha.
En el ejemplo `"0010F"` primero llegan dos naranjas
seguido de una banana y por último una naranja. Por último las frutas que entran
por la cinta transportadora pueden estar podridas. En este caso se marcan como
**Defectuosa** con la letra D después del caracter de la fruta.
En el ejemplo `"1D01D0F"` entra primero una banana, al estar seguida por la letra `D`, esa banana debe ser
descartada. La tercer fruta en entrar también es una banana que
debe ser descartada. El resultado de `"1D01D0F"` entonces es 2 bananas descartadas
y 2 naranjas OK.  

##### Salida: 
El programa debe imprimir **5** números.
Los primeros dos números son la cantidad de bananas y naranjas a entregar al distribuidor, respectivamente.
Los siguientes dos números son la cantidad de bananas y naranjas descartadas, respectivamente.
El último número son la cantidad total de fruta que se manda al orfanato de Oz.

### Manual de logistica de la distribuidora de Oz S.R.L.
```plaintext
Nro. Documento: DO-QAP-00039-ES
Titulo: MANUAL OPERACION PLANTA 1A (VERANO) -- DISTRIBUICION
Sección: PLANTA/CINTA BANANA+NARANJA
ENCARGADA/O: Miguela

Definiciones:
I. FRUTA(S): Puede referirse a una banana o naranja no defectuosa, o un conjunto de ambas.

Directivas:
1. Las primeras 5 FRUTAS son separadas para rellenar el tazón de frutas de la oficina.
2. Las siguientes 6 FRUTAS son entregadas al distribuidor.
3. Las FRUTAS restantes son reservadas para el orfanato de Oz.
```

#### Sugerencia
Miguela sugiere definir una función `procesar_frutas` que tome dos argumentos: El texto de cinta 
a procesar y la cantidad de frutas no defectuosas a procesar. Esta función debería
entonces devolver el texto de cinta sin las frutas procesadas y la cantidades de 
naranjas y bananas, OK y defectuosas.
'''
def procesar_frutas(cinta:str, numero_de_frutas:int):
    """Recibe una cadena representando la cinta y la cantidad de frutas a procesar.
    devuelve la cadena sin las primeras frutas procesadas y la cantidad de 
    bananas, naranjas buenas y descartadas como numeros enteros.
    El último caracter de cinta es ignorado siempre.
    """
    i = 0
    bananas = 0
    naranjas = 0
    bananas_desc = 0  # Descartadas
    naranjas_desc = 0
    while bananas+naranjas < numero_de_frutas and i < len(cinta)-1:
        caracter = cinta[i]
        prox_caracter = cinta[i+1]
        i += 1
        if caracter == "1" and prox_caracter != "D":
             bananas += 1
        elif caracter == "1" and prox_caracter == "D":
            bananas_desc += 1
        elif caracter == "0" and prox_caracter != "D":
            naranjas += 1
        elif caracter == "0" and prox_caracter == "D":
            naranjas_desc += 1
        
    return cinta[i:], bananas, naranjas, bananas_desc, naranjas_desc

# Aqui empieza el programa.
cinta = input()
bananas_descartadas = 0
naranjas_descartadas = 0

# Ignoramos la cantidad de frutas que van al tazón de oficina.
cinta, _, _, bd, nd = procesar_frutas(cinta, 5)
bananas_descartadas += bd
naranjas_descartadas += nd

# Frutas que van al distribuidor
cinta, banana_distribuidor, naranjas_distribuidor, bd, nd = procesar_frutas(cinta, 6)
bananas_descartadas += bd
naranjas_descartadas += nd

# Procesamos lo que queda de la cinta para que vaya al orfanato
cinta, b, n, bd, nd = procesar_frutas(cinta, len(cinta))
frutas_orfanato = b+n
bananas_descartadas += bd
naranjas_descartadas += nd

print(banana_distribuidor, naranjas_distribuidor, bananas_descartadas, naranjas_descartadas, frutas_orfanato)