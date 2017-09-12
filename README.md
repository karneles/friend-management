## Overview

Friend Management API

## API Endpoints:

- Add new member
```
POST https://friend-management.herokuapp.com/member
{
	"email": "test@test.com"
}
```

- Create friend connection
```
POST https://friend-management.herokuapp.com/friend/add
{
	"friends": [
		"test@test.com",
		"test2@test.com"
	]
}
```

- Retrieve friend list
```
POST https://friend-management.herokuapp.com/friend/retrieve
{
	"email": "test@test.com"
}
```

- Get common friend list
```
POST https://friend-management.herokuapp.com/friend/common
{
	"friends": [
		"test3@test.com",
		"test@test.com"
	]
}
```

- Subscribe updates
```
POST https://friend-management.herokuapp.com/update/subscribe
{
	"requestor": "test@test.com",
	"target": "test2@test.com"
}
```

- Block updates
```
DELETE https://friend-management.herokuapp.com/update/subscribe
{
	"requestor": "test@test.com",
	"target": "test2@test.com"
}
```

- Send updates
```
POST https://friend-management.herokuapp.com/update/send
{
	"sender": "test3@test.com",
	"text": "Hello world! test2@test.com"
}
```
