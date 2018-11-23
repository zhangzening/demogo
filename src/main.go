package main

import (
	"demogo/src/inter"
	"demogo/src/model"
	"demogo/src/router"
	"demogo/src/service"
	"demogo/src/store"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:               "demogo",
		Short:             "demogo server",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Long:              `Reboot demogo server`,
		PersistentPreRunE: func(*cobra.Command, []string) error {
			fmt.Println("----------")
			return nil
		},
	}
	startCmd = &cobra.Command{
		Use:     "start",
		Short:   "Start demogo server",
		Aliases: []string{"startAt"},
		RunE:    start,
	}
)

func init() {
	startCmd.PersistentFlags().StringVarP(&model.Host.Port, "port", "p", "8081", "port agrs")

	startCmd.PersistentFlags().StringVarP(&model.Host.Address, "address", "H", "127.0.0.1", "address args")

	rootCmd.AddCommand(startCmd)
}

func start(_ *cobra.Command, _ []string) error {
	store.Init()
	service.Init()

	engine := gin.Default()
	router.Init(engine)
	server := &http.Server{
		//Addr:    fmt.Sprintf("%s:%s", &model.Host.Address, &model.Host.Port),
		Addr:    ":8081",
		Handler: engine,
	}

	server.ListenAndServe()

	//engine.Run(":8081")

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}

	// test interface
	//testInterface()
}

func testInterface() {
	inter.TestStruct1()
	inter.TestStruct2()
	inter.TestFunc1()
	inter.TestFunc2()
}
