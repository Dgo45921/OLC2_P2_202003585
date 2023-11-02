print("--------------------------");
print("---FUNCIONES EMBEBIDAS----");
print("----------14 pts----------");
print("--------------------------");

print("");
func suma(_ numero1: Int, _ numero2: Int) -> Int {
    var resultado = numero1 + numero2
    return resultado
}

var resultado = suma(5, 3)
print("La suma es: ", resultado)

print("");
func saludo3() {
    print("saludos!")
}

func saludo2() {
    print("mundo")
    saludo3()
}

func saludo1() {
    print("hola")
    saludo2()
}

saludo1()


print("");
func duplicar(_ x: inout Int){
    x = x + 1
}

var numero1 = 1
duplicar(&numero1)
print("numero2:", numero1)

print("");
var num3 = 9.5
var num4 = 3.6
var num1 = Int("20")
var num2 = Int("20")
var suma = String(num1 + num2)
var resta = String(num3 - num4)
var cadena = String(true) + "->" + String(3.504)

print("valor true: ", String(true))
print("valor false: ", String(false))
print("valor 1: ", suma)
print("valor 2: ", resta)
print("valor cadena: ", cadena);

/*
--------------------------
---FUNCIONES EMBEBIDAS----
----------14 pts----------
--------------------------

La suma es:  8

hola
mundo
saludos!

251
85

numero2: 2

valor true:  true
valor false:  false
valor 1:  40
valor 2:  5.9
valor cadena:  true->3.504
*/