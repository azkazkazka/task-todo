
# API Task-Todo

A simple API made for Matchmade assessment:

## Users

### Register

```http
POST /register
```

`Headers`

| Key       | Value     |
| :-------- | :------- |
| - | - |

`Request Body`
| Key       | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `fullname` | `string` | **Required**. User's fullname |
| `username` | `string` | **Required**. User's username (unique) |
| `password` | `string` | **Required**. User's password (min. 8 chars, special characters) |
| `email`    | `string` | **Required**. User's email (unique) |

`Response`

```json
{
  "status": 201
  "data": {
    "fullname": <NAME>
    "username": <USERNAME>
    "email": <EMAIL>
  }
  "message": "Successfully registered user"
}
```

### Login

```http
POST /login
```

`Headers`

| Key       | Value     |
| :-------- | :------- |
| - | - |

`Request Body`
| Key       | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `username`      | `string` | **Required if email empty**. User's username |
| `password`      | `string` | **Required**. User's password |
| `email`      | `string` | **Required if username empty**.  User's email |

`Response`

```json
{
  "status": 200
  "data": {
    "fullname": <NAME>
    "username": <USERNAME>
    "email": <EMAIL>
  }
  "message": "Successfully logged in user"
}
```

### Get user profile

```http
GET /user
```

`Headers:`

| Key       | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <token>` |


`Request Body`

- None

`Response`

```json
{
  "status": 200
  "data": {
    "fullname": <NAME>,
    "username": <USERNAME>
    "email": <EMAIL>
  }
  "message": "Successfully get user profile"
}
```

### Update user

```http
PUT /user
```

`Headers:`

| Key       | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <token>` |

`Request Body`
| Key | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `fullname`      | `string` | User's fullname (optional) |
| `username`      | `string` | User's username (optional) |
| `email`      | `string` | User's email (optional) |

`Response`

```json
{
  "status": 200
  "data": {
    "fullname": <NAME>,
    "username": <USERNAME>
    "email": <EMAIL>
  }
  "message": "Successfully updated user profile"
}
```

### Delete user

```http
DELETE /user
```

`Headers:`

| Key       | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <token>` |

`Request Body`

- None

`Response`

```json
{
  "status": 204
  "data": {}
  "message": "Successfully deleted user profile"
}
```

## Tasks

### Get all tasks

```http
GET /tasks
```

`Headers:`

| Key       | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <token>` |

`Request Body`

- None

`Response`

```json
{
  "status": 200
  "data": [
    {
      "id": <ID>
      "title": <TITLE>
      "description": <DESCRIPTION>
      "due_date": <DUEDATE_RFC3339>
      "completion_status": <COMPLETIONSTATUS>
      "created_at": <CREATEDAT>
      "updated_at": <UPDATEDAT>
    },
    {
      "id": <ID>
      "title": <TITLE>
      "description": <DESCRIPTION>
      "due_date": <DUEDATE_RFC3339>
      "completion_status": <COMPLETIONSTATUS>
      "created_at": <CREATEDAT>
      "updated_at": <UPDATEDAT>
    },
    ...
  ]
  "message": "Successfully get user's to-do list"
}
```

### Get task

```http
GET /tasks/${id}
```

`Headers:`

| Key       | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <token>` |

`Request Body`

- None

`Response`

```json
{
  "status": 200
  "data": {
    "id": <ID>
    "title": <TITLE>
    "description": <DESCRIPTION>
    "due_date": <DUEDATE_RFC3339>
    "completion_status": <COMPLETIONSTATUS>
    "created_at": <CREATEDAT>
    "updated_at": <UPDATEDAT>
  }
  "message": "Successfully get to-do"
}
```

### Create task

```http
POST /tasks
```

`Headers:`

| Key       | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <token>` |

`Request Body`
| Key | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title`      | `string` | **Required**. Title of task |
| `description`      | `string` | **Required**. Description of task |
| `due_date`      | `string` | **Required**. Due date of task |

`Response`

```json
{
  "status": 201
  "data": {
    "id": <ID>
    "title": <TITLE>
    "description": <DESCRIPTION>
    "due_date": <DUEDATE_RFC3339>
    "completion_status": <COMPLETIONSTATUS>
    "created_at": <CREATEDAT>
    "updated_at": <UPDATEDAT>
  }
  "message": "Successfully update to-do"
}
```

### Update task

```http
PUT /tasks/${id}
```

`Headers:`

| Key       | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <token>` |

`Request Body`
| Key | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title`      | `string` | Title of task (optional) |
| `description`      | `string` | Description of task (optional) |
| `due_date`      | `string` | Due date of task (optional) |
| `completion_status`   | `string` | Completion status of task (optional) |

`Response`

```json
{
  "status": 200
  "data": {
    "id": <ID>
    "title": <TITLE>
    "description": <DESCRIPTION>
    "due_date": <DUEDATE_RFC3339>
    "completion_status": <COMPLETIONSTATUS>
    "created_at": <CREATEDAT>
    "updated_at": <UPDATEDAT>
  }
  "message": "Successfully update to-do"
}
```

### Delete task

```http
DELETE /tasks/${id}
```

`Headers:`

| Key       | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <token>` |

`Request Body`

- None

`Response`

```json
{
  "status": 204
  "data": {}
  "message": "Successfully deleted to-do"
}
```
