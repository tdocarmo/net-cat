package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	client := &Client{
		conn: conn,
	}

	// Demander le nom du client
	welcomeMessage := "Welcome to TCP-Chat!\n" +
		"         _nnnn_\n" +
		"        dGGGGMMb\n" +
		"       @p~qp~~qMb\n" +
		"       M|@||@) M|\n" +
		"       @,----.JM|\n" +
		"      JS^\\__/  qKL\n" +
		"     dZP        qKRb\n" +
		"    dZP          qKKb\n" +
		"   fZP            SMMb\n" +
		"   HZM            MMMM\n" +
		"   FqM            MMMM\n" +
		" __| \".        |\\dS\"qML\n" +
		" |    `.       | `' \\Zq\n" +
		"_)      \\.___.,|     .'\n" +
		"\\____   )MMMMMP|   .'\n" +
		"     `-'       `--'\n"

	// Envoyer le message de bienvenue au client
	fmt.Fprintf(conn, "%s\n", welcomeMessage)

	// Demander le nom du client
	scanner := bufio.NewScanner(conn)
	for {
		fmt.Fprintf(conn, "[ENTER YOUR NAME]:")
		if scanner.Scan() {
			name := scanner.Text()
			if name != "" {
				client.name = name
				break
			}
		}
	}

	// Demander au client de choisir un canal parmi les canaux disponibles
	channelOptions := "Available channels:\n"
	for _, ch := range channels {
		channelOptions += fmt.Sprintf("- %s\n", ch.name)
	}
	channelOptions += "Choose a channel or create a new one (Type the channel name):\n"
	fmt.Fprintf(conn, "%s", channelOptions)

	// Lire la réponse du client
	var chosenChannel string
	if scanner.Scan() {
		chosenChannel = scanner.Text()
	}

	// Vérifier si le canal choisi existe, sinon le créer
	var foundChannel *Channel
	mu.Lock()
	for _, channel := range channels {
		if channel.name == chosenChannel {
			foundChannel = channel
			break
		}
	}
	if foundChannel == nil {
		foundChannel = &Channel{
			name:    chosenChannel,
			clients: []*Client{client}, // Ajouter le client actuel à la liste des clients du canal
		}
		channels = append(channels, foundChannel)
	} else {
		foundChannel.clients = append(foundChannel.clients, client) // Ajouter le client actuel au canal existant
	}
	mu.Unlock()
	client.canal = foundChannel.name

	// Envoyer l'historique des messages au client
	for _, msg := range chatBuffer {
		fmt.Fprintf(client.conn, "%s\n", msg)
	}

	// Diffuser le message de connexion aux autres clients du même canal
	broadcast(fmt.Sprintf("[%s] %s has joined the chat.", time.Now().Format("2006-01-02 15:04:05"), client.name))

	// Boucle pour lire les messages du client
	reader := bufio.NewReader(conn)
	for {
		// Lire une ligne de texte
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			// Retirer le client du canal et diffuser un message de départ
			mu.Lock()
			for _, channel := range channels {
				if channel.name == client.canal {
					for i, c := range channel.clients {
						if c == client {
							// Supprimer le client du canal
							channel.clients = append(channel.clients[:i], channel.clients[i+1:]...)
							break
						}
					}
					break
				}
			}
			mu.Unlock()
			broadcast(fmt.Sprintf("[%s] %s has left our chat.", time.Now().Format("2006-01-02 15:04:05"), client.name))
			return
		}

		// Ignorer les lignes vides
		if strings.TrimSpace(message) == "\n" || message == "" || strings.TrimSpace(message) == "" {
			continue
		}

		// Ajouter le message à l'historique des messages
		chatBuffer = append(chatBuffer, fmt.Sprintf("[%s] [%s]: %s", time.Now().Format("2006-01-02 15:04:05"), client.name, strings.TrimSpace(message)))

		// Diffuser le message à tous les clients du même canal
		broadcast(fmt.Sprintf("[%s] [%s]: %s", time.Now().Format("2006-01-02 15:04:05"), client.name, strings.TrimSpace(message)))
	}
}
