
# 1. Create a Root CA (if you don’t already have one)

## Generate Root CA private key
openssl genrsa -out tyllaca.key 4096

## Create self-signed Root CA certificate
openssl req -x509 -new -nodes -key tyllaca.key -sha256 -days 3650 -out tyllaca.crt


## 
openssl req -new -key server.key -out server.csr

openssl req -new -key server.key -out server.csr

openssl x509 -req -in server.csr -CA tyllaca.crt -CAkey tyllaca.key -CAcreateserial -out server.crt -days 825 -sha256

openssl verify -CAfile tyllaca.crt server.crt

## With config

openssl req -new -config transition.conf -keyout trans-server.key.pem -out trans-server.csr.pem

openssl x509 -req -in trans-server.csr.pem -CA tyllaca.crt -CAkey tyllaca.key -CAcreateserial -out trans-server.crt.pem -days 825 -sha256


--
A challenge password:1234
An optional company name:qwerty