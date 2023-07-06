print("Bienvenido! validaremos si califica para tributar impuestos")
edad_usuario = int(input("Escriba su edad ---> "))
ingresos_usuario = int(input("Escriba sus ingresos mensuales ---> "))

if edad_usuario >= 16 and ingresos_usuario >= 1000:
  print("Califica para tributar impuestos")
else:
  print("No califica para tributar impuestos")