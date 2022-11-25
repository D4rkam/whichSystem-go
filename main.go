package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

func check(err error) {
	if err != nil {
		fmt.Print("A ocurrido un error:", err)
		os.Exit(1)
	}
}

func windows(ip_address string) {
	out, err := exec.Command("ping", ip_address).Output()
	check(err)
	output_command := string(out[:])
	fmt.Println(output_command)
	re := regexp.MustCompile(`Respuesta desde 181.85.123.83: bytes=32 tiempo<1m TTL=64 
								Respuesta desde 181.85.123.83: bytes=32 tiempo=1ms TTL=64
								Respuesta desde 181.85.123.83: bytes=32 tiempo<1m TTL=64 
								Respuesta desde 181.85.123.83: bytes=32 tiempo=1ms TTL=64`)
	ttl := re.FindStringSubmatch("TLL=64")
	check(err)
	fmt.Println("La salida del comando es:", ttl[0])

}

func main() {
	var ip_address string
	fmt.Print("Ingrese la direccion ip:")
	fmt.Scan(&ip_address)
	if runtime.GOOS == "windows" {
		windows(ip_address)
	}
}
