# Book Rental (Backend using Go)

## Table of Contents

## Features

## Project Structure

## Requirements

- Go 1.22.3 or later


## Installation

## Build

## Certificate Generation 

### 1. Generate the Private Key:
```Bash
$ openssl genrsa -out key.pem 2048
```
- openssl genrsa: The OpenSSL command for generating RSA keys.
- -out key.pem: Specifies the filename to save the private key. Important: Keep this file secure! Permissions should be 0600 (readable and writable only by the owner).
- 2048: The key size in bits. 2048 bits is a good standard; you can also use 4096 for even stronger security.

### 2. Generate a Self-Signed Certificate (for development):

```Bash
$ openssl req -x509 -new -nodes -key key.pem -out cert.pem -days 365 -subj "/C=US/ST=YourState/L=YourCity/O=YourOrganization/OU=YourUnit/CN=localhost"
```
- openssl req: The OpenSSL command for certificate requests and generation.
- -x509: Creates a self-signed certificate.
- -new: Creates a new certificate request.
- -nodes: (Optional) Do not encrypt the private key with a passphrase. For development, this is often convenient. In production, you should almost always encrypt your private key with a strong passphrase. If you omit -nodes, OpenSSL will prompt you for a passphrase.
- -key key.pem: Specifies the private key file.
- -out cert.pem: Specifies the filename to save the certificate.
- -days 365: Sets the certificate validity period to 365 days (1 year). Adjust as needed.
- -subj "/C=US/ST=YourState/L=YourCity/O=YourOrganization/OU=YourUnit/CN=localhost": Sets the subject of the certificate. Replace the values with your own information. CN (Common Name) is the most important; it should be the domain name or hostname for which the certificate is valid (e.g., localhost for local development).

### 3. (Optional) Encrypt the Private Key (for production):

If you want to encrypt the private key (highly recommended for production), omit the -nodes option in the openssl req command.  OpenSSL will then prompt you for a passphrase.  You'll need to provide this passphrase whenever you use the private key.   

### 4. Verify the Certificate:

```Bash
$ openssl x509 -in cert.pem -text -noout
```
This command will display the contents of the certificate in human-readable format, allowing you to verify the information.   

## Usage

## Contributing

## License