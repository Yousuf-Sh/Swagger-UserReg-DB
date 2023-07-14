// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"SwaggerPostGresApi/models"
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"SwaggerPostGresApi/restapi/operations"
	"SwaggerPostGresApi/restapi/operations/users"
)

//go:generate swagger generate server --target ../../Swagger-PostGreSQL-Api --name UserRegistrationAPI --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.UserRegistrationAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UserRegistrationAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.UsersCreateUserHandler == nil {
		api.UsersCreateUserHandler = users.CreateUserHandlerFunc(func(params users.CreateUserParams) middleware.Responder {
			user := &models.User{
				Name:  params.User.Name,
				Email: params.User.Email,
			}
			err := CreateUser(db, user)
			if err != nil {
				return users.NewCreateUserInternalServerError()
			}
			return users.NewCreateUserOK().WithPayload(user)
		})
	}
	if api.UsersDeleteUserByIDHandler == nil {
		api.UsersDeleteUserByIDHandler = users.DeleteUserByIDHandlerFunc(func(params users.DeleteUserByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.DeleteUserByID has not yet been implemented")
		})
	}
	if api.UsersGetAllUsersHandler == nil {
		api.UsersGetAllUsersHandler = users.GetAllUsersHandlerFunc(func(params users.GetAllUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetAllUsers has not yet been implemented")
		})
	}
	if api.UsersGetUserByIDHandler == nil {
		api.UsersGetUserByIDHandler = users.GetUserByIDHandlerFunc(func(params users.GetUserByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUserByID has not yet been implemented")
		})
	}
	if api.UsersUpdateUserByIDHandler == nil {
		api.UsersUpdateUserByIDHandler = users.UpdateUserByIDHandlerFunc(func(params users.UpdateUserByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.UpdateUserByID has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func CreateUser(db *sql.DB, user *models.User) error {
	stmt, err := db.Prepare("INSERT INTO users(name,email) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("Failed to prepare statementt: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}
	return nil
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
