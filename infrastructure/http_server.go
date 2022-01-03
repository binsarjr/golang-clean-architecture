package infrastructure

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"giapps/servisin/application"
	"giapps/servisin/domain/model"
	"giapps/servisin/domain/repository"
	"giapps/servisin/infrastructure/config"
	"giapps/servisin/infrastructure/database"
	"giapps/servisin/infrastructure/exception"
	"giapps/servisin/interfaces"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v4/pgxpool"
)

// HTTPServer define an application structure
type HTTPServer struct {
	router        chi.Router
	configuration config.Config
	database      *pgxpool.Pool
	tokenAuth     *jwtauth.JWTAuth
}

// Start run the application
func (app *HTTPServer) Start() {
	app.HandleNotFound()
	app.HandleMethodNotAllowed()
	app.MiddlewareRegistry()

	// Setup Repository
	userRepository := repository.NewUserRepository(app.database)

	// Setup Application Services
	authAppInterface := application.NewAuthenticateAppInterface(&userRepository, app.tokenAuth)

	// Setup Interface
	authInterfaces := interfaces.NewAuthenticate(&authAppInterface, &userRepository, app.tokenAuth)

	app.router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verify(app.tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)
	})

	// authentication routes
	app.router.Post("/login", authInterfaces.Login)
	app.router.Post("/register", authInterfaces.Register)

	app.Serve(app.configuration.Get("APP_PORT"))
}

// NewHTTPServer creates new HTTPServer with its dependencies
func NewHTTPServer() *HTTPServer {
	configuration := config.New()

	pool_max, err := strconv.Atoi(configuration.GetOrDefault("POSTGRESS_POOL_MAX", "5"))
	exception.PanicIfNeeded(err)
	pool_min, err := strconv.Atoi(configuration.GetOrDefault("POSTGRESS_POOL_MIN", "2"))
	exception.PanicIfNeeded(err)

	dbport, err := strconv.Atoi(configuration.Get("POSTGRESS_PORT"))
	exception.PanicIfNeeded(err)

	return &HTTPServer{
		router:        chi.NewRouter(),
		configuration: configuration,
		database: database.NewPostgres(database.Postgres{
			MaxConns: int32(pool_max),
			MinConns: int32(pool_min),
			Connection: database.PostgresConnection{
				Host:     configuration.Get("POSTGRESS_HOST"),
				User:     configuration.Get("POSTGRESS_USER"),
				Password: configuration.Get("POSTGRESS_PASSWORD"),
				Database: configuration.Get("POSTGRESS_DATABASE"),
				Port:     uint16(dbport),
			},
		}),
		tokenAuth: jwtauth.New("HS256", []byte(configuration.Get("JWT_SECRETKEY")), nil),
	}
}

// HandleNotFound
func (app *HTTPServer) HandleNotFound() {
	app.router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &exception.ErrResponse{Code: http.StatusNotFound, Message: "404 not found"})
	})
}

// MethodNotAllowed
func (app *HTTPServer) HandleMethodNotAllowed() {
	app.router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &exception.ErrResponse{Code: http.StatusMethodNotAllowed, Message: "method tidak diijinkan"})
	})
}

// Error handler
func (app *HTTPServer) ErrorHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				err, ok := rvr.(exception.ErrResponse)
				if ok {
					render.Render(w, r, &err)
					return
				}

				render.Render(w, r, &model.WebResponse{
					Code:   http.StatusInternalServerError,
					Status: "INTERNAL_SERVER_ERROR",
					Data:   rvr,
				})
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// Setup middleware
func (app *HTTPServer) MiddlewareRegistry() {
	app.router.Use(middleware.RealIP)
	app.router.Use(render.SetContentType(render.ContentTypeJSON))
	app.router.Use(middleware.Logger)
	app.router.Use(app.ErrorHandler)
	app.router.Use(middleware.RequestID)
}

func (app *HTTPServer) Serve(port string) {

	addr := fmt.Sprintf(":%s", port)
	server := &http.Server{
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
		Addr:         addr,
		Handler:      app.router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
