# THORChain Pause bot

Allows programmatic pausing of THORChain nodes.

## How does it work

This program is a proxy that has a password protected endpoint that allows you to pause the THORChain node, it will run `make pause` in the THORChain node directory if the password is correct.

## Build

docker run -d -p 8080:8080 -e PASSWORD=your_password_here -e MAKE_CWD=/root mygolangserver
