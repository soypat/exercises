'''
★★★ *Divide y vencerás*. Llegó el fin de los tiempos. La sociedad ha colapsado y lo único que ha quedado
a tu nombre es una lista de todos los temas que has escuchado en spotify. Planeas
pasar el resto de tus días buscando vinílos de tus temas preferidos entre los escombros
para luego escucharlos una vez llegada la noche para sentirte acompañada/o.  

Sin embargo, la lista de temas es extremadamente larga.  

Sería imposible recorrerla a ojo en un tiempo razonable. Por suerte sabes Python
por cursar IPC en esa universidad... cómo era el nombre? UdeSA... o algo asi.  

Querés empezar buscando los temas que más escuchaste durante tu vida, pero no querés
buscar temas de corta duración. Mucho más te va rendir un tema largo y tendido que
uno de unos pocos segundos.

#### Entrada
La lista tiene un tema por línea. A continuación un ejemplo:
```plaintext
Big Iron - Marty Robbins {4096,3:56}
Air Supply - Making Love Out Of Nothing At All {512,5:42}
Every Time the Sun Comes Up - Sharon Van Etten {1024,4:23}
Don't Look Up - Nicholas Britell {128,52} {456,4:08}
```
Cada línea para cada tema sigue un formato:
```plaintext
NOMBRE_DE_TEMA - CANTAUTOR/A {INFO(ORIGINAL)} {INFO(BONUS)} 
```
donde `INFO` tiene la forma `REPRODUCIONES,DURACION`. `DURACION` puede ser la cantidad
de segundos que dura la cancion, o la duración en `minutos:segundos`, o en raros casos
horas:minutos:segundos. La información BONUS puede o no estar presente.

#### Salida esperada
Imprimir el nombre y cantautor/a de los 3 temas más escuchados con duración superior
a 2:31 minutos/segundos (no inclusive). La duración de un tema es dada por la más 
alta duración entre su version original y bonus. La cantidad de reproducciones
es dada por la suma de sus dos versiones.

Para el ejemplo dado arriba el programa debería imprimir
```plaintext
Big Iron - Marty Robbins
Every Time the Sun Comes Up - Sharon Van Etten
Don't Look Up - Nicholas Britell
```

#### Sugerencias
Definir una función `duracion_a_tiempo` que convierta una cadena en
formato duración a segundos. Ejemplo: 
```python
> duracion_a_tiempo("2:31")
> 151
```

Definir una función `info` que a partir del texto informativo de
un tema devuelva la cantidad de reproducciones y la duración en segundos. Ejemplo
```python
> rep, seg = info("{2332,11:32:09}")
> print(rep, seg)
> 2332 41529
```

**Puede ser muy útil definir una función `procesar_tema` que procese la cadena de
un solo tema y validar que funcione bien antes de resolver el problema entero.**
'''
def duracion_a_tiempo(cadena_duracion):
    """
    Convierte texto en formato duración a segundos.
    Puede tener cualquiera de las siguientes formas:
    segundos,  minutos:segundos,  horas:minutos:segundos
    """
    hms = cadena_duracion.split(":")
    if len(hms) == 1:
        segundos = int(hms[0])
    elif len(hms) == 2:
        segundos = int(hms[0])*60 + int(hms[1])
    elif len(hms) == 3:
        segundos = int(hms[0])*60*60 + int(hms[1])*60 + int(hms[2])
    return segundos

def info(cadena_info:str):
    """
    Procesa texto informativo de un tema y devuelve
    cantidad de veces que se escucho un tema y segundos que dura un tema.
    cadena_info tiene el formato: "{NUMERO,DURACION}"
    """
    cadena_info = cadena_info.removeprefix("{")
    cadena_info = cadena_info.removesuffix("}")
    sep = cadena_info.split(",")
    return int(sep[0]), duracion_a_tiempo(sep[1])

def procesar_tema(tema):
    """Procesa un tema de la forma 
    NOMBRE_DE_TEMA_Y_CANTAUTOR/A {INFO(ORIGINAL)} {INFO(BONUS)}
    """
    bonus_reproducciones, bonus_duracion = 0, 0
    # En base a la posición de las llaves podemos slicear la cadena
    # para dividir el texto en partes de info y título+cantautor/a
    info_original_indice = tema.index("{")
    info_original_indice_fin = tema.index("}")
    titulo = tema[:info_original_indice-1]
    # Obtenemos datos de informacion del tema original
    orig_reproducciones, orig_duracion = info(tema[info_original_indice:info_original_indice_fin])
    
    # Investigamos si hay informacion bonus luego de info original.
    bonus_texto = tema[info_original_indice_fin+1:]
    if "{" in bonus_texto:
        info_bonus_indice = bonus_texto.index("{")
        info_bonus_indice_fin = bonus_texto.index("}")
        bonus_reproducciones, bonus_duracion = info(bonus_texto[info_bonus_indice:info_bonus_indice_fin])
    return titulo, orig_reproducciones, orig_duracion, bonus_reproducciones, bonus_duracion 

temas = input().split("\n")

duracion_minima = duracion_a_tiempo("2:31")
informaciones = []
for tema in temas:
    titulo, orig_rep, orig_dur, bon_rep, bon_dur = procesar_tema(tema)
    # Nos importa solo la más alta de las dos duraciones.
    max_duracion = max(bon_dur, orig_dur)
    # Descartamos el tema si es más corto que el limite.
    if max_duracion < duracion_minima:
        continue
    
    # Acá usamos un trucazo. Si guardamos una tupla en una lista y luego
    # la ordenamos usando sort, sort ordenará según el primer valor de la tupla.
    info_tupla = (orig_rep+bon_rep, titulo)
    informaciones.append(info_tupla)

informaciones.sort(reverse=True)
print(informaciones[0][1])
print(informaciones[1][1])
print(informaciones[2][1])