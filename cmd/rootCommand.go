package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/AF250329/tfs2pdf/pkg/tfs2pdf"
	"github.com/spf13/cobra"
)

var pdfOutputFolder string

// addCmd represents the add command
var rootCommand = &cobra.Command{
	Use:   "tfs2pdf",
	Short: "Application will convert TFS item to PDF document",
	Long:  `Application will convert TFS item to PDF document`,
	Args: func(cmd *cobra.Command, args []string) error {

		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		// Cobra will remove the first argument (executable file name)
		if _, err := strconv.Atoi(args[0]); err != nil {
			return fmt.Errorf("could not convert provided TFS item ID to number. Provided value: %v", args[0])
		}

		return nil
	},
	Run: runApplication,
}

func runApplication(cmd *cobra.Command, args []string) {

	err := tfs2pdf.Run(args, pdfOutputFolder)
	if err != nil {
		// Error should be already printed
	}
}

func init() {
	pdfOutputFolder, _ := os.Executable()
	pdfOutputFolder = filepath.Dir(pdfOutputFolder)

	rootCommand.PersistentFlags().StringVar(&pdfOutputFolder, "output", pdfOutputFolder, "Output folder where PDF file will be created")
	// rootCommand.MarkPersistentFlagRequired("tfs-token")
}

func Execute() error {

	helpTemplate := helpTemplate()
	rootCommand.SetUsageTemplate(helpTemplate)

	return rootCommand.Execute()
}

func helpTemplate() string {
	// Slightly modified template from
	// https://github.com/spf13/cobra/blob/284f4101043c4b1faebde411caec35a7b6e36494/command.go#L539
	return `Usage:{{if .Runnable}}
  {{.UseLine}} <TFS item ID>{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}
Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}
Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}
Available Commands:{{range $cmds}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{else}}{{range $group := .Groups}}
{{.Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if not .AllChildCommandsHaveGroup}}
Additional Commands:{{range $cmds}}{{if (and (eq .GroupID "") (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}
Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}
Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}
