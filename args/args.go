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

func NewCLIApp() *cli.App {
	args := &Args{} // Create a pointer to Args

	return &cli.App{
		Name:    "agent",
		Usage:   "Configure the agent and connect to a server",
		Version: "1.0.0",
		Authors: []*cli.Author{
			{
				Name:  "Author Name",
				Email: "author@example.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "nickname",
				Aliases:     []string{"n"},
				Usage:       "Nickname of the agent that will be displayed in the game",
				Destination: &args.Nickname,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "host",
				Usage:       "The IP address or domain name of the server to connect to",
				Value:       "localhost",
				Destination: &args.Host,
			},
			&cli.UintFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Usage:       "The port on which the server is listening",
				Value:       5000,
				Destination: &args.Port,
			},
			&cli.StringFlag{
				Name:        "code",
				Aliases:     []string{"c"},
				Usage:       "Optional access code required to join the server",
				Value:       "",
				Destination: &args.Code,
			},
		},
		Action: func(c *cli.Context) error {
			// Store the args in the app's metadata
			c.App.Metadata = map[string]interface{}{
				"args": args,
			}
			return nil
		},
	}
}

func (a *Args) GetArgs() *Args {
	return a
}
