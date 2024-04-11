# THORChain Pause

THORChain Pause is a tool that facilitates programmatic pausing of THORChain nodes. It provides a password-protected endpoint to pause the THORChain node, triggering the execution of `make pause` in the THORChain node directory if the correct password is provided.

## How It Works

The THORChain Pause acts as a proxy server, exposing a single endpoint `/pause/{password}`. When a request is made to this endpoint with the correct password, the bot executes the `make pause` command in the THORChain node directory. This effectively pauses the THORChain node, allowing operators to manage node operations programmatically.

## Building the Binary

To build the THORChain Pause binary, follow these steps:

1. Ensure you have Go installed on your system. If not, download and install it from the official Go website: https://golang.org/dl/

2. Clone this repository to your local machine:

```sh
git clone https://github.com/Hey/THORChain-Pause
```

Navigate to the root directory of the cloned repository:

```sh
cd THORChain-Pause
```

Build the binary using the go build command:

```sh
go build -o thorchain-pause
```

After building, you will find the thorchain-pause-bot executable file in the current directory.

3. Running the THORChain Pause
   Before running the THORChain Pause, ensure you have set the required environment variables:

```sh
cp .env.example .env
```

Then open the .env file and set the following environment variables:

- PASSWORD: The password required to authenticate requests to pause the THORChain node.
- MAKE_CWD: The directory path of the THORChain node where the make pause command will be executed.

Once the environment variables are set, you can run the THORChain Pause binary. Here's an example command:

```sh
./thorchain-pause
```

By default, the server will start on port 8080. You can access the pause endpoint at http://localhost:8080/pause/{password}.
