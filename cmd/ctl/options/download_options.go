package options

import (
	"github.com/spf13/cobra"
)

type CliDownloadOptions struct {
	Version  string
	Md5sum   string
	Manifest string
	BaseDir  string
}

func NewCliDownloadOptions() *CliDownloadOptions {
	return &CliDownloadOptions{}
}

func (o *CliDownloadOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.Version, "version", "", "Set install-wizard version, e.g., 1.7.0, 1.7.0-rc.1, 1.8.0-20240813")
	cmd.Flags().StringVar(&o.Md5sum, "md5sum", "", "Set install-wizard package md5 sum")
	cmd.Flags().StringVar(&o.Manifest, "manifest", "", "Set download package manifest file , default value $HOME/.terminus/installation.manifest")
	cmd.Flags().StringVar(&o.BaseDir, "base-dir", "", "Set download package base dir , default value $HOME/.terminus")
}
