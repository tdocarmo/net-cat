package main

import (
	"net"
	"sync"
)

type Client struct {
	conn  net.Conn
	name  string
	canal string // Ajout de la propriété "canal" pour le client
}

type Channel struct {
	name    string
	clients []*Client
}

var (
	channels   []*Channel // Liste des canaux disponibles
	chatBuffer []string   // Historique des messages
	maxConnections = 10 // Limite de connexions
	sem            = make(chan struct{}, maxConnections)
	mu             sync.Mutex // Mutex pour protéger l'accès concurrent à la liste des canaux
)
