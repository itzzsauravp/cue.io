package cmd

import (
	"fmt"

	"github.com/itzzsauravp/go-rem/constants"
	"github.com/itzzsauravp/go-rem/helpers"
	"github.com/itzzsauravp/go-rem/internal/setup"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   constants.ROOT_CMD,
	Short: "`cue` Your personal reminder in CLI",
	Long: fmt.Sprintf(`%sVisit %s to view more details or contribute to it

                                                                                                      
 ██████╗██╗   ██╗███████╗                                                             
██╔════╝██║   ██║██╔════╝                                                             
██║     ██║   ██║█████╗                                                               
██║     ██║   ██║██╔══╝                                                               
╚██████╗╚██████╔╝███████╗                                                             
 ╚═════╝ ╚═════╝ ╚══════╝                                                             
                                                                                      
██████╗ ██╗   ██╗                                                                     
██╔══██╗╚██╗ ██╔╝                                                                     
██████╔╝ ╚████╔╝                                                                      
██╔══██╗  ╚██╔╝                                                                       
██████╔╝   ██║                                                                        
╚═════╝    ╚═╝                                                                        
                                                                                      
██╗████████╗███████╗███████╗███████╗ █████╗ ██╗   ██╗██████╗  █████╗ ██╗   ██╗██████╗ 
██║╚══██╔══╝╚══███╔╝╚══███╔╝██╔════╝██╔══██╗██║   ██║██╔══██╗██╔══██╗██║   ██║██╔══██╗
██║   ██║     ███╔╝   ███╔╝ ███████╗███████║██║   ██║██████╔╝███████║██║   ██║██████╔╝
██║   ██║    ███╔╝   ███╔╝  ╚════██║██╔══██║██║   ██║██╔══██╗██╔══██║╚██╗ ██╔╝██╔═══╝ 
██║   ██║   ███████╗███████╗███████║██║  ██║╚██████╔╝██║  ██║██║  ██║ ╚████╔╝ ██║     
╚═╝   ╚═╝   ╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝  ╚═══╝  ╚═╝     
                                                                                      

`, fmt.Sprintf("\n%s is a small CLI project created by me for fun and to learn Golang + Cobra\n", helpers.ColorCyan("cue.io")), helpers.ColorCyan("github.com/itzzsauravp/cue.io")),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		helpers.PreRun(setup.FILE_PATH)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() error {
	return RootCmd.Execute()
}
