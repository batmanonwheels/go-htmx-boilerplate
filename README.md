# Go + HTMX Boilerplate

Everything you need to kick off your next Go + HTMX web app!

## Getting Started

Getting started is a simple as cloning the repository

```
git clone git@github.com:batmanonwheels/go-htmx-boilerplate.git
```

Changing into the new directory

```
cd go-htmx-boilerplate
```

Removing the .git folder (and any additional files, folders or dependencies you may not need)

```
rm -rf .git
```

Downloading and installing the latest standalone Tailwind CLI (No need to run "./tailwindcss init", the config file already exists!) https://tailwindcss.com/blog/standalone-cli

```
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
chmod +x tailwindcss-macos-arm64
mv tailwindcss-macos-arm64 tailwindcss
```

Congrats! You're ready to starting working on that new project!

If you'd rather run the commands above in one go, check out the command below:

```
git clone git@github.com:batmanonwheels/go-htmx-boilerplate.git &&\
cd go-htmx-boilerplate &&\
rm -rf .git &&\
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
chmod +x tailwindcss-macos-arm64
mv tailwindcss-macos-arm64 tailwindcss

```

## Development

Run your app locally with the following commands. You'll want to run each command in their own terminal.

```
./tailwindcss -i ./static/css/input.css -o ./static/css/styles.css --watch

air
```

This will run your app on localhost:8080, and the styles.css file will get updated with any css/tailwind changes made.

## Preparing for Deployment

Instructions are provided for deploying both with and without Docker. Both options still require a platform to host the application.

### Without Docker

Deploying is as easy as running

```
./tailwindcss -i ./static/css/input.css -o ./static/css/styles.css --minify &&\

go build
```

and pointing your web server to the generated `index.html` file found at `templates/index.html`

### With Docker

A Dockerfile with an [NGINX](https://www.nginx.com) base image is also provided for quick and easy deployments. Simply execute the following commands:

1. `./tailwindcss -i ./static/css/input.css -o ./static/css/styles.css --minify`
2. `go build`
3. `docker build . -t <container_name>`
   - Example: `docker build . -t todo-app`
4. `docker run  -p <port_number>:80 <container_name>`
   - Example: `docker run -p 8080:80 todo-app`
