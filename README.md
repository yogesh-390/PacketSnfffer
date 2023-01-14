Basic Packet Sniffer

This is a basic packet sniffer tool written in Go using the gopacket library. The tool captures and prints out information on all packets passing through a specified network interface.
Usage

To use the tool, first build and run the program with the command go run sniffer.go [interface], where [interface] is the network interface you wish to capture packets on (e.g. "eth0" or "wlan0").

The tool will continuously capture and print out information on all packets passing through the specified interface.

You can also customize the tool to filter for specific packets or extract specific information from the packets using the functions and options provided by the gopacket library.
Requirements

    Go version 1.13 or above
    gopacket library (can be installed using go get github.com/google/gopacket)

Note

Please note that capturing packets on a network interface requires administrative privileges.
License

This project is licensed under the MIT License.
