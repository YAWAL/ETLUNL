package logging

import (
	"github.com/Sirupsen/logrus"
)

var Log = logrus.New()

func init() {

	Log.Formatter = &logrus.JSONFormatter{}
}
