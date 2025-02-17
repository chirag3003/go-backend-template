# Go Backend Template

This is a starter template in Go using Go Fiber. This simple template allows even beginners to start writing their backend code without worrying about the boilerplate. The folder structure is easy to understand, and it includes implementations for authentication, MongoDB connection, and file uploading using S3.

## Features

- Go Fiber
- Authentication (JWT)
- MongoDB connection
- File uploading using S3
- Image optimization

## Folder Structure

```
/home/chirag/Projects/go-backend-template/
├── main.go
├── config/
├── controllers/
├── db/
├── helpers/
├── middlewares/
├── models/
├── repository/
├── routes/
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.23+
- MongoDB

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/chirag3003/go-backend-template.git
    ```
2. Navigate to the project directory:
    ```sh
    cd go-backend-template
    ```
3. Install the dependencies:
    ```sh
    go mod tidy
    ```
4. Copy the example environment file to `.env`:
    ```sh
    cp .env.example .env
    ```

### Usage

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start at `http://localhost:5000`.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request.

## Author

Chirag Bhalotia  
GitHub: [chirag3003](https://github.com/chirag3003)

## License

This project is licensed under the MIT License.