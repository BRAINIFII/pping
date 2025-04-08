package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Color codes for terminal output
const (
	Reset  = "\033[0m"
	RedBG  = "\033[41m"
	GreenBG = "\033[42m"
	BlueBG = "\033[44m"
)

func tcpPing(host string, port int) (bool, float64, int) {
	address := fmt.Sprintf("%s:%d", host, port)
	start := time.Now()

	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		if opErr, ok := err.(*net.OpError); ok {
			if syscallErr, ok := opErr.Err.(*os.SyscallError); ok {
				return false, 0, int(syscallErr.Err.(syscall.Errno))
			}
		}
		return false, 0, 0
	}
	elapsed := time.Since(start).Seconds() * 1000
	conn.Close()
	return true, elapsed, 200
}

func ping(host string, port int, count int, interval time.Duration) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	pingLoop:
	for i := 0; i < count || count == 0; i++ {
		select {
		case <-sigChan:
			fmt.Println("\nPing interrupted.")
			break pingLoop
		default:
			ok, elapsed, status := tcpPing(host, port)

			statusStr := ""
			if ok {
				statusStr = fmt.Sprintf("%sOPEN%s", GreenBG, Reset)
			} else {
				statusStr = fmt.Sprintf("%sCLOSED%s", RedBG, Reset)
			}

			timeStr := ""
			if elapsed > 0 {
				timeStr = fmt.Sprintf("time=%.2fms", elapsed)
			}

			statusCodeStr := ""
			if status != 0 {
				statusCodeStr = fmt.Sprintf("%sST: %d%s", BlueBG, status, Reset)
			}

			msg := fmt.Sprintf("Probing %s:%d/tcp - %s", host, port, statusStr)
			if timeStr != "" {
				msg += " - " + timeStr
			}
			if statusCodeStr != "" {
				msg += " - " + statusCodeStr
			}

			fmt.Println(msg)
			time.Sleep(interval)
		}
	}
}

func main() {
	var (
		port    int
		count   int
		interval float64
	)

	flag.IntVar(&port, "port", 80, "Port to ping (default: 80)")
	flag.IntVar(&count, "n", 0, "Number of pings to send (0 = infinite)")
	flag.Float64Var(&interval, "c", 1.0, "Interval between pings in seconds (min: 0.1)")
	flag.Parse()

	if interval < 0.1 {
		fmt.Println("Interval cannot be less than 0.1 seconds.")
		os.Exit(1)
	}

	if flag.NArg() < 1 {
		fmt.Println("Usage: go-tcpping [options] host")
		flag.PrintDefaults()
		os.Exit(1)
	}

	host := flag.Arg(0)
	ping(host, port, count, time.Duration(interval*1000)*time.Millisecond)
}
