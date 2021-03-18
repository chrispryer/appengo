package cmd

import (
	"fmt"
	"os"

	"github.com/cnpls/appengo/utils"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "appengo",
	Short: "appengo command line tool",
	Long:  `Append files from the command line`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := Run(args)
		if err != nil {
			panic(err) // TODO: look into non-panic
		}
	},
}

func Run(args []string) error {
	fmt.Printf("Targeted files: %s", args)

	for i, filepath := range args {
		fmt.Printf("\nLoading file (%d, %s)", i, filepath)
		readFile(filepath)
	}

	return nil
}

// TODO: return []byte
func readFile(filepath string) {
	// TODO
	// 1. does this open for read-only?
	// 2. if so is there a way to get fileinfo from metadata?
	// ^ would reduce the number of times you're loading the file.
	file, err := os.Open(filepath)
	check(err)
	defer file.Close()

	fileinfo, err := file.Stat()
	check(err)

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	nbytes, err := file.Read(buffer)
	check(err)
	fmt.Printf("\n%d bytes read from %s", nbytes, filepath)
}

func check(err error) {
	utils.Check(err)
}
