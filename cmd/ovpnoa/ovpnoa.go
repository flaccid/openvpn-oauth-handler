package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/flaccid/openvpn-oauth-handler/providers/microsoft"
	"github.com/urfave/cli"
)

var (
	VERSION = "v0.0.0-dev"
)

func beforeApp(c *cli.Context) error {
	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "ovpnoa"
	app.Version = VERSION
	app.Usage = "external oauth handler for openvpn authentication"
	app.Action = start
	app.Before = beforeApp
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "provider",
			Usage:  "oauth provider",
			EnvVar: "OAUTH_PROVIDER",
		},
		cli.StringFlag{
			Name:   "provider-method",
			Usage:  "provider oauth method",
			EnvVar: "OAUTH_PROVIDER_METHOD",
		},
		cli.StringFlag{
			Name:   "client-id",
			Usage:  "oauth client id",
			EnvVar: "CLIENT_ID",
		},
		cli.StringFlag{
			Name:   "client-secret",
			Usage:  "oauth client secret",
			EnvVar: "CLIENT_SECRET",
		},
		cli.StringFlag{
			Name:   "redirect-url",
			Usage:  "oauth redirect url aka the callback url",
			EnvVar: "REDIRECT_URL",
		},
		cli.StringFlag{
			Name:   "tenant-id",
			Usage:  "azure tenant id (used with microsoft provider)",
			EnvVar: "TENANT_ID",
		},
		// cli.BoolFlag{
		// 	Name:  "callback-daemon",
		// 	Usage: "run the oauth callback daemon only",
		// },
		cli.BoolFlag{
			Name:  "dry",
			Usage: "run in dry mode",
		},
		cli.BoolFlag{
			Name:  "debug,d",
			Usage: "run in debug mode",
		},
	}
	app.Run(os.Args)
}

func start(c *cli.Context) error {
	log.Infof("starting auth with %v provider", c.String("provider"))

	// if c.Bool("callback-daemon") {
	// 	microsoft.CallbackDaemon()
	// } else {
	switch c.String("provider") {
	case "microsoft":
		microsoft.AuthAzureCli()
	default:
		log.Error("please specify a provider")
	}

	return nil
}
