password_user = "jhampier23edu"
password = ""

while password_user != password:
    password = input("Escriba la contraseña ")
    if password == password_user:
        print("Contraseña correcta")
    else:
        print("Contraseña incorrecta! Intente de nuevo")