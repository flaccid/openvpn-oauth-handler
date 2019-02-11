package microsoft

import (
	"os"
	"os/exec"
	"strings"

	//"fmt"
	//"reflect"
	//"github.com/davecgh/go-spew/spew"

	"github.com/Azure/go-autorest/autorest/azure/auth"
	log "github.com/Sirupsen/logrus"
)

var (
	authorized bool = false
)

func AuthAzureCli() {
	authorizer, err := auth.NewAuthorizerFromCLI()
	log.Debug(authorizer)

	//fmt.Println(reflect.TypeOf(&authorizer))
	//spew.Dump(authorizer)

	if err != nil {
		log.Errorf(err.Error())

		// check if the error is recoverable, e.g. needs az login
		if strings.Contains(err.Error(), "Please run 'az login' to setup account") {
			log.Info("user needs to login, will now execute az login")
			cmd := "az"
			args := []string{"login"}
			if err := exec.Command(cmd, args...).Run(); err != nil {
				log.Fatalf("az login failed", os.Stderr, err)
				authorized = false
			} else {
				authorized = true
			}
		}
	} else {
		authorized = true
	}

	if authorized {
		log.Info("user authenticated")
	} else {
		log.Fatalf("failed to get authorization", os.Stderr, err)
		os.Exit(1)
	}
}
