# MANUAL TECNICO

## Introducción

Este manual proporciona una descripción técnica del intérprete desarrollado en Go que utiliza ANTLR para realizar el análisis léxico y sintáctico de un lenguaje de programación personalizado. El intérprete se divide en varios módulos, cada uno con un propósito específico.

## Módulos del Proyecto

### 1. block

El módulo `block` se encarga de construir el bloque de código que contiene todos los statements del programa. Esto incluye la estructura de control de flujo, como bucles y condicionales, así como las declaraciones de variables y funciones.

### 2. environment

El módulo `environment` se encarga de la construcción y gestión de la tabla de símbolos y funciones. La tabla de símbolos almacena información sobre las variables y sus valores, mientras que las funciones se registran y gestionan para su posterior llamado y ejecución.

### 3. expressions

El módulo `expressions` maneja todas las expresiones definidas en la gramática del lenguaje. Esto incluye expresiones aritméticas, de cadena, lógicas y cualquier otra operación que devuelva un valor.

### 4. instructions

El módulo `instructions` contiene todas las instrucciones y statements definidos en la gramática. Esto incluye la implementación de la impresión (`print`), asignaciones, llamadas a funciones y otras operaciones de alto nivel.

### 5. operators

El módulo `operators` se encarga de implementar todas las operaciones definidas en el lenguaje, ya sean aritméticas, relacionales o lógicas. Esto permite que el intérprete realice cálculos y evaluaciones de expresiones de manera precisa.

### 6. type

El módulo `type` define un enum con los tipos de datos primitivos disponibles en el lenguaje. Estos tipos son fundamentales para la gestión de variables y operaciones en el intérprete.

### 7. value

El módulo `value` contiene los structs que representan los diferentes tipos de valores manejados por el intérprete. Esto incluye estructuras de datos para representar valores, errores, funciones y parámetros.

### 8. visitor

El módulo `visitor` implementa la funcionalidad del patrón Visitor. Esto facilita la traversión y evaluación de la estructura del código fuente durante el proceso de interpretación.

## Requisitos Previos

Antes de utilizar este intérprete, se debe de tener instalado en el sistema:

1. Go (Golang).
2. ANTLR 4 instalado y configurado para generar el código Go correspondiente.