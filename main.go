package main

import (
	"flag"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func check(err error) {
	if err != nil {
		color.Red("[!] A ocurrido un error:", err)
		os.Exit(1)
	}
}

func get_ttl(output_ping string) {
	pattern := ` \w\w\w\=\d\d\d?`
	delimitador := "="
	const WINDOWS_TTL = 128
	const LINUX_TTL = 64

	re := regexp.MustCompile(pattern)
	ttl_array := re.FindStringSubmatch(output_ping)

	if ttl_array != nil {

		color.Green("[+] Respondio")
		ttl, err := strconv.Atoi(strings.Split(ttl_array[0], delimitador)[1])
		check(err)
		if ttl == LINUX_TTL {
			color.Green("[+] El ttl es: %v --> linux", ttl)
			os.Exit(0)
		}
		if ttl == WINDOWS_TTL {
			color.Green("[+] El ttl es: %v --> windows", ttl)
			os.Exit(0)
		} else {
			color.Yellow("[-] El ttl es: %v --> Sistema Operativo desconocido", ttl)
			os.Exit(0)
		}
	}
	color.Red("[!] La maquina no respondio")
	os.Exit(1)
}

func windows(ip_address string) {
	out, err := exec.Command("C:/Windows/System32/ping", ip_address).Output()
	check(err)
	output_ping := string(out[:])
	get_ttl(output_ping)
}

func linux(ip_address string) {
	out, err := exec.Command("/usr/bin/ping", ip_address).Output()
	check(err)
	output_ping := string(out[:])
	get_ttl(output_ping)
}

func verify_os(ip_address string) {

	if runtime.GOOS == "windows" {
		color.Green("[+] Ejecutando ping en windows...")
		windows(ip_address)
		return
	}
	if runtime.GOOS == "linux" {
		color.Green("[+] Ejecutando ping en linux...")
		linux(ip_address)
		return
	} else {
		color.Red("[-] Sistema Operativo no apto")
		os.Exit(1)
	}
}

func main() {
	ip_address_flag := flag.String("ip", "", "Direccion IP")
	flag.Parse()
	ip_address := *ip_address_flag
	color.Yellow("[*] Verificando tu Sistema Operativo...")
	verify_os(ip_address)
}
