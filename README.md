# pong

A blazing fast TCP-based ping utility written in Go. Ideal for networks where ICMP is restricted or blocked. Works smoothly on Windows, macOS, and Linux.

## ⚡ Features

- Ping a host and port using TCP.
- Determine if a specific service is reachable.
- Specify the number of pings to send.
- Control the interval between pings.
- Lightweight and dependency-free (except for Go stdlib).
- Cross-platform compatibility (builds on all major OSes).

## 📦 Installation

Make sure you have Go installed (https://golang.org/dl).

To install:

1. Clone this repo:
```
git clone https://github.com/brainifii/pong.git
```
2. Navigate into the directory:
```
cd pong
```
3. Build it:
```
5. go build -o pong pong.go
```
You’ll get a binary named `pong` (or `pong.exe` on Windows) in the current directory.

## 🚀 Usage

Basic usage:
```
./pong example.com 80
```

Ping with a custom interval:
```
./pong example.com 443 -c 0.5
```

Send a specific number of pings:
```
./pong example.com 443 -n 10
```

## ⛏ Arguments

positional:
- host: The target hostname or IP address.
- port: The TCP port to check (default: 80)

flags:
- -n: Number of times to ping (default: infinite)
- -c: Interval between pings in seconds (default: 1s, minimum: 0.1s)

## 🔥 Example

Ping google.com on port 443 every 0.5 seconds, 5 times:
```
./pong google.com 443 -n 5 -c 0.5
```
## 💬 Why Go?

- Fast and compiled
- Portable across platforms
- Perfect for system utilities

## 🧠 Pro Tip

You can cross-compile for another OS easily:
```
GOOS=windows GOARCH=amd64 go build -o pong.exe pong.go
```
## 🛠 Dependencies

None. Everything runs on Go's standard library.

## 🐛 Got Issues?

Open one on the repo or hit me up at hello@brainifii.com
