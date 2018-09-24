# SPaaS (Small Product as a Service) CLI Application

CLI for the SPaaS Server

**!!! You have to have a running instance of the [SPaaS Server](https://github.com/magrandera/SPaaS) !!!**

### Prerequisites
- golang

# Install CLI

```
go get github.com/magrandera/SPaaS-cli
```

If golang is properly setup you will be able to run the CLI by executing `SPaaS-cli`

# Setting up

To get started run:

```
SPaaS-cli setup
```

It will ask you for the URL of the server so please enter the url of the spaas instance on the server including `http://` or `https://` depending on your server configuration. (Example: `https://spaas.example.com`)

After this you have to login to the server. To do this run the following command:

```
SPaaS-cli login
```

Put in your credentials and wait to be verified. Once this has gone through you are ready to use the rest of the commands.

# Commands

| Command | Explanation |
| ------- | ----------- |
| `add`     | Add a new application to be deployed |
| `change-password` | change the password of the SPaaS Server |
| `delete`  | Delete a app from the serve |
| `deploy`  | Deploy an application |
| `help`    | Help about any command |
| `inspect` | Inspect a running application |
| `list`    | List all applications created |
| `login`   | Login to SPaaS server |
| `setup`   | used to setup the cli application |
| `start`   | Start a stopped application |
| `stop`    | Stop a application |