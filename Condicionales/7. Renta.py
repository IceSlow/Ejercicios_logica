print("Calularemos su Tipo Impositivo según su renta anual")
renta_user = int(input("Escriba su renta anual ---> "))

if renta_user < 10000:
  print("Su tipo impositivo es del 5%")
elif 10000 <= renta_user < 20000:
  print("Su tipo impositivo es del 15%")
elif 20000 <= renta_user < 35000:
  print("Su tipo impositivo es del 20%")
elif 35000 <= renta_user < 60000:
  print("Su tipo impositivo es del 30%")
else:
  print("Su tipo impositivo es del 45%")