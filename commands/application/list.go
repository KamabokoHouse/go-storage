package application

import (
	"fmt"
	"go-storage/commands/domain/list"

	"github.com/spf13/cobra"
)

var (
	listObjects = &cobra.Command{
		Use:  "listObjects",
		RunE: action,
	}

	option = &ListOption{}
)

//ListOption ...
type ListOption struct {
	bucketName string
}

func action(cmd *cobra.Command, args []string) error {
	fmt.Println("Surching bucket in: " + option.bucketName)
	list := list.NewList(option.bucketName)
	err := list.Display()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	RootCmd.AddCommand(listObjects)
	listObjects.Flags().StringVarP(&option.bucketName, "bucket", "b", "default", "string option")
}
