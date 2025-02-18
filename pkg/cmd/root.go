package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCommand - Creates new root command
func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "go-skeleton",
		Short: "GoSkeleton is a clean and minimal Go project template that provides a structured foundation for scalable applications. ",
		Long:  "GoSkeleton is a clean and minimal Go project template that provides a structured foundation for scalable applications. ",
	}
}
