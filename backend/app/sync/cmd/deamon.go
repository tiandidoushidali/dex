package cmd

import (
	"context"
	_ "dex/app/sync/internal/config"
	"dex/app/sync/internal/service"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var deamonCmd = &cobra.Command{
	Use:   "deamon",
	Short: "deamon commands",
	Long:  "deamon commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deamon called")
		var license, _ = cmd.Flags().GetString("license")
		fmt.Println("flag license:", license)
		//fmt.Println(cmd.Usage())

		ctx, cancel := context.WithCancel(context.Background())
		syncExit := make(chan os.Signal)
		signal.Notify(syncExit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)

		var wg sync.WaitGroup
		go func() {
			wg.Add(1)
			defer wg.Done()

			s := service.New()
			s.Run(ctx)
		}()

		select {
		case <-syncExit:
			fmt.Println("deamon exit")
			cancel()
		}
		wg.Wait()
		fmt.Println("deamon finish")
	},
}

func init() {
	rootCmd.AddCommand(deamonCmd)
}
