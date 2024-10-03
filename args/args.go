package args

import (
	"github.com/urfave/cli/v2"
)

type Args struct {
	Nickname string
	Host     string
	Port     uint
	Code     string
}

func (a *Args) GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "nickname",
			Aliases:     []string{"n"},
			Usage:       "Nickname of the agent that will be displayed in the game",
			Destination: &a.Nickname,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "host",
			Usage:       "The IP address or domain name of the server to connect to",
			Value:       "localhost",
			Destination: &a.Host,
		},
		&cli.UintFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Usage:       "The port on which the server is listening",
			Value:       5000,
			Destination: &a.Port,
		},
		&cli.StringFlag{
			Name:        "code",
			Aliases:     []string{"c"},
			Usage:       "Optional access code required to join the server",
			Value:       "",
			Destination: &a.Code,
		},
	}
}
