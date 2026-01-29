# LazyPorts

> **The Modern Visual Port Manager for Linux**

![License](https://img.shields.io/badge/license-MIT-blue.svg) ![Go](https://img.shields.io/badge/Go-1.24+-00ADD8.svg)

`lazyports` is a powerful Terminal UI (TUI) tool designed to visualize and manage your network ports effortlessly. built with **Bubble Tea** and **Lipgloss**, it brings a modern, interactive experience to the terminal.

## 🚀 Features

-   **Interactive Table**: Navigate through open ports with a clean, responsive UI.
-   **Visual Status**: Instantly see `LISTEN` (●) vs `ESTAB` (↔) states.
-   **Process Management**: Kill process blocking ports directly from the list.
-   **Auto-Refresh**: state updates automatically after actions.
-   **Zero Config**: Auto-detects shell (bash/zsh) and configures PATH.

## 📥 Installation

Install in seconds with a single command:

```bash
curl -sL https://raw.githubusercontent.com/v9mirza/lazyports/main/install.sh | bash
```

> **Note**: This will install `lazyports` to your system. It may ask for your password to install to `/usr/local/bin` for global access.

## 🎮 Usage

Run the tool from anywhere:

```bash
lazyports
```

### Controls

| Key | Action |
| :--- | :--- |
| `↑` / `↓` | Navigate the list |
| `k` | Kill the selected process |
| `r` | Refresh the list manually |
| `q` | Quit application |

## 🛠️ Built With

-   [Bubble Tea](https://github.com/charmbracelet/bubbletea) - The TUI framework
-   [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling definitions
-   [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components

## 📄 License

MIT
