**README - TCP Chat**

**Introduction**
Le programme TCP Chat est une application de chat basée sur le protocole TCP. Il permet à plusieurs utilisateurs de communiquer entre eux via un réseau.

**Fonctionnalités**
- Connexion à un port spécifique pour l'écoute des connexions entrantes.
- Gestion des canaux de discussion.
- Enregistrement des messages dans un historique.
- Conservation des logs des informations de connexion.

**Installation**
1. Assurez-vous d'avoir Go installé sur votre système.
2. Clonez ce dépôt Git ou téléchargez les fichiers du programme.
3. Exécutez la commande suivante pour compiler le programme :
go build -o TCPChat main.go

markdown
Copy code
4. Lancez le programme en spécifiant le port sur lequel écouter les connexions :
./TCPChat <port>

markdown
Copy code

**Utilisation**
- Lors du lancement de l'application, les utilisateurs peuvent se connecter en spécifiant un nom d'utilisateur.
- Les utilisateurs peuvent choisir un canal de discussion parmi ceux disponibles ou en créer un nouveau.
- Une fois connectés, les utilisateurs peuvent envoyer et recevoir des messages dans le canal choisi.
- L'historique des messages est automatiquement conservé par le programme pour référence ultérieure.
- Les logs des informations de connexion sont automatiquement conservés par le programme pour référence ultérieure.

**Contributions**
Les contributions sont les bienvenues ! Si vous souhaitez contribuer à ce projet, n'hésitez pas à ouvrir une issue ou à soumettre une pull request.

**Auteurs**
Ce programme a été développé par Toni Do Carmo, Baran Aksoy, Martin Depreau.

**Licence**
Ce projet est sous licence MIT. Consultez le fichier LICENSE.md pour plus de détails.
