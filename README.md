# AmazeCC CLI

<p align="center">
  <img src="https://img.shields.io/badge/AmazeCC_CLI-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="AmazeCC CLI">
</p>

<p align="center">
  <strong>Access AmazeCC from your terminal — grades, attendance, and more</strong>
</p>

<p align="center">
  <a href="https://github.com/AmazeContinuityProjects/AmazeCC-CLI"><strong>Repository</strong></a>
</p>

<p align="center">
  <img src="https://img.shields.io/github/last-commit/AmazeContinuityProjects/AmazeCC-CLI/main?style=flat-square&label=Last%20Commit" alt="Last Commit">
  <img src="https://img.shields.io/github/repo-size/AmazeContinuityProjects/AmazeCC-CLI?style=flat-square&label=Repo%20Size&color=blueviolet" alt="Repo Size">
  <br>
  <img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
</p>

---

## Overview

AmazeCC CLI is a native command-line tool that lets you interact with the AmazeCC student ecosystem directly from your terminal. Check your grades, attendance, timetable, and more without opening a browser.

---

## Features

- 📊 View your current grades and CGPA
- 📋 Check attendance percentages
- 📅 Browse your class timetable
- 🔐 Secure credential storage
- ⚡ Lightning fast — built in Go

---

## Tech Stack

| Category | Technology |
|----------|------------|
| **Language** | Go |
| **API** | AmazeCC Backend API |

---

## Installation

```bash
# Clone the repository
git clone https://github.com/AmazeContinuityProjects/AmazeCC-CLI.git
cd AmazeCC-CLI

# Build
go build -o amazecc

# Run
./amazecc
```

---

## Usage

```
amazecc [command]

Available commands:
  grades       Fetch your current grades
  attendance   Check attendance status
  timetable    View class schedule
  login        Authenticate with VTOP credentials
  help         Show this help message
```

---

## Contributing

Contributions are welcome! Feel free to fork the repo and submit a pull request.

---

## License

MIT
