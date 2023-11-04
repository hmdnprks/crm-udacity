# CRM Backend Project

This project represents the backend of a Customer Relationship Management (CRM) web application. The server supports various functionalities for managing customer data.

## Description

The CRM Backend project is developed using Go language and provides a RESTful API to interact with customer data. The API allows users to perform the following operations:

* Get a list of all customers
* Get data for a single customer
* Add a new customer
* Update a customer's information
* Remove a customer

## Installation

To run the CRM Backend project, follow these steps:

1. Make sure you have Go installed on your system. If not, you can download and install it from the official Go website: <https://golang.org/>
2. Clone the project repository from GitHub:

   ```
   git clone <repository_url>
   ```
3. Change into the project directory:

   ```
   cd crm-udacity
   ```
4. Install the project dependencies:

   ```
   go mod download
   ```

## Launch

To launch the CRM Backend server, run the following command:

```
go run main.go
```

The server will start running on `localhost` at port `3000`.

## Usage

Once the server is up and running, you can interact with the API using tools like Postman or cURL. Here are the available API endpoints:

* **GET /customers**: Get a list of all customers
* **GET /customers/{id}**: Get data for a specific customer
* **POST /customers**: Add a new customer
* **PUT /customers/{id}**: Update a customer's information
* **DELETE /customers/{id}**: Remove a customer

Make sure to replace `{id}` with the actual customer ID in the URL when performing operations on a specific customer.

For example, to get a list of all customers, send a GET request to `http://localhost:3000/customers`. Similarly, you can use the respective HTTP methods and endpoints to perform other operations.

Please note that the server will return JSON responses for all API endpoints except the home route (`/`), which serves a static HTML page.

## Additional Information

* The project uses the `encoding/json` package to handle JSON data.
* The server uses a router package (`gorilla/mux`, `http.ServeMux`, etc.) for HTTP method-based routing and URL path variables.
* The project includes a mock "database" to store customer data.
* Error handling is implemented for non-existent customers, returning a 404 status code and appropriate response.

For more details and implementation guidelines, please refer to the [project rubric](https://review.udacity.com/#!/rubrics/4856/view).

Feel free to reach out if you have any questions or need further assistance. Happy coding!
