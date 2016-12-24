package resource

import "github.com/bearchit/kiss/log"

func NewLogger(c *LogConfig) (logger *log.Logger, err error) {
	logger = log.New()

	if err = logger.SetFormat(c.Format); err != nil {
		return
	}

	if err = logger.SetLevel(c.Level); err != nil {
		return
	}

	if err = logger.SetOutput(c.Out, c.Path); err != nil {
		return
	}

	return
}
