package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retornar a aplicação de linha de comando pronta
//para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca de IPs e Nomes de Servidores na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	//App.commands vai receber um Slice de comandos
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de IPs",
			Flags:  flags,
			Action: buscaIps,
		},
		{
			Name:   "servers",
			Usage:  "Busca de Servidores",
			Flags:  flags,
			Action: buscaServers,
		},
	}

	return app
}

func buscaIps(c *cli.Context) {
	host := c.String("host")

	//Usando pacote Net
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
func buscaServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host) //NS Name Server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
