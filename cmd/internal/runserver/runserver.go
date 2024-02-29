package runserver

import (
	"context"
	"fmt"
	"github.com/atpuxiner/grapi/app/initializer/conf"
	"github.com/atpuxiner/grapi/app/initializer/db"
	"github.com/atpuxiner/grapi/app/initializer/logger"
	"github.com/atpuxiner/grapi/app/router"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type params struct {
	port string
}

type handler struct {
	cmd  *cobra.Command
	args []string
	pars params
}

func (r *handler) init() {
	conf.InitConfig()
	logger.InitLogger()
	defer func() {
		if err := logger.Logger.Sync(); err != nil {
			log.Printf("Sync log error: %v", err)
		}
	}()
	db.InitDB(conf.Conf.DB.Driver)

	r.pars.port, _ = r.cmd.Flags().GetString("port")
	if r.pars.port == "" {
		r.pars.port = conf.Conf.Server.Port
		if r.pars.port == "" {
			r.pars.port = "9000"
		}
	}
}

func (r *handler) handle() {
	r.start()
	defer r.clean()
}

func (r *handler) start() {
	c, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	engine := router.InitRouter()
	srv := &http.Server{
		Addr:    ":" + r.pars.port,
		Handler: engine,
	}
	go func() {
		fmt.Println("")
		fmt.Println(time.Now())
		fmt.Printf("Starting server at http://127.0.0.1:%s/\n", r.pars.port)
		fmt.Printf("Testing server at http://127.0.0.1:%s/ping\n", r.pars.port)
		fmt.Println("Quit the server with Ctrl+C")
		fmt.Println("")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server start error: %v", err)
		}
	}()

	<-c.Done()

	stop()

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(c); err != nil {
		log.Fatalf("Server stop error: %v", err)
	}
}

func (r *handler) clean() {
	fmt.Println("")
	fmt.Println(time.Now())
	fmt.Println("Exiting the server.....")
	// TODO:
}

func Run(cmd *cobra.Command, args []string) {
	h := handler{
		cmd:  cmd,
		args: args,
	}
	h.init()
	h.handle()
}
