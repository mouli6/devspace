package list

import (
	"strconv"

	"github.com/devspace-cloud/devspace/pkg/devspace/config/configutil"
	"github.com/devspace-cloud/devspace/pkg/devspace/config/generated"
	"github.com/devspace-cloud/devspace/pkg/util/log"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

type profilesCmd struct{}

func newProfilesCmd() *cobra.Command {
	cmd := &profilesCmd{}

	profilesCmd := &cobra.Command{
		Use:   "profiles",
		Short: "Lists all DevSpace profiles",
		Long: `
#######################################################
############## devspace list profiles #################
#######################################################
Lists all DevSpace configuartions for this project
#######################################################
	`,
		Args: cobra.NoArgs,
		RunE: cmd.RunListProfiles,
	}

	return profilesCmd
}

// RunListProfiles runs the list profiles command logic
func (cmd *profilesCmd) RunListProfiles(cobraCmd *cobra.Command, args []string) error {
	// Set config root
	configExists, err := configutil.SetDevSpaceRoot(log.GetInstance())
	if err != nil {
		return err
	}
	if !configExists {
		return errors.New("Couldn't find a DevSpace configuration. Please run `devspace init`")
	}

	profiles, err := configutil.GetProfiles(".")
	if err != nil {
		return err
	}

	// Load generated config
	generatedConfig, err := generated.LoadConfig("")
	if err != nil {
		return err
	}

	// Specify the table column names
	headerColumnNames := []string{
		"Name",
		"Active",
	}

	configRows := make([][]string, 0, len(profiles))

	for _, profile := range profiles {
		configRows = append(configRows, []string{
			profile,
			strconv.FormatBool(profile == generatedConfig.ActiveProfile),
		})
	}

	log.PrintTable(log.GetInstance(), headerColumnNames, configRows)
	return nil
}
