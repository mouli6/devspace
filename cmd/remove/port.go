package remove

import (
	"errors"

	"github.com/devspace-cloud/devspace/cmd/flags"
	"github.com/devspace-cloud/devspace/pkg/devspace/config/configutil"
	"github.com/devspace-cloud/devspace/pkg/devspace/configure"
	"github.com/devspace-cloud/devspace/pkg/util/log"
	"github.com/spf13/cobra"
)

type portCmd struct {
	*flags.GlobalFlags

	LabelSelector string
	RemoveAll     bool
}

func newPortCmd(globalFlags *flags.GlobalFlags) *cobra.Command {
	cmd := &portCmd{GlobalFlags: globalFlags}

	portCmd := &cobra.Command{
		Use:   "port",
		Short: "Removes forwarded ports from a devspace",
		Long: `
#######################################################
############### devspace remove port ##################
#######################################################
Removes port mappings from the devspace configuration:
devspace remove port 8080,3000
devspace remove port --label-selector=release=test
devspace remove port --all
#######################################################
	`,
		Args: cobra.MaximumNArgs(1),
		RunE: cmd.RunRemovePort,
	}

	portCmd.Flags().StringVar(&cmd.LabelSelector, "label-selector", "", "Comma separated key=value selector list (e.g. release=test)")
	portCmd.Flags().BoolVar(&cmd.RemoveAll, "all", false, "Remove all configured ports")

	return portCmd
}

// RunRemovePort executes the remove port command logic
func (cmd *portCmd) RunRemovePort(cobraCmd *cobra.Command, args []string) error {
	// Set config root
	configExists, err := configutil.SetDevSpaceRoot(log.GetInstance())
	if err != nil {
		return err
	}
	if !configExists {
		return errors.New("Couldn't find a DevSpace configuration. Please run `devspace init`")
	}

	config, err := configutil.GetBaseConfig(cmd.ToConfigOptions())
	if err != nil {
		return err
	}

	err = configure.RemovePort(config, cmd.RemoveAll, cmd.LabelSelector, args)
	if err != nil {
		return err
	}

	log.Done("Successfully removed port")
	return nil
}
