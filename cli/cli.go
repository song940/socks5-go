package cli

import "github.com/song940/socks5/socks5"

func Run() {

	cred := socks5.DemoCredentials{}
	cator := socks5.UserPassAuthenticator{Credentials: cred}

	// Create a SOCKS5 server
	conf := &socks5.Config{AuthMethods: []socks5.Authenticator{cator}}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "127.0.0.1:1080"); err != nil {
		panic(err)
	}
}
