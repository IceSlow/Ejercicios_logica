print("Programa para calcular cuanto costara la entrada ")
edad_user = int(input("Escriba su edad ---> "))

if edad_user < 4:
  entrada = 0
elif 4 <= edad_user < 18:
  entrada = 5
else:
  entrada = 10

print("Su entrada será de ---> " + str(entrada) + "$")