# {{.LowerEntity}}s-API


## REST API



### Start the project

**Build the API**

    go build ./cmd/api
    go build -o {{.LowerEntity}}API ./cmd/api/

    go run ./cmd/api



**Create sqlite database**

If you are using the Ubuntu Linux operating system, you can start the SQLite database by running the following command:

    sh script.sh


**Execute the API**

    ./api
    ./{{.LowerEntity}}API


**cURL requests**

You can test the generated API using the cURL request in requests.txt file.


