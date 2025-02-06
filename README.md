Here’s the entire `README.md` properly formatted for Markdown, ensuring that **all text** is aligned with Markdown syntax for a professional and clean display:

```markdown
# 🎨 ASCII Art Project

## 📖 Overview

The **ASCII Art Project** is a Go-based application that converts strings into **graphical ASCII art representations**. The program uses pre-defined templates stored in text files, allowing for stylized output of user-provided text.

---

## ✨ Features

- 🖋 Converts input strings into ASCII art using templates.
- 🛠 Supports letters, numbers, spaces, special characters, and newlines (`\n`).
- ⚡ Includes an automated testing script to validate program functionality.

---

## 📂 File Structure

```
.
├── main.go           # Main Go program logic.
├── main_test.go      # Unit tests for the program.
├── ascii_audit.sh    # Bash script for automated testing.
├── standard.txt      # ASCII templates (Standard font).
├── shadow.txt        # ASCII templates (Shadow font).
├── thinkertoy.txt    # ASCII templates (Thinkertoy font).
├── README.md         # Documentation (this file).
├── go.mod            # Go module dependencies.
└── other files       # Supporting files (if applicable).
```

---

## ⚙️ How It Works

1. **Input Handling**:
   - Accepts user input via the command line.
2. **Character Mapping**:
   - Maps each character in the input to its ASCII art representation using the specified template.
3. **Output**:
   - Prints the ASCII art line by line to the terminal.

---

## 🚀 Getting Started

### 📋 Prerequisites

- **Go** installed on your machine.
  - Download: [Go's Official Site](https://golang.org/dl/)

### 🏗 Setup and Build

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Ramona-Ekanayake/ascii-art.git
   cd ascii-art
   ```

2. **Build the Program**:
   ```bash
   go build -o ascii-art main.go
   ```

3. **Run the Program**:
   ```bash
   ./ascii-art "Your Text Here"
   ```
   Replace `"Your Text Here"` with your desired input.

---

## 🧪 Testing

### 🔍 Unit Tests
Run the unit tests to validate the program logic:
```bash
go test
```

### 🔧 Audit Script
Use the `ascii_audit.sh` script to test multiple inputs automatically.

#### 📜 How to Run
1. Make the script executable:
   ```bash
   chmod +x ascii_audit.sh
   ```

2. Execute the script:
   ```bash
   ./ascii_audit.sh
   ```

The script will run predefined test cases and display results for each case.

---

## 🎨 ASCII Templates

The program supports three fonts:
1. **Standard**: Defined in `standard.txt`
2. **Shadow**: Defined in `shadow.txt`
3. **Thinkertoy**: Defined in `thinkertoy.txt`

Each file contains 8-line ASCII representations for every supported character.

---

## 🌟 Example Usage

### Input:
```bash
./ascii-art "Hello, World!"
```

#### Output:
```plaintext
 _    _      _ _         __        __         _     _ 
| |  | |    | | |        \ \      / /        | |   | |
| |__| | ___| | | ___     \ \ /\ / /__  _ __ | | __| |
|  __  |/ _ \ | |/ _ \     \ V  V / _ \| '_ \| |/ _` |
| |  | |  __/ | | (_) |     \_/\_/ (_) | | | | | (_| |
|_|  |_|\___|_|_|\___( )                      \__,_|
                      |/                              
```

---

## 📜 Project Information

This project is part of **GritLab Åland Islands** by **01Edu**. It showcases foundational Go programming skills and ASCII art generation.

- 🔗 [Project Details](https://github.com/01-edu/public/tree/master/subjects/ascii-art)
- 🔗 [Audit Guide](https://github.com/01-edu/public/tree/master/subjects/ascii-art/audit)

---

## 🤝 Contributing

Contributions are always welcome! If you have ideas or find bugs:
- **Submit an Issue**: Open an issue in the repository.
- **Submit a Pull Request**: Make improvements and open a PR.

---

## 📜 License

This project is licensed under the **MIT License**. See the `LICENSE` file for more details.

---

## 💌 Contact

Have questions, suggestions, or feedback? Feel free to reach out:


---

🌟 **Thank you for checking out my ASCII Art Project!** 🌟
```

