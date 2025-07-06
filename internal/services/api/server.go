package api

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/kiraralab/snowbankcp/ui/webpages"
)

// Thanks to otraore for the code example
// https://gist.github.com/otraore/4b3120aa70e1c1aa33ba78e886bb54f3

const (
	userkey = "user"   // key used to store the username in the session
	secret  = "secret" // random and secure key used to encrypt the session cookie
)

func StartServer() {
	// Initialize the engine
	e, err := engine()
	if err != nil {
		log.Fatal("Unable to initialize engine:", err)
	}

	// Run the engine on port 8080
	if err := e.Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func engine() (*gin.Engine, error) {
	// Create a new gin engine
	r := gin.New()

	// Enable Gin's logging middleware for request logging
	r.Use(gin.Logger())

	// Setup the cookie store for session management
	// This middleware will automatically handle session cookie reading/writing
	// r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte(secret))))

	fs, err := static.EmbedFolder(webpages.WebPageFS, "build/client")
	if err != nil {
		return nil, err
	}
	r.Use(static.Serve("/", fs))

	// Public routes that don't require authentication
	r.POST("/login", login)  // Handles user login
	r.GET("/logout", logout) // Handles user logout

	// Private route group, protected by AuthRequired middleware
	// All routes within this group require a valid session
	private := r.Group("/private")
	private.Use(AuthRequired) // Enable the middleware on these routes
	{
		private.GET("/me", me)         // Returns current user info
		private.GET("/status", status) // Returns login status
	}

	return r, nil
}
