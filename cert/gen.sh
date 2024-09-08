rm *.key *.crt

openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650 -subj "/C=CN/ST=Beijing/L=Beijing/O=GameSale/OU=Personal/CN=localhost" -addext "subjectAltName = DNS:localhost"