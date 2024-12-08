package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Swapnil1296/crud-tutorial/internal/config"
)

// go run cmd/students-api/main.go
// go run cmd/students-api/main.go -config config/local.yaml
func main() {

	

	// laod config 
	cfg :=config.MustLoad()

//database setup
//setup router
router:=http.NewServeMux()
router.HandleFunc("GET /",func(w http.ResponseWriter, r*http.Request){
	w.Write([]byte("welcome to crud operations"))
})
//setup server
server:=http.Server{
	Addr:cfg.Addr,
	Handler:router,


}
slog.Info("server started",slog.String("address",cfg.Addr))
fmt.Printf("server started %s",cfg.HTTPServer.Addr)
//creating a  channel 
done:= make(chan os.Signal,1)
signal.Notify(done,os.Interrupt,syscall.SIGINT, syscall.SIGTERM)

// gracefull shutdown if any interuption error occurs
go func(){
	err:=server.ListenAndServe()
if err != nil{
	log.Fatal("failed to start server")
}
}()
<-done
slog.Info("shutting down the server")
ctx, cancel:= context.WithTimeout(context.Background(),5*time.Second)
defer cancel()

err:=server.Shutdown(ctx)
if err != nil{
	slog.Error("failed to shutdown server", slog.String("error",err.Error()))
}

slog.Info("server shutdown successfully")
}