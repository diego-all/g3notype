**REQUESTS**


**CREATE**

curl --location 'http://localhost:9090/products' \
--header 'Content-Type: application/json' \
--data '{
"name": "Teléfono móvil", "description": "Smartphone de última generación", "price": 799, "quantity": 10}'



**READ**

curl --location 'http://localhost:9090/products/get/6'



**UPDATE**  

curl --location --request PUT 'http://localhost:9090/products/update/6' \
--header 'Content-Type: application/json' \
--data '{
"name": "Camiseta", "description": "Camiseta de algodón", "price": 20, "quantity": 50}'



**LIST**

curl --location 'http://localhost:9090/products/all'




**DELETE**

curl --location --request DELETE 'http://localhost:9090/products/delete/6'