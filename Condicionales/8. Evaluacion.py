print("Programa para saber su nivel y puntuación del empleado")
puntuación_empleado = float(input("Escriba la puntuación del empleado "))
dinero = puntuación_empleado * 2400
valor = True

if puntuación_empleado == 0:
  nivel = "Inaceptable"
elif puntuación_empleado == 0.4:
  nivel = "Aceptable"
elif puntuación_empleado >= 0.6:
  nivel = "Meritorio"
else:
  valor = False
  print("Valor Incorrecto")

if valor == True:
  print("Su nivel respecto a su puntuación es: " + nivel )
  print("El dinero obtenido es: " + str(dinero))