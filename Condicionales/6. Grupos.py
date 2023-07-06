print("Te separaremos en grupo A y B, segun tu nombre y sexo.")
name = input("Escribe tu nombre ")
sexo = input("Escribe la letra ---> M(Masculino) o F(Femenino) ")

if (name.lower() <= "m" and sexo == "F") or (name.lower() >= "n" and sexo == "M"):
  print("Perteneces al grupo A")
else:
  print("Perteneces al grupo B")