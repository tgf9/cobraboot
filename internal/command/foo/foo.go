package foo

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "foo",
		RunE: cmdMain,

		// Silence usage when an error occurs.
		SilenceUsage: true,
		// Take control of reporting errors.
		SilenceErrors: true,
	}

	addFlags(cmd.Flags())

	return cmd
}

func addFlags(f *pflag.FlagSet) {
	f.BoolP("help", "h", false, "Show help")
	f.StringP("foostr", "s", "", "Set foostr")
}

func cmdMain(cmd *cobra.Command, args []string) error {
	f := cmd.Flags()

	if ok, err := f.GetBool("help"); err != nil && ok {
		return cmd.Usage()
	}

	fooStr, err := f.GetString("foostr")
	if err != nil {
		return err
	}
	log.Printf("foostr=%q", fooStr)

	return nil
}
