import requests

url = "http://localhost:5000"

###############INSERINDO CARRO###############
car = {"name":"Unão da massa","price":15000.00,"year":2009}

r = requests.post(url+"/cars", json=car)

print(r.text)
##############################################

###################BUSCANDO CARROS#######################
car = {"name":"Ferrari","price":1500.00,"year":2003}
r = requests.post(url+"/cars", json=car) #INSERINDO UM SEGUNDO CARRO

r = requests.get(url+"/cars")

print(r.text)
#########################################################

##################BUSCANDO O PRIMEIRO CARRO#################
r = requests.get(url+"/cars/1")

print(r.text)
#############################################################


#####################EDITANDO O SEGUNDO CARRO########################
car = {"name":"Ferrari EDITADA","price":1500.00,"year":2003}

r = requests.put(url+"/cars/2", json=car)

print(r.text)
#####################################################################

#########################DELETANDO O SEGUNDO CARRO#########################
r = requests.delete(url+"/cars/2")

print(r.text)
############################################################################
