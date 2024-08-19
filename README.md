# GoCommerce

GoCommerce is a powerful, flexible, and easy-to-use e-commerce platform written in Go (Golang). It is designed to be simple for developers to customize while offering all the essential features needed to run a modern online store.

## Features

- **Product Management:** Add, edit, delete, and manage products.
- **Order Management:** Process and track customer orders.
- **User Authentication:** Secure user registration and login.
- **Shopping Cart:** Add and remove items from the cart.

## Getting Started

Follow these instructions to set up and run the GoCommerce project on your local machine for development and testing purposes.

### Prerequisites

Ensure that you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.23 or later)
- Docker
- Docker compose

### Installation

1. Clone the repository:   
	```sh
	git clone https://github.com/nubufi/GoCommerce.git 
	cd GoCommerce
	```
    
2. Build and deploy the app:   
    ```sh
    docker compose up --build
    ```
    
### Running Tests

To run the tests, use:

```sh
go test ./...
```

## Configuration

The application can be configured using environment variables. Refer to the `.env` file for available configuration options such as database credentials, API keys, and more.

## Usage

The API documentation can be accessed at `http://localhost:8080/docs/index.html`.

## Contributing

We welcome contributions from the community! To contribute to GoCommerce:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Create a new Pull Request.

Please read our [Contributing Guide](CONTRIBUTING.md) for more details.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

For any inquiries or feedback, feel free to reach out to us:

- **Website:** [www.nubufi.com](https://www.nubufi.com)
- **Email:** numanburakfidan@yandex.com
