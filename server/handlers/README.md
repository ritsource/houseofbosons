# API Documentation
> Version 1, released on August 15, 2019

## Authentication

### `GET - https://houseofbosons.com/api/auth/google`
Redirects the client to Google-OAuth2 authentication flow, where admin can login with their google account. Only, authorized gmail accounts are allowed.

### `GET - https://houseofbosons.com/api/auth/current_user`
Writes back data about the currently logged in user.

- 200 Response:
```json
{
  "_id": "123...",
  "email":"houseofbosons@gmail.com",
  "id":"000000..."
}
```


## Post C.R.U.D.

### `GET - https://houseofbosons.com/api/api/post/new`
Writes back data about the currently logged in user.

- Request Body:
```json
{
  "id_str": "blog-post-no-1",
  "title": "Title of Blog Post 1",
  "description": "Soem description",
  "formatted_date": "15 August, 2019",
  "doc_type": 0,
  "md_src": "https://content...",
  "html_src": "https://content...",
  "thumbnail": "https://content...",
  "topics": [],
  "is_featured": false,
  "is_public": true,
  "is_deleted": false,
  "is_series": false"
}
```

- 200 Response:
```json
{
  "_id": "123...",
  "id_str": "blog-post-no-1",
  "title": "Title of Blog Post 1",
  "description": "Soem description",
  "formatted_date": "15 August, 2019",
  "doc_type": 0,
  "md_src": "https://content...",
  "html_src": "https://content...",
  "thumbnail": "https://content...",
  "topics": [],
  "is_featured": false,
  "is_public": true,
  "is_deleted": false,
  "is_series": false,
  "sub_blogs": []
}
```


Rest maybe later


- POST -`/api/post/new`

- GET -`/api/post/all`
<!-- /api/post/all?skip=10&limit=2 -->

- GET -`/api/post/single?id=123`

- PUT -`/api/post/edit?id=123`

- DELETE -`/api/post/delete/temp?id=123`

- DELETE -`/api/post/delete/perm?id=123`



- GET -`/api/post/idstr/available?idstr=123` - Checks if idstr is available or not/



- POST -`/api/topic/new`

- GET -`/api/topic/all`

- PUT -`/api/topic/edit?id=123`

- DELETE -`/api/topic/delete?id=123`



## Error-Response:
```json
{
  "message": "422 query failed"
}
```