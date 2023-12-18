# APIwithGofr

# Car Garage Management Service

# Overview:
The Car Garage Management Service is a RESTful API developed in Go using the GoFr framework. It facilitates the management of cars entering, staying, and leaving a hypothetical garage.

Key Features:

Create Entry: Enables users to add new car entries to the garage by providing details such as make, model, and repair status.

List Cars in Garage: Provides a comprehensive list of cars currently present in the garage, including their respective details.

Update Entry: Allows modifications to the car's details, such as updating the repair status, marking completion of repair, or any other pertinent information.

Delete Entry: Facilitates the removal of a car's entry when it leaves the garage.

# Technology Stack:

Go (Golang): The primary programming language used to build the entire project.
GoFr Framework: Employed for creating the RESTful API endpoints and managing HTTP requests and responses.
MongoDB: Chosen as the database to store and manage car details due to its flexibility and scalability.
Unit Testing: Implemented for ensuring the reliability and correctness of the functionalities.

# Functionality:
The service emulates a straightforward car garage management system, allowing users to execute CRUD operations on car entries. It mimics real-world scenarios of cars entering, being serviced, and leaving a garage.

# Objective:
To provide a robust and scalable backend service that manages cars in a garage efficiently, allowing for easy integration with frontend applications or other services.

# Scalability & Future Enhancements:
The project's modular structure allows for easy extension with additional features, such as user authentication, advanced search functionalities, or analytics based on the car data stored in the database.

## Setting Up Development Environment

### Prerequisites

To run this project locally, ensure you have the following software/tools installed:

- **Go 1.15+**: [Download and install Go](https://golang.org/dl/).
- **MongoDB 4.0+**: Install MongoDB by following the instructions for your OS available [here](https://docs.mongodb.com/manual/installation/).

### Installation Steps

Follow these steps to set up the development environment:

1. **Clone the Repository:**

    bash
    git clone <repository_URL>
    cd <repository_name>
    

2. **Install Dependencies:**

    bash
    go mod download
    use command to install modules: go mod init <yourdirectoryname>
    
    use go get command to import packages 

    This command will download and install all required Go dependencies for the project.

3. **Start MongoDB Locally:**

    - Open a terminal and start the MongoDB server by running the appropriate command for your OS.

        Example for macOS/Linux:
        ```bash
        mongod
        ```

        Example for Windows (with default installation path):
        ```bash
        "C:\Program Files\MongoDB\Server\4.4\bin\mongod.exe"
        ```

    - Ensure MongoDB is running on the default port `27017`.

4. **Run the Application:**

    ```bash
    go run main.go
    ```

    This command starts the application locally, and it should be accessible at `http://localhost:8080`.

## API Endpoints

**POST /cars**
Description: Endpoint to add a new car.
Postman Equivalent:
Method: POST
URL: http://localhost:8080/cars
Body:
Select raw and JSON (application/json)
Body content:
json
Copy code
{
    "brand": "Toyota",
    "model": "Corolla",
    "description": "Test car"
}
Send the request.

**GET /cars**
Description: Endpoint to retrieve all cars.
Postman Equivalent:
Method: GET
URL: http://localhost:8080/cars
Send the request.
GET /cars/{id}
Description: Endpoint to retrieve a specific car by ID.
Postman Equivalent:
Method: GET
URL: http://localhost:8080/cars/1 (replace 1 with the desired car ID)
Send the request.

**PUT /cars/{id}**
Description: Endpoint to update a specific car by ID.
Postman Equivalent:
Method: PUT
URL: http://localhost:8080/cars/1 (replace 1 with the desired car ID)
Body:
Select raw and JSON (application/json)
Body content:
json
Copy code
{
    "brand": "UpdatedBrand",
    "model": "UpdatedModel",
    "description": "UpdatedDescription"
}
Send the request.

**DELETE /cars/{id}**
Description: Endpoint to delete a specific car by ID.
Postman Equivalent:
Method: DELETE
URL: http://localhost:8080/cars/1 (replace 1 with the desired car ID)
Send the request.

## Testing

### Running Unit Tests

To run the unit tests, execute the following command:

```bash
go test ./...
