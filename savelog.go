package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func savelog(remoteAddr net.Addr) {
	// Obtenir l'heure et la date actuelles
	now := time.Now()

	// Formater l'heure et la date
	currentTime := now.Format("2006-01-02 15:04:05")

	// Ouvrir le fichier de log en mode écriture, avec ajout et création s'il n'existe pas
	file, err := os.OpenFile("logfile.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier de logs:", err)
		return
	}
	defer file.Close()

	// Écrire dans le fichier de log
	_, err = fmt.Fprintf(file, "Date et heure: %s - Adresse IP: %s\n", currentTime, remoteAddr)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier de logs:", err)
		return
	}
	fmt.Println("Informations de connexion enregistrées dans le fichier.")
}
