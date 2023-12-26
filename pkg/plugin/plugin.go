package plugin

import (
	"context"
	"os"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	cmdlogs "k8s.io/kubectl/pkg/cmd/logs"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func ioStreams() genericiooptions.IOStreams {
	return genericiooptions.IOStreams{In: os.Stdin, Out: defaultWriter(), ErrOut: os.Stderr}
}

func RunPlugin(_ context.Context) error {
	cf := genericclioptions.NewConfigFlags(false)
	f := cmdutil.NewFactory(cf)
	cmdLog := cmdlogs.NewCmdLogs(f, ioStreams())

	cf.AddFlags(cmdLog.Flags())
	return cmdLog.Execute()
}
