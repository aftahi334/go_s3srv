# GoS3Server

GoS3Server is a simple implementation of an S3 (Simple Storage Service) server written in Go. It provides basic functionalities to store, retrieve, and manage data in a way similar to AWS S3.

## Features

- Object Storage: Upload, download, and delete objects.
- Bucket Management: Create and delete buckets.
- Metadata Support: Store and retrieve object metadata.
- Authentication: Basic authentication support.
- S3 API Compatibility: Supports a subset of the S3 API for easy integration with existing tools.

## Requirements

- Go 1.18+
- Docker (optional, for containerized deployment)

## Installation

1. Clone the repository:
	git clone https://github.com/aftahi334/go_s3srv.git
	cd go_s3srv

2. Build the project:
	go build -o gos3srv


3. Run the server:
	 ./gos3srv

## Configuration

Configuration is done through environment variables:

- `S3_SERVER_PORT`: The port on which the server will listen (default: `8080`).
- `S3_DATA_DIR`: Directory to store the data (default: `./data`).
- `S3_ACCESS_KEY`: Access key for authentication.
- `S3_SECRET_KEY`: Secret key for authentication.

Example:
	
	export S3_SERVER_PORT=8080

	export S3_DATA_DIR=/path/to/data

	export S3_ACCESS_KEY=your_access_key

	export S3_SECRET_KEY=your_secret_key

	./gos3server

## Usage

### Upload Object
	curl -XPOST --data-binary @filename localhost:8080/filename

### Download Object
	curl -XGET localhost:8080/filename

### Delete Object
	curl -XDELETE localhost:8080/filename

### List Objects
	curl XGET localhost:8080/list

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

Inspired by AWS S3 and various open-source S3 server implementations.