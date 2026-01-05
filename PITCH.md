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
- POST /users
- GET /users
- GET /users/{id}
- PUT /users/{id}
- DELETE /users/{id}

- /login
- /refresh

- GET /users/{id}/locations
- GET /users/{id}/groups

Locations:
- POST /locations
- GET /locations
- GET /locations/{id}
- PUT /locations/{id}
- DELETE /locations/{id}

- GET /locations/{id}/groups

Groups:
- POST /groups
- GET /groups
- GET /groups/{id}
- PUT /groups/{id}
- DELETE /groups/{id}

--- 
- POST /groups/{id}/users ({"user_id" : id})
- GET /groups/{id}/users
- DELETE /groups/{id}/users/{id}

---
- POST /groups/{id}/locations ({"location_id" : id})
- GET /groups/{id}/locations
- PUT /groups/{id}/locations/{id} 
- DELETE /groups/{id}/locations/{id}