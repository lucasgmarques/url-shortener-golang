# Go URL Shortener

Go URL Shortener is a simple web application written in Go that allows users to shorten URLs. It provides a minimalist user interface to generate short URLs from long ones, which can be useful for sharing links in a concise format.

## Features

-   Shorten long URLs to easily shareable short URLs.
-   Redirect users from short URLs to their original long URLs.
-   Lightweight and easy-to-use web interface.

## Prerequisites

-   Go 1.17 or later installed on your system.
-   Docker (optional) if you prefer running the application in a containerized environment.

## Getting Started

### Running Locally

1.  Clone the repository:
    
    `git clone https://github.com/your-username/go-url-shortener.git
    cd go-url-shortener` 
    
2.  Build and run the application:
    
    `go build -o app .
    ./app` 
    
3.  Access the application at [http://localhost:8080](http://localhost:8080/) in your web browser.
    

### Running with Docker

1.  Build the Docker image:
    
    `docker build -t go-url-shortener .` 
    
2.  Run the Docker container:
    
    `docker run -p 8080:8080 go-url-shortener` 
    
3.  Access the application at [http://localhost:8080](http://localhost:8080/) in your web browser.
    

## Usage

-   Visit the homepage of the application.
-   Enter the long URL you want to shorten and submit the form.
-   Copy the generated short URL and use it for sharing.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request for any improvements or features you'd like to see.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgements

-   This project was inspired by the need for a simple URL shortener service.
-   Special thanks to the Go community for providing excellent resources and libraries.
