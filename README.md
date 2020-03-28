# EnterpriseComputingCA
email

using; 
GO-KIT
    - transport layer
    - end point layer each of which represents an rpc method 
    - service layer

POSTGRESSQL
    - as a way to store the data given
    -SQL

POSTMAN
    - To interface with the API 
    - allows the user to POST GET DELETE and many more
    - POST localhost:8080/user sends the data to the server 
            {
                "email":"test@here.com",
                "password":"Hudders1",
                "body":"hello david hope you are well, sorry my hand in is late",
                "toEmail":"henry@there.com"
            }
    - GET localhost:8080/user/b7b95bad-9fa7-4e9e-91ff-afe10fca3e5b
            the hexidecimal string following the user/ is used to get that users mail, it is there id

GORILLA MUX 
    - nice way to interface with go and gokit
