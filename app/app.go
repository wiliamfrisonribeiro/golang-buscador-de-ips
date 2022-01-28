package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {

	app := cli.NewApp()

	app.Name = "Aplicação de Linha de Comando "
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca Ips de endereços na interner",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome do servidores na internet",
			Flags:  flags,
			Action: buscaServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {

		fmt.Println(ip)
	}
}

func buscaServidores(c *cli.Context) {

	host := c.String("host")

	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, s := range servers {

		fmt.Println(s.Host)
	}
}
