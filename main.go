package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"resource-service/dbmodel"
	"resource-service/src/controller"
	applogger "resource-service/utils/logging"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

/*
********************
 RESOURCE MGMNT APP
********************
*/

var log *zerolog.Logger = applogger.GetInstance()
var configFilePath *string
var cfg dbmodel.Config

type application struct {
	config dbmodel.Config
	logger *zerolog.Logger
}

func main() {

	gin.DisableConsoleColor()
	r := gin.New()
	//Allow CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	app := &application{
		config: cfg,
		logger: log,
	}

	db := dbmodel.SetupDB()
	db.AutoMigrate()

	setupLogger(r)
	setupRoutes(r)
	startServer(r, app)
}

// startServer - Start server
func startServer(r *gin.Engine, a *application) {

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("Listen: %s\n", err)
	}

	log.Info().Msg("Shutting down server...\n")

	log.Info().Msg("Server exiting\n")

}

// setupLogger - Configure logging for the server
func setupLogger(r *gin.Engine) {
	// Configure logger
	zlog := applogger.GetInstance()
	r.Use(logger.SetLogger(logger.Config{
		Logger: zlog,
		UTC:    true,
	}))
}

//setupRoutes - Define all the Routes
func setupRoutes(r *gin.Engine) {
	// Instantiate controllers
	var s controller.ResourceController
	var l controller.LoginController

	http.HandleFunc("/resource/get-resource", getResc)
	r.POST("/resource/login-user", l.LoginUser)
	r.POST("/resource/add-resc", s.AddResource)
	r.GET("/resource/get-resc", s.GetResource)
	r.GET("/resource/get-resc-links", s.Getfile_imgLink)

}

// get Resource handler to handle the route
func getResc(w http.ResponseWriter, req *http.Request) {
	a, _ := strconv.Atoi(req.URL.Query().Get("id"))

	str_data := map[string]string{
		"title":       "civo",
		"category":    "WhiteP",
		"status":      "Published",
		"type":        "PDF",
		"image_links": "hyt.jpg",
		"file_link":   "gtw.pdf",
		"content":     "GR",
	}
	int_data := map[string]int32{
		"id":         1,
		"created_by": 68,
	}

	resc_str_data := getData(str_data)
	resc_int_data := getData(int_data)

	if a == 1 {
		fmt.Fprint(w, resc_str_data, resc_int_data)
	} else {
		fmt.Fprint(w, "Wrong Id caught")
	}

}

// use generic func to provide smooth response
func getData[V compararble, b int32 | float64 | string](Z map[V]b) []b{

	var x_data []b
	for _, v := range Z {
		x_data = append(x_data, v)
	}
	return x_data
}
