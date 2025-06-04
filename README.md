# Go Notes Application

A simple, clean, and modular CLI note-taking application written in Go.

## Project Overview

This project demonstrates practical Go programming skills by building a command-line notes app that allows users to create, display, and save notes as JSON files. The app emphasizes:

- **Modular design** with multiple packages (`main` and `note`)
- **Structs and methods** for data encapsulation
- **Error handling** and input validation
- **File operations** including JSON serialization
- **User-friendly CLI interaction**

---

## Features

- Prompt users to input a note **title** and **content**
- Validate input and handle errors gracefully
- Display note content in a clear, formatted output
- Save notes locally as JSON files with auto-generated filenames based on the title
- Timestamp each note on creation

---

## Code Structure

- `main` package: Handles user input and orchestrates note creation, display, and saving.
- `note` package: Defines the `Note` struct, its methods (`DisplayNote`, `SaveNote`), and constructor function `NewNote`.

---

## Why This Project?

- Demonstrates strong Go fundamentals: packages, structs, methods, JSON handling, and file I/O.
- Shows ability to design clean, maintainable, modular code.
- Practical example of CLI app development.
- Ready foundation for more complex note-taking or data persistence apps.

---

## How to Use

1. Clone the repository:
   ```bash
   git clone https://github.com/omartamer002/Go-Notes.git
   cd Go-Notes
