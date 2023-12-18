# APIwithGofr

# Car Management API

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
    use command to install modules: go mod init <your directory name>
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

### `POST /cars`

- Description: Endpoint to add a new car.
- Example Request:
    ```bash
    curl -X POST http://localhost:8080/cars -d '{"brand":"Toyota","model":"Corolla","description":"Test car"}'
    ```
- Expected Response: HTTP Status 201 Created

### `GET /cars`

- Description: Endpoint to retrieve all cars.
- Example Request:
    ```bash
    curl http://localhost:8080/cars
    ```
- Expected Response: A JSON array of cars.

### `GET /cars/{id}`

- Description: Endpoint to retrieve a specific car by ID.
- Example Request:
    ```bash
    curl http://localhost:8080/cars/1
    ```
- Expected Response: Details of the car with the specified ID.

### `PUT /cars/{id}`

- Description: Endpoint to update a specific car by ID.
- Example Request:
    ```bash
    curl -X PUT http://localhost:8080/cars/1 -d '{"brand":"UpdatedBrand","model":"UpdatedModel","description":"UpdatedDescription"}'
    ```
- Expected Response: Updated details of the car with the specified ID.

### `DELETE /cars/{id}`

- Description: Endpoint to delete a specific car by ID.
- Example Request:
    ```bash
    curl -X DELETE http://localhost:8080/cars/1
    ```
- Expected Response: HTTP Status 200 OK on successful deletion.

## Testing

### Running Unit Tests

To run the unit tests, execute the following command:

```bash
go test ./...
