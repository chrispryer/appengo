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

	files := make([]byte, len(args))

	for i, filepath := range args {
		fmt.Printf("\nLoading file (%d, %s)", i, filepath)
		file := readFile(filepath)

		files = append(files, file...)
	}

	fmt.Printf("\nResult: %s", files)

	return nil
}

func readFile(filepath string) []byte {
	// TODO
	// 1. does this open for read-only?
	// 2. if so is there a way to get fileinfo from metadata?
	// ^ would reduce the number of times you're loading the file.
	file, err := os.Open(filepath)
	utils.Check(err)
	defer file.Close()

	fileinfo, err := file.Stat()
	utils.Check(err)

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	nbytes, err := file.Read(buffer)
	utils.Check(err)
	fmt.Printf("\n%d bytes read from %s", nbytes, filepath)

	return buffer
}
