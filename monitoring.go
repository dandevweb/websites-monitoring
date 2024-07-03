package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {
	showIntro()

	for {
		showMenu()
		command := readCommand()
		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Danilo"
	version := 1.1

	fmt.Println("Hello,", name)
	fmt.Println("Este programa está na versão,", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("O comando escolhido foi:", commandRead)
	fmt.Println("")

	return commandRead
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites := readFileSites()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	statusCode := resp.StatusCode

	if statusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		logRegister(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", statusCode)
		logRegister(site, false)
	}
}

func readFileSites() []string {
	var sites []string
	file, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if err == io.EOF || line == "" {
			break
		}

		sites = append(sites, line)
	}

	file.Close()

	return sites
}

func logRegister(site string, isOnline bool) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(isOnline) + "\n")

	file.Close()

}

func showLogs() {
	fmt.Println("Exibindo Logs...")

	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(file))

}
