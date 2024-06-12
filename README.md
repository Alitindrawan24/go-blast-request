# Go Blast Request

This Go program sends parallel HTTP requests to a specified endpoint, with the ability to define the HTTP method, request headers, body, and query parameters. The configuration is read from a JSON file.

## Prerequisites

- Go installed on your machine (version 1.16 or higher recommended)

## Configuration

Create a `target.json` file in the same directory as the `main.go` file with the following structure:

```json
{
    "url": "https://jsonplaceholder.typicode.com/posts/1",
    "method": "GET",
    "headers": {
        "Content-Type": "application/json"
    },
    "body": {},
    "query_params": {
        "param1": "value1",
        "param2": "value2"
    },
    "count": 10
}
```

- url: The endpoint to which the requests will be sent.
- method: The HTTP method to use for the - requests (e.g., GET, POST, PUT, PATCH).
- headers: An object containing any headers to include in the requests.
- body: An object representing the body of the request (only used for methods like POST, PUT, PATCH).
- query_params: An object containing query parameters to include in the URL.
- count: The number of parallel requests to send.

## Usage

- Clone the repository or copy the code files to your local machine.
- Ensure that the target.json file is in the same directory as main.go.
- Open a terminal and navigate to the directory containing main.go.
- Run the program using the following command:
```
go run main.go
```
The program will read the configuration from target.json and send the specified number of parallel HTTP requests to the endpoint. The responses will be printed to the console.

## Example

Given the target.json configuration above, the program will send 10 parallel GET requests to https://jsonplaceholder.typicode.com/posts/1 with the specified headers and query parameters.

## License
This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/MarketingPipeline/README-Quotes/blob/main/LICENSE) file for details.

## Contributors
<a href="https://github.com/alitindrawan24/go-blast-request/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=alitindrawan24/go-blast-request" />
</a>