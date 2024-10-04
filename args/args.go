package args

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type Args struct {
	Nickname string
	Host     string
	Port     uint
	Code     string
}

func NewCLIApp() *cli.App {
	args := &Args{}

	return &cli.App{
		Name:    "hackathon_2024_h2_go_client",
		Usage:   "Go client written for Hackathon 2024 H2 organized by KN inIT. The client is used to communicate with the server using WebSocket protocol. Your task is to implement bot logic. Each time the game state updates on the server, it is sent to you and you have to respond with your move. The game is played on a 2D grid. The game is turn-based. The player with the most points at the end of the game wins. Let the best bot win!",
		Version: "0.1.0",
		Authors: []*cli.Author{
			{
				Name: "KN inIT",
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
				Usage:       "The port on which the server is listening (1-65535)",
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
			// Validate the port number
			if args.Port < 1 || args.Port > 65535 {
				return fmt.Errorf("port must be between 1 and 65535")
			}

			// Set the metadata for the application
			c.App.Metadata = map[string]interface{}{
				"args": args,
			}
			return nil
		},
	}
}

// GetArgs returns the current instance of Args
func (a *Args) GetArgs() *Args {
	return a
}
