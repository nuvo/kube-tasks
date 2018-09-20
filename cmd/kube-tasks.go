package main

import (
	"io"
	"log"
	"os"

	"github.com/maorfr/kube-tasks/pkg/kubetasks"
	"github.com/maorfr/kube-tasks/pkg/utils"
	"github.com/spf13/cobra"
)

func main() {
	cmd := NewRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		log.Fatal("Failed to execute command")
	}
}

// NewRootCmd represents the base command when called without any subcommands
func NewRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kube-tasks",
		Short: "",
		Long:  ``,
	}

	out := cmd.OutOrStdout()

	cmd.AddCommand(NewSimpleBackupCmd(out))

	return cmd
}

type simpleBackupCmd struct {
	namespace string
	selector  string
	container string
	path      string
	dst       string
	parallel  int
	tag       string

	out io.Writer
}

// NewSimpleBackupCmd performs a backup
func NewSimpleBackupCmd(out io.Writer) *cobra.Command {
	b := &simpleBackupCmd{out: out}

	cmd := &cobra.Command{
		Use:   "simple-backup",
		Short: "backup files to S3",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := kubetasks.SimpleBackup(b.namespace, b.selector, b.container, b.path, b.dst, b.parallel, b.tag); err != nil {
				log.Fatal(err)
			}
		},
	}
	f := cmd.Flags()

	f.StringVarP(&b.namespace, "namespace", "n", "", "namespace to find pods")
	f.StringVarP(&b.selector, "selector", "l", "", "selector to filter on")
	f.StringVarP(&b.container, "container", "c", "", "container name to act on")
	f.StringVar(&b.path, "path", "", "path to act on")
	f.StringVar(&b.dst, "dst", "", "destination to backup to. Example: s3://bucket/backup")
	f.IntVarP(&b.parallel, "parallel", "p", 1, "number of files to copy in parallel. set this flag to 0 for full parallelism")
	f.StringVar(&b.tag, "tag", utils.GetTimeStamp(), "tag to backup to. Default is Now (yyMMddHHmmss)")

	return cmd
}
