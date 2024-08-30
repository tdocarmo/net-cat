package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var port string

	// Vérifier les arguments de la ligne de commande
	if len(os.Args) == 1 {
		port = "8989"
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(1)
	} else {
		port = os.Args[1]
		// Vérifier si le port est un entier valide
		if !isValidPort(port) {
			fmt.Println("[USAGE]: ./TCPChat $port")
			os.Exit(1)
		}
	}

	// Initialiser quelques canaux par défaut
	channels = append(channels, &Channel{name: "General"})
	channels = append(channels, &Channel{name: "Random"})
	// Vous pouvez ajouter d'autres canaux si nécessaire

	// Écoute sur le port spécifié
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Printf("Listening on the port: %s\n", port)

	// Boucle pour accepter les connexions entrantes
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Acquérir un jeton du sémaphore
		sem <- struct{}{}

		// Création d'une goroutine pour gérer la connexion du client
		go func(conn net.Conn) {
			defer func() {
				// Libérer un jeton du sémaphore
				<-sem
			}()

			// Enregistrer les informations de connexion avec l'adresse IP de l'hôte
		savelog(conn.RemoteAddr())

			handleClient(conn)
		}(conn)
	}
}
