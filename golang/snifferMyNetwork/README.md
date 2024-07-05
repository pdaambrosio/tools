# snifferMyNetwork

This folder contains the source code and related files for the `snifferMyNetwork` project.

## Description

The `snifferMyNetwork` project is a network sniffer tool developed in the Go programming language. It allows you to capture and analyze network traffic on your local network.

## Features

- Capture network packets
- Analyze packet headers
- Display packet information
- Filter packets based on protocols or IP addresses

## Installation

1. Clone the repository:

    ```shell
    git clone https://github.com/your-username/snifferMyNetwork.git
    ```

2. Change into the project directory:

    ```shell
    cd snifferMyNetwork
    ```

3. Build the project:

    ```shell
    go build -o packCap
    ```

4. Run the executable:

    ```shell
    ./packCap
    ```

## Usage

- To start capturing network packets, run the `packCap` executable.
- Type the interface name (e.g., `en0`, `eth0`) to start capturing packets on that interface.
- Type BPF filter expressions to filter packets based on protocols or IP addresses. For example, to capture only TCP packets, type `tcp`.
- Others examples:<br>
    `tcp` - Capture only TCP packets<br>
    `udp` - Capture only UDP packets<br>
    `ip` - Capture only IP packets<br>
    `src host <ip>` - Capture packets with a specific source IP address<br>
    `dst host <ip>` - Capture packets with a specific destination IP address<br>
    `src net <ip>` - Capture packets with a specific source network<br>
    `dst net <ip>` - Capture packets with a specific destination network<br>
    `port <port>` - Capture packets with a specific port number<br>
    `src port <port>` - Capture packets with a specific source port number<br>
    `dst port <port>` - Capture packets with a specific destination port number<br>
    `tcp and port <port>` - Capture only TCP packets with a specific port number<br>
    `udp and port <port>` - Capture only UDP packets with a specific port number<br>
    `tcp and src host <ip>` - Capture only TCP packets with a specific source IP address<br>
    `tcp and dst host <ip>` - Capture only TCP packets with a specific destination IP address<br>
- Press `Ctrl+C` to stop capturing packets.
- Use the available command-line options to filter and analyze the captured packets.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
