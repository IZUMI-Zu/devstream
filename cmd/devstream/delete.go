package main

import (
	"os"

	"github.com/devstream-io/devstream/internal/pkg/configloader"

	"github.com/spf13/cobra"

	"github.com/devstream-io/devstream/internal/pkg/pluginengine"
	"github.com/devstream-io/devstream/pkg/util/log"
)

var deleteCMD = &cobra.Command{
	Use:   "delete",
	Short: "Delete DevOps tools according to DevStream configuration file",
	Long: `Delete DevOps tools according to DevStream configuration file. 
DevStream will delete everything defined in the config file, regardless of the state.`,
	Run: deleteCMDFunc,
}

func deleteCMDFunc(cmd *cobra.Command, args []string) {
	log.Info("Delete started.")

	gConfig, err := configloader.LoadGeneralConf(configFile)
	if err != nil {
		log.Errorf("Delete error: %s.", err)
		os.Exit(1)
	}
	log.Debugf("config file content is %s.", gConfig)

	if err := pluginengine.Remove(gConfig.ToolFile, gConfig.VarFile, continueDirectly, isForceDelete); err != nil {
		log.Errorf("Delete error: %s.", err)
		os.Exit(1)
	}

	log.Success("Delete finished.")
}

func init() {
	deleteCMD.PersistentFlags().BoolVarP(&isForceDelete, "force", "", false, "force delete by config")
}
