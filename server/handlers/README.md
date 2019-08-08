# API Handlers

# API Docs (Kinda)


- GET - `/api/auth/google` - Redirects the client to Google-OAuth2 authentication flow.

- GET - `/api/auth/google/callback` - Google-OAuth2 callback handler.

- GET - `/api/auth/current_user` - To fetch data about currently logged in user/admin.



- POST -`/api/post/new`

- GET -`/api/post/all`

- GET -`/api/post/single?id=123`

- PUT -`/api/post/edit?id=123`

- DELETE -`/api/post/delete/temp?id=123`

- DELETE -`/api/post/delete/perm?id=123`



- GET -`/api/post/idstr/available?idstr=123` - Checks if idstr is available or not/



- POST -`/api/topic/new`

- GET -`/api/topic/all`

- PUT -`/api/topic/edit?id=123`

- DELETE -`/api/topic/delete?id=123`

