package main

import "fmt"

// Diffuser un message à tous les clients, sauf à l'expéditeur, dans le canal spécifié
func broadcast(message string) {
    for _, channel := range channels {
        for _, client := range channel.clients {
            fmt.Fprintf(client.conn, "%s\n", message)
        }
    }
}
