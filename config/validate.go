package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Validate the configuration file for sanity
func (c *config) Validate() error {
	if c.ApiKeys.Slack == "" || c.ApiKeys.Pagerduty.Key == "" {
		return fmt.Errorf("you must provide API keys for Slack and Pagerduty")
	}

	if c.ApiKeys.Pagerduty.Org == "" {
		return fmt.Errorf("you must provide an org name for Pagerduty (<org>.pagerduty.com)")
	}

	if len(c.Groups) == 0 {
		return fmt.Errorf("you must specify at least one group")
	}

	for i, group := range c.Groups {
		if group.Name == "" {
			return fmt.Errorf("must specify group name for group %d", i)
		}

		if len(group.Schedules) == 0 {
			return fmt.Errorf("must specify at least one schedule for group %s", group.Name)
		}
	}

	log.WithFields(log.Fields{"groups": len(c.Groups)}).Debug("Loaded config")

	return nil
}
