# Komuback
Komublog is a microblogging platform, but also aims to be a full fledged blogging/social platform in the future.

Our goal is to always be 100% open source and get to a point where it can be easily configured for different uses. Like forums, personal blogs 

This project is built using a combination of SvelteKit for the frontend, and PocketBase(as a framework) for the backend. This is the backend project.

You can check the running website here: [komublog](https://komu.blog)

## Features
The project is still a **work in progress** current features include:
* creating accounts
* configuring name and avatar
* posting messages
* posting images
* user pages
* limited markdown support
* tagging other users (albeit without a notification center this is quite pointless atm)

Only one custom backend feature is implemented right now:
* Likes

You could easily get away with running the standard executable for pocketbase and you would only lose out on the likes.
Having said that, more backend features will be implemented in the future.

## Running and building
If you want to run your own version of komublog you will also need to run the frontend project, you can follow the instructions here it [here](https://github.com/Komunumos/komublog-sv).

For running the backend you will need to build it, run it, and then configure it.

### Building/running in windows
Build the project and create an exe called "komuback"
```
go build -o komuback.exe
```

Run the executable with the "serve" command.
```
.\komuback.exe serve
```

### Building/running in Linux and macOS
Build the project and create an exe called "komuback"
```
go build -o komuback
```

Run the executable with the "serve" command.
```
./komuback serve
```

## Configure
1. Go to the api url that appears in the console after running
2. Create an admin user
3. Go to Import collections under Sync
4. Import pb_schema.json

Komublog requires custom columns, tables and views to be able to run. It won't run without importing the schema first.

For more information visit [PocketBase](https://pocketbase.io).

## Contributing
I don't have a defined way of contributing right now, if you are interesting contact me on twitter/email. 👀

## Credits
This project was built using the following open source libraries and services:

- [PocketBase](https://www.pocketbase.io/) - A backend as a service written in Go with a SQLite db.
- [SQLite](https://sqlite.org/) - A serverless, self-contained SQL database engine

Thanks to the authors of these libraries, couldn't have created this project without them. 🙏🏻

## Contact
You can find me on 
- twitter [@renatofontes](https://twitter.com/renatofontes)
- email renato[at]komu.blog
- komublog [@renato](https://komu.blog/@renato)