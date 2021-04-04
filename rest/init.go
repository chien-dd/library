package rest

import (
	"os"

	"github.com/chien-dd/library/pencil"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = pencil.NewLogger(
		os.Stdout,
		pencil.InfoLevel,
		pencil.ProductFormatter,
	)
	logger.Infof("[REST] Initialize logger ...\n")
}
