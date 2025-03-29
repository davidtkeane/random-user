```markdown
# Random User Generator

This project is a fork of [https://github.com/faressoft/random-user](https://github.com/faressoft/random-user) and has been modified to include additional features and improvements. The script generates random user data from the [Random User API](https://randomuser.me/) and displays it in the terminal. It also includes an option to export the output to a text file.

## Improvements Made

This project has been enhanced with the following features:

1. **Forked and Modified**: The original script was forked and modified by [https://github.com/davidtkeane/random-user](https://github.com/davidtkeane/random-user).
2. **Command-Line Flags**:
   - `-nat`: Filter users by nationality (e.g., `us`, `gb`, `ie`).
   - `-gender`: Filter users by gender (`male` or `female`).
   - `--text`: Export the output to a text file (`random_user.txt`).
3. **ANSI Color Removal**: Added functionality to remove ANSI color codes from the text file output for clean formatting.
4. **Improved Output**: The script displays user information in a well-structured format in the terminal and optionally saves it to a file.
5. **Dependency Management**: The script checks for required Go modules and installs them automatically if they are missing.
6. **Default Behavior**: Running the script without any flags generates a completely random user.

## How to Use

### Prerequisites

- Install Go (see instructions below).
- Ensure you have an internet connection to fetch data from the Random User API.

### Running the Script

1. Clone the repository:
   ```bash
   git clone https://github.com/davidtkeane/random-user.git
   cd random-user
   ```

2. Run the script:
   ```bash
   go run main.go
   ```

   This will generate a completely random user and display the information in the terminal.

3. Use command-line flags for additional functionality:
   - **Filter by nationality**:
     ```bash
     go run main.go -nat ie
     ```
     This will generate a random user from Ireland.

   - **Filter by gender**:
     ```bash
     go run main.go -gender male
     ```
     This will generate a random male user.

   - **Export to a text file**:
     ```bash
     go run main.go -nat ie -gender male --text
     ```
     This will generate a random male user from Ireland and save the output to random_user.txt.

### Example Output

#### Terminal Output
```plaintext
--- User Information ---
Joel

Gender: male
Name: Mr Joel Frazier

--- Location ---
Street: 6188 Victoria Road
City: Wexford
State: Galway City
Country: Ireland
Postcode: 38862
Coordinates: -56.1637 6.5999
Timezone: +7:00 Bangkok, Hanoi, Jakarta

--- Login Information ---
Email: joel.frazier@example.com
Username: organictiger145
Password: longhorn
  Password Strength: Unbreakable (as long as you don't tell anyone!)

UUID: 4c9ef55e-4852-4c6c-89e0-15e510e7ab51
ID: PPS 6592798T

--- Login Hashes ---
Salt: 7qGuiWzY
MD5: ca3cfd7976a1f47cd24516d5493455ed
SHA1: a3f96400a43ca665322e57f91f880fc7119db457
SHA256: 2e5a684fc9ed79dd703b7260213afd38e9729af1ae7c9f3c84101ffc141214df

--- Personal Details ---
Date of Birth: 1949-06-01T16:01:26Z (Age: 75)
Registered: 2008-07-22T16:57:47Z (Age: 16)
Phone: 031-108-4673
Cell: 081-125-1869

--- Picture ---
Large: https://randomuser.me/api/portraits/men/90.jpg
Medium: https://randomuser.me/api/portraits/med/men/90.jpg
Thumbnail: https://randomuser.me/api/portraits/thumb/men/90.jpg

Nationality: IE
```

#### Text File Output (`random_user.txt`)
```plaintext
--- User Information ---
Joel

Gender: male
Name: Mr Joel Frazier

--- Location ---
Street: 6188 Victoria Road
City: Wexford
State: Galway City
Country: Ireland
Postcode: 38862
Coordinates: -56.1637 6.5999
Timezone: +7:00 Bangkok, Hanoi, Jakarta

--- Login Information ---
Email: joel.frazier@example.com
Username: organictiger145
Password: longhorn
  Password Strength: Unbreakable (as long as you don't tell anyone!)

UUID: 4c9ef55e-4852-4c6c-89e0-15e510e7ab51
ID: PPS 6592798T

--- Login Hashes ---
Salt: 7qGuiWzY
MD5: ca3cfd7976a1f47cd24516d5493455ed
SHA1: a3f96400a43ca665322e57f91f880fc7119db457
SHA256: 2e5a684fc9ed79dd703b7260213afd38e9729af1ae7c9f3c84101ffc141214df

--- Personal Details ---
Date of Birth: 1949-06-01T16:01:26Z (Age: 75)
Registered: 2008-07-22T16:57:47Z (Age: 16)
Phone: 031-108-4673
Cell: 081-125-1869

--- Picture ---
Large: https://randomuser.me/api/portraits/men/90.jpg
Medium: https://randomuser.me/api/portraits/med/men/90.jpg
Thumbnail: https://randomuser.me/api/portraits/thumb/men/90.jpg

Nationality: IE
```

## Installation Instructions

### Install Go

#### Linux
1. Download the latest Go binary from [https://go.dev/dl/](https://go.dev/dl/).
2. Extract the archive:
   ```bash
   tar -C /usr/local -xzf go1.x.x.linux-amd64.tar.gz
   ```
3. Add Go to your PATH:
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```
4. Verify installation:
   ```bash
   go version
   ```

#### macOS
1. Install Go using Homebrew:
   ```bash
   brew install go
   ```
2. Verify installation:
   ```bash
   go version
   ```

#### Windows
1. Download the Go installer from [https://go.dev/dl/](https://go.dev/dl/).
2. Run the installer and follow the instructions.
3. Verify installation:
   ```cmd
   go version
   ```

### Required Modules

The script uses the following Go modules:
- `github.com/common-nighthawk/go-figure`
- `github.com/fatih/color`

The script automatically checks for these modules and installs them if they are missing.

## License

This project is licensed under the MIT License. See the original repository for more details.
