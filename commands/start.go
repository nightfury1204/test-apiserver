package commands

import (
	"github.com/spf13/cobra"
	"github.com/nightfury1204/test-apiserver/pkg/server"
)

// NewCommandStartWardleServer provides a CLI handler for 'start master' command
// with a default ServerOptions.
func NewCommandStartServer(defaults *server.TryapiServerOptions, stopCh <-chan struct{}) *cobra.Command {
	o := *defaults
	cmd := &cobra.Command{
		Short: "Launch a wardle API server",
		Long:  "Launch a wardle API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}
			if err := o.RunWardleServer(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	o.RecommendedOptions.AddFlags(flags)

	return cmd
}