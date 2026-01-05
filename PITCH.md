# PITCH

Application Mobile d'enregistrement et partage de données de localisations

L'application affiche la direction et la proximité de la localisation
L'application peut afficher une carte avec les localisations 
On peut partager une localisation dans des groupes 
- Plusieurs localisations peuvent etres partage dans un groupe
- Une localisation peut etre partager dans plusieurs groupes

Une utilisateur peut faire partie de plusieurs groupes
Il choissit quels localisations sont partager dans quels groupes


API:

Database: PostgreSQL
Tables:
- Users
- Locations
- Groups

- user-group
- location-group

Routes:

Users:
- POST /users        => créé un utilisateur     --- inscription
- GET /users         => pour debug
- GET /users/{id}    => récupéré 1 utilisateur  --- connexion
- PUT /users/{id}    => modifier un utilisateur --- parametre compte 
- DELETE /users/{id} => supprime un utilisateur --- parametre le compte

- /login     => recupération de token
- /refresh   => refresh du token

- GET /users/{id}/locations  => recupère toutes les localisations d'un utilisateur              --- compte/page principale
- GET /users/{id}/groups     => recupère tous les groupes dans lesquels se trouve l'utilisateur --- page secondaire groupes

Locations:
- POST /locations         => créé une localisation                    --- page principale => + localisation
- GET /locations          => pour debug
- GET /locations/{id}     => récupère une localisation en particulier --- page principale ou page groupe vers page localisation
- PUT /locations/{id}     => modifie une localisation                 --- page localisation
- DELETE /locations/{id}  => supprime une localisation                --- page localisation

- GET /locations/{id}/groups => récupère tous les groupes où la localisation est partager --- page localisation

Groups:
- POST /groups => créé un groupe --- page secondaire groupes
- GET /groups => pour debug
- GET /groups/{id} => récupère un groupe --- page secondaires groupes vers page groupe
- PUT /groups/{id} => modifie un groupe --- page groupe vision admin
- DELETE /groups/{id} => supprime un groupe --- page groupe vision admin

---
- POST /groups/{id}/users ({"user_id" : id}) => ajoute un utilisateur a un groupe --- page group => + user
- GET /groups/{id}/users => recupère tous les utilisateurs d'un groupes --- page groupe vers section utilisateurs
- DELETE /groups/{id}/users/{id} => supprime un utilisateur d'un groupe --- part d'un groupe/ban du groupe vision admin

---
- POST /groups/{id}/locations ({"location_id" : id})    => partage d'une localisattions dans un groupe --- page localisation/page groupe boutton partage
- GET /groups/{id}/locations => recupère toutes les localisations partager dans le groupes --- page groupe
- PUT /groups/{id}/locations/{id} => modifie le partage de localisation --- page groupe vision partageur
- DELETE /groups/{id}/locations/{id} => supprime la localisation du groupe --- page groupe vision partageur