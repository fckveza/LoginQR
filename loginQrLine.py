import requests, json, time

hostVH = "https://api.vhtear.com/"
apikey = "PREMIUMKEY" #Chat me on whatsapp https://wa.me/+628123552767

#Header sesuai'in dengan scripmu
headerLine = {
  "User-Agent": "Line/6.3.2",
  "X-Line-Application": "DESKTOPWIN\t6.3.2\tVH-PC\t10.0.14;SECONDARY",
  "x-lal": "en_id"
}

def NewLogin():
  sagne = requests.post(hostVH+"zloginqr="+apikey, json=headerLine)
  if sagne.status_code == 200:
     tH = sagne.json()
     print("Link Login: "+tH["result"]["url"]) #Atau ganti dengan client.sendMessage(to,tH["result"]["url"])
  time.sleep(4)
  ve = requests.get(hostVH+"zgetcode="+apikey)
  if ve.status_code == 200:
    bG = ve.json()
    if bG["result"]["code"] != "":
        print("Pin code : ",bG["result"]["code"])
  time.sleep(4)
  ve = requests.get(hostVH+"zgettoken="+apikey)
  if ve.status_code == 200:
      rH = ve.json()
      if rH["result"]["token"] != "":
         print("Token : ",rH["result"]["token"]) #Atau ganti dengan client.sendMessage(to,rH["result"]["token"])
         print("Cert : ",rH["result"]["certificate"]) #Atau ganti dengan client.sendMessage(to,rH["result"]["certificate"])
         
  
def LoginWithCertificate(cert):
  sagne = requests.post(hostVH+"zloginqrwithcert="+apikey+"&"+cert, json=headerLine)
  if sagne.status_code == 200:
     ve = sagne.json()
     print("Link Login: "+ve["result"]["url"]) #Atau ganti dengan client.sendMessage(to,ve["result"]["url"])
  else:
     print("Error : ", sagne.status_code)
  time.sleep(4)
  ve = requests.get(hostVH+"zgettoken="+apikey)
  if ve.status_code == 200:
      cok = ve.json()
      if cok["result"]["token"] != "":
        print("Token : ",cok["result"]["token"]) #Atau ganti dengan client.sendMessage(to,cok["result"]["token"])
        
  
#LoginWithCertificate("certificate mu")
#NewLogin()
