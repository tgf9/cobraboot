package command

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/tgf9/cobraboot/internal/command/foo"
)

func Run() error {
	cmd := &cobra.Command{
		Use:  "cobraboot",
		RunE: cmdMain,

		// Silence usage when an error occurs.
		SilenceUsage: true,
		// Take control of reporting errors.
		SilenceErrors: true,
	}
	// Remove help subcommand.
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	// Remove completion subcommand.
	cmd.CompletionOptions.DisableDefaultCmd = true

	addFlags(cmd.Flags())

	cmd.AddCommand(foo.Command())

	return cmd.Execute()
}

func addFlags(f *pflag.FlagSet) {
	f.BoolP("help", "h", false, "Show help")
	f.StringP("str", "s", "", "Set str")
}

func cmdMain(cmd *cobra.Command, args []string) error {
	f := cmd.Flags()

	if ok, err := f.GetBool("help"); err != nil && ok {
		return cmd.Usage()
	}

	str, err := f.GetString("str")
	if err != nil {
		return err
	}
	log.Printf("str=%q", str)

	return nil
}
