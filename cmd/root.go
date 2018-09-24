package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tkmru/cibo/core"
)

const (
	version = "dev"
)

var debugFlag bool
var beginAddress int
var bitMode int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cibo",
	Short: "x86 CPU emulator",
	Long:  `cibo - x86 CPU emulator`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		var path string
		path = args[0]

		emu := cibo.NewEmulatorWithLoadFile(bitMode, beginAddress, path, debugFlag)
		cpu := emu.CPU
		reg := &cpu.X86registers

		reg.Init()
		emu.Run()
		reg.Dump()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute(path []string) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().IntVarP(&beginAddress, "address", "a", 0x7c00, "begin address")
	rootCmd.Flags().IntVarP(&bitMode, "bit", "b", 32, "bit mode")
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "debug mode")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cibo version %s\n", version)
	},
}
