package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// grepcloneCmd represents the grepclone command
var grepcloneCmd = &cobra.Command{
	Use:   "grepclone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grepclone called")
		filename, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Printf("Couldn't get file: %v", err)
		}
		sTerm, err := cmd.Flags().GetString("sterm")
		if err != nil {
			fmt.Printf("Couldn't get string: %v", err)
		}
		res, err := searchFile(filename, sTerm, false)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(grepcloneCmd)
	grepcloneCmd.Flags().StringP("file", "f", "", "Filename | Path to a file")
	grepcloneCmd.Flags().StringP("sterm", "s", "", "Search Term")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grepcloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grepcloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// not = true is like -v flag of grep
func searchFile(path, sTerm string, not bool) (string, error) {
	scanner, err := openFile(path)
	if err != nil {
		return "", err
	}
	var res []string
	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for scanner.Scan() {
		// if the search term is found on the current line, append it to the resulting slice
		if not {
			if !strings.Contains(scanner.Text(), sTerm) {
				res = append(res, scanner.Text())
			}
		} else {
			if strings.Contains(scanner.Text(), sTerm) {
				res = append(res, scanner.Text())
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if len(res) < 1 {
		return "", errors.New("nothing found by that search term")
	}
	return buildStrFromSlice(res), nil
}

func openFile(path string) (*bufio.Scanner, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(f), nil
}

// Build response as a single string from a slice of strings
func buildStrFromSlice(ss []string) string {
	var sb strings.Builder
	for _, str := range ss {
		sb.WriteString(str)
		sb.WriteString("\n")
	}
	return sb.String()
}