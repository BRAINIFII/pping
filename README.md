# pping

A Python-based utility to simulate ping using TCP on specific open ports. This tool is particularly useful in environments where ICMP packets are disabled for security reasons. 

## Features

- Ping a host and port using TCP.
- Check if specific applications are working properly.
- Specify the number of pings to send.
- Set the interval between pings.

## Usage

```bash
usage: pping.py [-h] [-n N] [-c C] host [port]

Ping a host and port using TCP

positional arguments:
  host        The host to ping
  port        The port to ping (default: 80)

options:
  -h, --help  Show this help message and exit
  -n N        Number of pings to send (default: indefinitely)
  -c C        Interval between pings in seconds (default: 1, min: 0.1)
```

## Installation

To use `pping`, you need Python installed on your system. You can install the required dependencies using pip:

```bash
pip install -r requirements.txt
```

## Example
Ping a host on the default port 80:

```bash
python pping.py example.com
```

Ping a host on port 443:
```bash
python pping.py example.com 443
```