package main

import (
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"golang.org/x/oauth2"
)

func main() {
	// Credenciais da API do Dropbox
	config := &oauth2.Config{
		ClientID:     "seu_client_id",
		ClientSecret: "seu_client_secret",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.dropbox.com/oauth2/authorize",
			TokenURL: "https://api.dropboxapi.com/oauth2/token",
		},
	}

	// Token de acesso do usuário
	token := oauth2.Token{
		AccessToken: "seu_access_token",
	}

	// Cria o cliente da API do Dropbox com as credenciais e token de acesso
	dbx := files.New(dropbox.Config{
		Token: token.AccessToken,
	})

	// Abre o arquivo que será enviado
	file, err := os.Open("caminho/do/arquivo")
	if err != nil {
		// tratamento de erro
	}
	defer file.Close()

	// Upload do arquivo para o Dropbox
	_, err = dbx.Upload(&files.CommitInfo{
		Path: "/caminho/no/dropbox/nome_do_arquivo",
	}, file)
	if err != nil {
		// tratamento de erro
	}
}
