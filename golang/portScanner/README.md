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
    go build -o portScannerV1
    ```

5. Run the tool using the generated executable.

## Usage

To scan for open ports on a host, use the following command:

```
./portScannerV1 -i <ip> -p <portRange>
```

Options:

- `-i: Specify the IP address or hostname of the target machine.
- `-p: Specify a range of ports to scan (e.g., 1024 are ports 1 to 1024).

## Examples

Scan for open ports on `192.168.10.1` using the default settings:

```
./portScannerV1 -i 192.168.10.1

Scan for open ports on `192.168.10.1` ports range `1-65535`:

```
./portScannerV1 -i 192.168.10.1 -p 65535
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
