# Development Steps for this project

These are the steps we have followed and came across while developing the project.

## Step 1: Installation of GO according to the Operating system

for `Windows System`:

- Go to the link  `https://go.dev/dl`
- Download the pre-compilled binary from the page for windows.
- Run the downloded `exe` file and follow the prompt.

for `MacOS`:

for `Linux`:

## Step 2: Creating a Project Root

Open the prefered terminal application on your system and navigate to the projects directory where you want to save your project files.

```terminal
mkdir dcabinet && cd dcabinet
```

## Step 3: Init the Go Project

Init the Backend API project inside the root of the directory with this command

```terminal
go mod init github.com/your_username/dcabinet

# eg. 
# go mod init github.com/ajaysinghnp/dcabinet
```

This will generate a `go.mod` file that contains all the metadata and information regarding the project like `project_name`, `dependencies`, `go version` etc. Basically it is equivalent to the `package.json` in the `Nodejs` environment.
