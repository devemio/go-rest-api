package contracts

import "github.com/sirupsen/logrus"

type Logger interface {
	logrus.FieldLogger
}
