# Virtual File System CLI

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=blackhorseya_iscool-assessment&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=blackhorseya_iscool-assessment)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=blackhorseya_iscool-assessment&metric=coverage)](https://sonarcloud.io/summary/new_code?id=blackhorseya_iscool-assessment)

This is a command-line application for managing a virtual file system with user and file management capabilities using
Go and Cobra; it allows you to perform various operations such as registering users, creating folders, and managing
files in a structured and efficient manner.

## Features

- Register a new user with a unique username to manage personal virtual files.
- Create, delete, and rename folders for a specific user, providing an organized structure.
- Create, delete, and list files within a specified folder, ensuring efficient file management.

## Installation

1. Ensure that you have Go installed on your system; if not, you can download it
   from [Go's official website](https://golang.org/dl/).
2. Clone this repository by executing the following command in your terminal:

   ```sh
   git clone https://github.com/blackhorseya/iscool-assessment.git
   ```

3. Navigate to the project directory with the command:

   ```sh
   cd iscool-assessment
   ```

4. Build the application using Go by running:

   ```sh
   go build
   ```

5. Once built, run the application with the following command:

   ```sh
   ./iscool-assessment
   ```

## Usage

### Register a New User

To register a new user, utilize the `register` command as shown below:

```sh
./iscool-assessment register [username]
```

For example, to register a user with the username `john_doe`, execute:

```sh
./iscool-assessment register john_doe
```

This command will register a new user with the username `john_doe`, creating an entry in the virtual file system.

### Additional Commands

While the `register` command is illustrated, additional commands can be seamlessly integrated in a similar fashion using
Cobra; potential commands include, but are not limited to:

- **Create Folder**: This command would allow the creation of a new folder within a user's directory:
  ```sh
  ./iscool-assessment create-folder [username] [foldername] [description]
  ```

- **Delete Folder**: For deleting an existing folder:
  ```sh
  ./iscool-assessment delete-folder [username] [foldername]
  ```

- **Rename Folder**: To rename an existing folder:
  ```sh
  ./iscool-assessment rename-folder [username] [foldername] [new-foldername]
  ```

- **Create File**: To create a new file within a specified folder:
  ```sh
  ./iscool-assessment create-file [username] [foldername] [filename] [description]
  ```

- **Delete File**: For deleting an existing file:
  ```sh
  ./iscool-assessment delete-file [username] [foldername] [filename]
  ```

- **List Folders**: To list all folders belonging to a user, optionally sorted by name or creation date:
  ```sh
  ./iscool-assessment list-folders [username] [--sort-name|--sort-created] [asc|desc]
  ```

- **List Files**: To list all files within a folder, with sorting options:
  ```sh
  ./iscool-assessment list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
  ```

## Architecture Design Explanation

Based on the source code of the `iscool-assessment` project, the architecture design can be explained as follows:

1. **Main Module (`main.go`)**:
    - This is the entry point of the application, defining the command-line interface (CLI) and related commands.
    - It handles user input parsing and invokes the corresponding functionality modules.

2. **Virtual File System Module (`vfs.go`)**:
    - Implements the main logic for the virtual file system, including operations for creating, deleting, and renaming
      files and folders.
    - Uses structs to represent files and directories, providing methods for various file system operations.

3. **Data Persistence (`vfs_data.json`)**:
    - Utilizes a JSON file to store user and file system data, ensuring data persistence across application restarts.
    - Includes serialization and deserialization logic to read from and write to the data file.

### Detailed Explanation

1. **Main Module**:
    - Uses the `cobra` package to handle command-line arguments and commands.
    - Sets up multiple subcommands like `register`, `mkdir`, `rm`, `rename`, each corresponding to a specific operation.

2. **Virtual File System Module**:
    - The `FileSystem` struct represents the virtual file system, containing users and directory structures.
    - Provides methods such as `CreateFolder`, `DeleteFolder`, `RenameFolder` for managing the file system.
    - Internally uses map structures for efficient management and lookup of files and folders.

3. **Data Persistence**:
    - Reads from the `vfs_data.json` file at startup to load existing data.
    - Updates the `vfs_data.json` file after each operation to ensure data consistency and persistence.

This design ensures modularity, maintainability, and extensibility of the code while guaranteeing data persistence and
stability.
