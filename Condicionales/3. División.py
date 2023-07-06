print("Bienvenido! vamos a dividir")
dividendo = float(input("Escriba el dividendo ---> "))
divisor = float(input("Escriba el divisor ---> "))

if divisor != 0:
  print("La división es: " + str(round (dividendo / divisor, 2)))
else:
  print("¡Error!, no se puede dividir entre cero.")