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
	cmd.AddCommand(NewWaitForPodsCmd(out))
	cmd.AddCommand(NewExecuteCmd(out))

	return cmd
}

type simpleBackupCmd struct {
	namespace  string
	selector   string
	container  string
	path       string
	dst        string
	parallel   int
	tag        string
	bufferSize float64

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
			if _, err := kubetasks.SimpleBackup(b.namespace, b.selector, b.container, b.path, b.dst, b.parallel, b.tag, b.bufferSize); err != nil {
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
	f.Float64VarP(&b.bufferSize, "buffer-size", "b", 6.75, "in-memory buffer size (in MB) to use for files copy (buffer per file_")

	return cmd
}

type waitForPods struct {
	namespace string
	selector  string
	container string
	replicas  int

	out io.Writer
}

// NewWaitForPodsCmd waits for a given number of replicas
func NewWaitForPodsCmd(out io.Writer) *cobra.Command {
	w := &waitForPods{out: out}

	cmd := &cobra.Command{
		Use:   "wait-for-pods",
		Short: "Wait for a given number of ready pods",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if err := kubetasks.WaitForPods(w.namespace, w.selector, w.replicas); err != nil {
				log.Fatal(err)
			}
		},
	}
	f := cmd.Flags()

	f.StringVarP(&w.namespace, "namespace", "n", "", "namespace to find pods")
	f.StringVarP(&w.selector, "selector", "l", "", "selector to filter on")
	f.IntVarP(&w.replicas, "replicas", "r", 1, "number of ready replicas to wait for")

	return cmd
}

type execCmd struct {
	namespace string
	selector  string
	container string
	command   string

	out io.Writer
}

// NewExecuteCmd executes a simple command in a container
func NewExecuteCmd(out io.Writer) *cobra.Command {
	e := &execCmd{out: out}

	cmd := &cobra.Command{
		Use:   "execute",
		Short: "Execute a command in a container. Only executes the command in the first pod",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if err := kubetasks.Execute(e.namespace, e.selector, e.container, e.command); err != nil {
				log.Fatal(err)
			}
		},
	}
	f := cmd.Flags()

	f.StringVarP(&e.namespace, "namespace", "n", "", "namespace to find pods")
	f.StringVarP(&e.selector, "selector", "l", "", "selector to filter on")
	f.StringVarP(&e.container, "container", "c", "", "container name to act on")
	f.StringVar(&e.command, "command", "", "command to execute in container")

	return cmd
}
