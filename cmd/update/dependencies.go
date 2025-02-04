package update

import (
	"github.com/devspace-cloud/devspace/cmd/flags"
	"github.com/devspace-cloud/devspace/pkg/devspace/config/configutil"
	"github.com/devspace-cloud/devspace/pkg/devspace/config/generated"
	"github.com/devspace-cloud/devspace/pkg/devspace/dependency"
	"github.com/devspace-cloud/devspace/pkg/util/log"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// dependenciesCmd holds the cmd flags
type dependenciesCmd struct {
	*flags.GlobalFlags

	AllowCyclicDependencies bool
}

// newDependenciesCmd creates a new command
func newDependenciesCmd(globalFlags *flags.GlobalFlags) *cobra.Command {
	cmd := &dependenciesCmd{GlobalFlags: globalFlags}

	dependenciesCmd := &cobra.Command{
		Use:   "dependencies",
		Short: "Updates the git repositories of the dependencies defined in the devspace.yaml",
		Long: `
#######################################################
############ devspace update dependencies #############
#######################################################
Updates the git repositories of the dependencies defined
in the devspace.yaml
#######################################################
	`,
		Args: cobra.NoArgs,
		RunE: cmd.RunDependencies,
	}

	dependenciesCmd.Flags().BoolVar(&cmd.AllowCyclicDependencies, "allow-cyclic", false, "When enabled allows cyclic dependencies")

	return dependenciesCmd
}

// RunDependencies executes the functionality "devspace update dependencies"
func (cmd *dependenciesCmd) RunDependencies(cobraCmd *cobra.Command, args []string) error {
	// Set config root
	configExists, err := configutil.SetDevSpaceRoot(log.GetInstance())
	if err != nil {
		return err
	}
	if !configExists {
		return errors.New("Couldn't find a DevSpace configuration. Please run `devspace init`")
	}

	// Get the config
	configOptions := cmd.ToConfigOptions()
	config, err := configutil.GetConfig(configOptions)
	if err != nil {
		return err
	}

	// Load generated config
	generatedConfig, err := generated.LoadConfig(cmd.Profile)
	if err != nil {
		return errors.Errorf("Error loading generated.yaml: %v", err)
	}

	err = dependency.UpdateAll(config, generatedConfig, cmd.AllowCyclicDependencies, configOptions, log.GetInstance())
	if err != nil {
		return err
	}

	log.Donef("Successfully updated all dependencies")
	return nil
}
