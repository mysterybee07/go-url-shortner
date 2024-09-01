# go-url-shortner

This project is a simple URL shortener service built using Go and the Fiber web framework. It allows users to shorten long URLs and retrieve the original URLs using the shortened versions. The service is designed to be lightweight, fast, and easy to use.

# Features

i. Shorten long URLs

ii. Retrieve original URLs from shortened versions

iii. Basic error handling for invalid URLs
 
iv. Unit tests using github.com/stretchr/testify/assert

# Prerequisites

Before running the service, ensure you have the following installed on your machine:

Go (go1.23.0)

Git

# Installation
Clone the repository:

Using https:

    git clone https://github.com/mysterybee07/go-url-shortner.git
    
or Using ssh:

    git clone git@github.com:mysterybee07/go-url-shortner.git

# Install dependencies: 

Redirect to the project directory:

       cd go-url-shortener

Use Go modules to install the required dependencies:
  
       go mod tidy

# Running the Service
To start the URL shortener service, run the following command in your terminal:

        go run main.go

OR, you can use air command to enjoy live-reloading function if you have air installed in your device 
        
        air
    
Follow the url to install air: 
        https://github.com/air-verse/air

The service will start on http://localhost:8080 by default. 

# API Endpoints

POST /shorten

Request Body: JSON object with the long URL

Response: JSON object with the shortened URL

![alt text](images/image-1.png)

GET /{shortcode}

Request: Access the shortened URL directly

Response: Redirects to the original long URL

![alt text](images/image-2.png)

# Running Tests
To ensure that the service is functioning correctly, you can run the tests using the following command:

       go test ./tests -v

![alt text](images/image-3.png)

This command will execute all tests in the project and display the results in the terminal. The tests are written using the github.com/stretchr/testify/assert package for easy assertions.

