# CLI To-Do List App

A simple Command-Line Interface (CLI) To-Do List application written in Go. The app uses flags to parse inputs, stores the to-do lists in a JSON file, and displays lists in a clean tabular format.

## Features
- Add a new to-do with a specified title.
- Delete a to-do by its index.
- Edit a to-do by specifying its index and a new title.
- List all to-dos in a tabular format with details like creation and completion times.
- Toggle the completion status of a to-do by index.
- Data persistence through JSON for storing and managing to-dos.
- Displays to-dos using a formatted table (via [github.com/aquasecurity/table](https://github.com/aquasecurity/table)).

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Wambita/cli-to-to-list.git
   cd cli-to-to-list
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Build the application**:
   ```bash
   go build -o todo
   ```

4. **Run the application**:
   ```bash
   ./todo
   ```

## Usage

Run the `todo` executable with the following flags to manage your to-do list:

### Add a To-Do
```bash
./todo -Add "Buy groceries"
```

### Delete a To-Do by Index
```bash
./todo -Del 1
```

### Edit a To-Do by Index
```bash
./todo -Edit "1:Go for a walk"
```

### List All To-Dos
```bash
./todo -List
```
Example Output:
```
┌───┬───────────┬───────────┬───────────────────────────────┬───────────────────────────────┐
│ # │   Title   │ Completed │          Created At           │         Completed At          │
├───┼───────────┼───────────┼───────────────────────────────┼───────────────────────────────┤
│ 0 │ Buy bread │ ❌        │ Sun, 26 Jan 2025 18:39:37 EAT │                               │
│ 1 │ Buy milk  │ ✅        │ Sun, 26 Jan 2025 18:40:10 EAT │ Sun, 26 Jan 2025 18:42:05 EAT │
│ 2 │ Walk      │ ❌        │ Sun, 26 Jan 2025 19:08:57 EAT │                               │
└───┴───────────┴───────────┴───────────────────────────────┴───────────────────────────────┘
```

### Toggle Completion Status
```bash
./todo -toggle 2
```

## Technologies Used
- **Go Programming Language**
- **JSON** for data persistence
- **github.com/aquasecurity/table** for displaying to-dos in a tabular format

## Key Learnings
While developing this project, I gained hands-on experience with:
- JSON marshalling and unmarshalling
- Date and time handling
- Pointers, structs, and generics
- Constructors and string manipulation
- Parsing CLI flags

## Contributions
Contributions are welcome! To contribute:

1. Fork this repository.
2. Create a new branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add some feature"
   ```
4. Push to the branch:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request.

## Author
Wambita Sheila Fana

GitHub: [Wambita](https://github.com/Wambita)

## Repository
[cli-to-to-list](https://github.com/Wambita/cli-to-to-list)

---
Happy task management! 🚀

