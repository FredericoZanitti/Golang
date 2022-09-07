package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser importada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidores na internet"

	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "zanitti.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endreços na web",
			Flags:  flag,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "busca nome dos servidores na internet",
			Flags:  flag,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	//net
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	//net
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
