# password

**password** is a powerful and customizable command-line tool (CLI) built in Go using the Cobra library. It helps you generate secure passwords, manage them by saving with additional details, and supports features like bulk generation, updating, listing, and deleting passwords.

## Features

- **Password Generation**: Generate random, secure passwords of customizable lengths.
- **Password Customization**: Include numbers, special characters, or use a custom string to generate the password.
- **Bulk Generation**: Generate multiple passwords at once.
- **Password Management**: Save, update, list, and delete passwords for future use.

## Installation

1. **Clone the repository**:

   ```bash
   git clone <repository-url>
   ```

2. **Navigate to the directory**:

   ```bash
   cd go-password-gen
   ```

3. **Build the application**:

   ```bash
   go build -o password
   ```

4. **Run the application**:
   ```bash
   ./password
   ```

## Usage

### Password Generation

1. **Generate a simple password**:

   ```bash
   ./password -g
   ```

   This generates a password of default length (10 characters).

2. **Generate a password with specific length**:

   ```bash
   ./password -g -l 12
   ```

   Generates a password of 12 characters.

3. **Include numbers and special characters**:

   ```bash
   ./password -g -l 16 -d -s
   ```

   Generates a 16-character password including numbers and special characters.

4. **Generate a password using a custom string**:

   ```bash
   ./password -g -c "abcdef123"
   ```

   Generates a password using only the characters from the custom string provided.

5. **Bulk password generation**:
   ```bash
   ./password -g -b 5 -l 10
   ```
   Generates 5 passwords, each with a length of 10 characters.

### Password Management

1. **Save a password with details**:

   ```bash
   ./password save -g -l 12 -d -s --purpose "WiFi Password"
   ```

   Generates a password and saves it with a purpose (e.g., WiFi password).

2. **List all saved passwords**:

   ```bash
   ./password -h
   ```

   Lists all the saved passwords.

3. **Display all passwords (including unsaved)**:

   ```bash
   ./password -a
   ```

   Shows all passwords, even those not saved.

4. **Update the purpose of a saved password**:

   ```bash
   ./password -u <id> --purpose "Updated Purpose"
   ```

   Updates the purpose of a saved password by using its unique ID.

5. **Delete a password**:
   ```bash
   ./password -d <id>
   ```
   Deletes the password associated with the provided ID.

## Command Flags

- `-g, --generate` : Generates a password (default length is 10).
- `-l, --length` : Specify the length of the generated password (e.g., `-l 12` for a 12-character password).
- `-d, --digits` : Include digits in the generated password.
- `-s, --special` : Include special characters in the generated password.
- `-c, --custom` : Provide a custom string to generate the password from.
- `-b, --bulk` : Generate multiple passwords (e.g., `-b 10` to generate 10 passwords).
- `save` : Saves the generated password with a custom purpose (e.g., for WiFi, bank account, etc.).
- `-a, --all` : Display all passwords (saved and unsaved).
- `-h, --list` : List all saved passwords.
- `-u, --update` : Update the purpose of a saved password using its ID.
- `-d, --delete` : Delete a saved password using its ID.

## Example Commands

1. Generate a 12-character password with numbers and special characters:

   ```bash
   ./password -g -l 12 -d -s
   ```

2. Generate and save a password with a purpose:

   ```bash
   ./password save -g -l 16 -d -s --purpose "Email Account"
   ```

3. Generate 10 passwords with a length of 20 characters:

   ```bash
   ./password -g -b 10 -l 20
   ```

4. List all saved passwords:

   ```bash
   ./password -h
   ```

5. Update the purpose of a saved password:

   ```bash
   ./password -u 1 --purpose "Updated Purpose"
   ```

6. Delete a password:
   ```bash
   ./password -d 1
   ```

## Contributing

Feel free to fork this project, make your changes, and submit a pull request. Contributions are welcome!
