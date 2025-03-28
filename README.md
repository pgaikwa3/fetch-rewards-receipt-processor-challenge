Name: Prachi Rajesh Gaikwad
Email: pgaikwa3@binghamton.edu

Receipt Processor

Go Implementation: 
This implementation provides:
•	POST /receipts/process – To process a receipt and return a generated receipt ID.
•	GET /receipts/{id}/points – To retrieve the points for a given receipt ID.
•	In-Memory Storage – Uses a simple map to store receipt data.

Steps to Run the Project in VS Code: 
[1] Install Go
Ensure you have Go installed. Check by running:
>	go version
>	PS C:\Users\Prachi-Gaikwad\receipt-processor> go version
go version go1.24.1 windows/amd64

[2] Initialize a Go Module
In the terminal (inside the project directory), run:
>	go mod init receipt-processor
>	PS C:\Users\Prachi-Gaikwad\receipt-processor> go mod init receipt-processor
go: creating new go.mod: module receipt-processor
go: to add module requirements and sums:
        go mod tidy

[3] Install Dependencies
This installs the uuid package for generating unique receipt IDs.
>	go get github.com/google/uuid
>	PS C:\Users\Prachi-Gaikwad\receipt-processor> go get github.com/google/uuid
go: downloading github.com/google/uuid v1.6.0
go: added github.com/google/uuid v1.6.0

[4] Run the Server
Start the server with: (This will start the server on port 8080.)
>	go run main.go
>	PS C:\Users\Prachi-Gaikwad\receipt-processor> go run main.go
Server is running on port 8080...

[5] Test API Endpoints using Postman
A) Test POST /receipts/process in Postman
1. Open Postman
2. Creating a New Request
    1.	Click + New Tab in Postman.
    2.	Change the request type to POST.
    3.	Enter the URL: http://localhost:8080/receipts/process
3. Adding Headers
    •	In the “Headers” tab, add a header:
    Key: Content-Type
    Value: application/json
4. Adding the JSON Body
    •	Go to the “Body” tab.
    •	Select “raw” and set the format to “JSON”
    •	Put the below JSON: 
    {
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
        { "shortDescription": "Mountain Dew 12PK", "price": "6.49" },
        { "shortDescription": "Emils Cheese Pizza", "price": "12.25" },
        { "shortDescription": "Knorr Creamy Chicken", "price": "1.26" },
        { "shortDescription": "Doritos Nacho Cheese", "price": "3.35" },
        { "shortDescription": "Klarbrunn 12-PK 12 FL OZ", "price": "12.00" }
    ],
    "total": "35.35"
    }
5. Send the Request
    •	Click the “Send” button.
    •	If the server is running correctly, you should receive a json response similar to:
    { "id": "6ec78085-ef3a-435a-99d5-ae787ad48084" }
![POST request output on postman](./images/POST.png)

B) Test GET /receipts/{id}/points
1.	Copy the id from the previous response.
2.	Open a new Postman tab and change the request type to “GET”.
3.	Enter the URL: http://localhost:8080/receipts/{id}/points
    (Replace {id} with the actual receipt ID, e.g., 6ec78085-ef3a-435a-99d5-ae787ad48084)
4.	Click “Send”.
5.	You should receive a json response like:
    { "points": 28 }
![GET request output on postman](./images/GET.png)

[6] Stop the Server
To stop the server, press CTRL + C in the terminal.
