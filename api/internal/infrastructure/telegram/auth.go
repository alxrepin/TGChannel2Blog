package telegram

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
)

type Authenticator struct{}

func (a Authenticator) Code(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
	fmt.Print("Enter code: ")

	var code string

	fmt.Scanln(&code)

	return code, nil
}

func (a Authenticator) Password(ctx context.Context) (string, error) {
	fmt.Print("Enter 2FA password: ")

	reader := bufio.NewReader(os.Stdin)
	pass, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	pass = strings.TrimSpace(pass)

	return pass, nil
}

func NewAuthenticator() auth.Flow {
	return auth.NewFlow(
		auth.CodeOnly("", Authenticator{}),
		auth.SendCodeOptions{},
	)
}
