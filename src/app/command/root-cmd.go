/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package command

const (
	AppEmoji        = "🦄"
	ApplicationName = "gummy"
	RootPsName      = "root-ps"
	SourceID        = "github.com/plastikfan/gummy"
)

func Execute() error {
	return (&Bootstrap{}).Root().Execute()
}

// CLIENT-TODO: define valid properties on the root parameter set
type RootParameterSet struct {
	Language string
}
