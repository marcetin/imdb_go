package actions

import (
	"fmt"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/cookiestore"
)

func init() {
	gothic.Store = App().SessionStore

	goth.UseProviders(
		cookiestore.New(os.Getenv("COOKIESTORE_KEY"), os.Getenv("COOKIESTORE_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/cookiestore/callback")),
	)
}

func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	// Do something with the user, maybe register them/sign them in
	return c.Render(200, r.JSON(user))
}
