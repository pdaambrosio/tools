# portScanner

This folder contains the source code and related files for the `portScanner` tool.

## Description

The `portScanner` is a command-line tool written in Go that allows you to scan for open ports on a given host. It provides a simple and efficient way to check the availability of specific ports on a target machine.

## Features

- Scan for open ports on a specified host
- Specify a range of ports to scan
- Choose between TCP and UDP protocols
- Display detailed information about open ports

## Installation

1. Make sure you have Go installed on your system.
2. Clone this repository to your local machine.
3. Navigate to the `portScanner` directory.
4. Build the executable by running the following command:
    ```
    go build
    ```
5. Run the tool using the generated executable.

## Usage

To scan for open ports on a host, use the following command:

```
./portScanner <host> [options]
```

Options:

- `-p, --ports`: Specify a range of ports to scan (e.g., `1-100` or `80,443,8080`).
- `-t, --tcp`: Scan using TCP protocol (default).
- `-u, --udp`: Scan using UDP protocol.
- `-v, --verbose`: Display detailed information about open ports.

## Examples

Scan for open ports on `localhost` using the default settings:

```
./portScanner localhost
```

Scan for open ports on `example.com` using TCP protocol and ports `80` and `443`:

```
./portScanner example.com -p 80,443 -t
```

Scan for open ports on `192.168.0.1` using UDP protocol and ports `53-100`:

```
./portScanner 192.168.0.1 -p 53-100 -u
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
