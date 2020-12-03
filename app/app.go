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

	//App.commands vai receber um Slice de comandos
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca de IPs",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com",
				},
			},
			Action: buscaIps,
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
