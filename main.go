package main

import (
	"github.com/sirupsen/logrus"
	"github.com/eosioafrica/ecte/ecte"
)

func main() {

	ecte := ecte.New()

	ecte.Run()

	if ecte.Err != nil {

		logrus.Error("Ooooops! ", ecte.Err)
	}
}




