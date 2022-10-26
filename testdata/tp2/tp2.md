# Análisis bursátil

## Introducción
Análisis del mercado de valores significa analizar las tendencias actuales e históricas en el mercado de valores para tomar decisiones futuras de compra y venta. El análisis del mercado de valores es cada vez más frecuente en el mundo de los negocios digitales. 

## Motivación
En este TP se les va pedir que trabajen con funciones de lectura/escritura a archivo. Algunos puntos de interés a aprender o reforzar en el transcurso de este TP son los siguientes:

 * **Modularización.** Las funciones son un concepto poderoso. Pueden ahorrar mucho laburo si son definidas adecuadamente.

* **Extensión de funciones.** Si uno necesita crear una nueva función que haría algo parecido a otra función ¿Cómo reconcilian las dos funciones? ¿Copian y pegan el código? O ven cómo modularizar el programa y descomponer la función existente en más funciones.

* **Memoria persistente.** Cómo bien ya saben, los archivos son la memoria persistente de una computadora. No se borran una vez terminado el programa. Es muy importante tener un buen manejo de escritura y lectura a archivo para poder aprovecharlo.

* **Gráficos.** La habilidad de presentar datos en una forma clara y convincente es invaluable. Su utilidad va más allá del mundo de negocios y forma una parte central en la argumentación. Este TP no les va enseñar esto, pero puede considerarse un ejercicio para ese fin.

## Datos de entrada
La cátedra les entregará un archivo .csv con el precio de cotización de las acciones de distintas empresas. Cada fila del archivo representa una fecha y cada columna representa el precio de la acción en esa fecha.

```csv
Date, SATL, MELI, MHV, …
2021-10-04, 9.93, 9.88, 10.12 …
2021-10-05, 9.92, 9.89, 10.00 …
2021-10-06, 9.90, 9.90, 10.92 …
2021-10-07, 9.91, 9.88, 11.23 …
```

## Sugerencias antes de comenzar
* Para leer el archivo .csv se sugiere que utilicen la función [read_csv()](https://pythonbasics.org/read-csv-with-pandas/), 
de la librería pandas, de la siguiente forma:

    ```python
    import pandas as pd
    f = open("bolsa.csv")
    data = pd.read_csv(f)
    print(data.head())
    ```

* Crear una función que a partir del nombre de la acción se obtengan los datos de la misma, i.e:
    ```python
    fechas, precio = get_stock("MELI")
    ```

* Definir funciones gráficas para uso interno que devuelvan un objeto matplotlib.plot. Al objeto se le pueden agregar más gráficos o mostrarlo en ventana o guardarlo en un archivo. Permite mayor flexibilidad programática y es ideal para trabajar con funciones.

> Recuerde que el contenido de bolsa.csv es demostrativo! Sus programas serán probados con archivos con distintas fechas y acciones!

## Consigna

1. Haga una función `plot_price` que reciba el nombre de una acción y genere un gráfico de líneas del precio, fecha por fecha, de dicha acción.
    ```python
    plot_price("MELI")
    ```
    La función también debe poder recibir los parámetros opcionales start y end representando las fechas que se quiere graficar.
    ```python
    from datetime import date
    plot_price("MELI", start=date(2022,1,1), end=date(2022,6,1))
    ```

2. Haga una función `average_monthly_bar_plot` que genere un gráfico de barras con el precio promedio de la acción, mes a mes.
    ```python
    average_monthly_bar_plot("SATL")
    ```

3. Mirando el precio de hoy, es fácil determinar cuándo hubiera sido el día ideal, en el pasado, para comprar una acción y generar la mayor ganancia vendiéndola hoy.

    Escriba una función max_gain que reciba el nombre de una acción y una fecha de venta. La función debe buscar la fecha de compra que hubiera generado la mayor ganancia. La función debe devolver esta fecha y el retorno de la inversión.

    Se define retorno de la inversión como: (Precio de venta - precio de compra) / precio de compra

    ```python
    fecha_compra, retorno = max_gains("RTX", fecha_venta)
    ```

4. Podríamos, también, realizar el análisis del punto anterior, para todas las acciones del archivo y asi determinar qué acción me hubiera convenido comprar.

    Crear una función que reciba una fecha de venta y genere un gráfico de barra con el retorno de la mejor inversión para TODAS las acciones. Es decir: para cada acción, debe haber una barra en el gráfico que represente el retorno que hubiera generado la mejor inversión.

    Se debe realizar esta función utilizando la función del item anterior.

    ```python
    plot_all_returns(fecha_venta)
    ```

5. Escriba una función save_report que cree un informe financiero en formato markdown. El informe se debe llamar "report.md".
    ```python
    save_report()
    ```
    El informe deberá tener las siguientes secciones:
    
    1. **Resumen:**
        1. Información sobre los datos y gráficos que se muestran en el informe.
        2. Fecha inicial y fecha final del análisis bursátil según el archivo `bolsa.csv`
        3. Nombre de las acciones que se estudian en el informe
    2. **Máximos rendimientos:**
        1. El máximo rendimiento posible de todas las acciones, considerando la última fecha del archivo como la fecha de venta. Estos rendimientos deben ser representados en formato de tabla o lista junto con las fechas de compra y venta. También, se debe señalar en negrita la acción con mayor retorno
        2. Un gráfico con el máximo rendimiento posible de todas las acciones. También utilizando la última fecha del archivo como fecha de venta.
    3. **Mejor inversión:**
        1. Un gráfico de precios de la acción de mayor retorno donde también se marquen las fechas de compra y venta para dicho retorno.
    4. (Opcional): **Peor inversión:**
        1. Un gráfico del precio promedio mes a mes de la acción con peor retorno. Señalizar esta acción con itálica en la tabla de retornos

6. Subir TP al campus.

## Anexo
### Tutorial Markdown VSCode
1. Cree un archivo llamado **example.md** en VSCode. Este archivo va a definir nuestro documento escrito con markdown.

2. Guarde una imagen en el mismo directorio que el archivo markdown. Puede ser jpg tambien.

3. Abra el archivo example.md y copie el siguiente texto a el
    ```markdown
    # Titulo
    Esto es texto comun. **Texto en negrita**.
 
    Esto es un [link](https://www.example.com).
 
    Esto es una imágen:
 
    ![](imagen.png)
    ```

4. Reemplace el nombre de la imagen del texto **imagen.png** con el nombre de la image guardada en el paso 2.

5. Haga click en el botón de visualización en la esquina superior derecha. Es un icono
que tiene una lupa en frente de dos cuadrados.

6. Se debería haber abierto un nuevo tab. A medida que edites el contenido de example.md los cambios se verán reflejados en el tab Preview example.md (ver imagen abajo). Este tab es el resultado del documento markdown. El [gobierno de argentina provee una guía simple](https://www.argentina.gob.ar/contenidosdigitales/markdown) para empezar a escribir markdown como un pro.

Si lo desean, pueden convertir un archivo markdown a pdf: [https://www.markdowntopdf.com/](Si lo desean, pueden convertir un archivo markdown a pdf: https://www.markdowntopdf.com/ )