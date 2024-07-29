import socket
import argparse
import sys
import time
from colorama import Fore, Style, Back

def tcp_ping(host, port):
    try:
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.settimeout(1)
        start_time = time.time()
        sock.connect((host, port))
        elapsed_time = time.time() - start_time
        sock.close()
        return True, elapsed_time, 200
    except Exception as e:
        if hasattr(e, 'errno'):
            return False, None, e.errno
        else:
            return False, None, None

def ping(host, port, count=None, interval=1):
    if count is None:
        count = sys.maxsize

    try:
        for _ in range(count):
            result, elapsed_time, status_code = tcp_ping(host, port)
            if result:
                status = f"{Back.GREEN}OPEN{Style.RESET_ALL}"
            else:
                status = f"{Back.RED}CLOSED{Style.RESET_ALL}"
            
            time_str = f"time={elapsed_time*1000:.2f}ms" if elapsed_time is not None else ""
            status_code_str = f"{Back.BLUE}ST: {status_code}{Style.RESET_ALL}" if status_code else ""

            if time_str and status_code_str:
                print(f"Probing {host}:{port}/tcp - {status} - {time_str} - {status_code_str}".replace("  ", ""))
            elif time_str:
                print(f"Probing {host}:{port}/tcp - {status} - {time_str}")
            elif status_code_str:
                print(f"Probing {host}:{port}/tcp - {status} - {status_code_str}".replace("  ", ""))
            else:
                print(f"Probing {host}:{port}/tcp - {status}")

            time.sleep(interval)
    except KeyboardInterrupt:
        print("\nPing interrupted.")
        sys.exit(0)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Ping a host and port using TCP")
    parser.add_argument("host", help="The host to ping")
    parser.add_argument("port", nargs='?', type=int, default=80, help="The port to ping (default: 80)")
    parser.add_argument("-n", type=int, help="Number of pings to send (default: indefinitely)")
    parser.add_argument("-c", type=float, default=1, help="Interval between pings in seconds (default: 1, min: 0.1)")
    args = parser.parse_args()
    
    if args.c < 0.1:
        print("Interval cannot be less than 0.1 seconds.")
        sys.exit(1)
    
    ping(args.host, args.port, args.n, args.c)