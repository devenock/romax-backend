# 🎰 Roma X Game API

A simplistic backend API built using the **Fiber** framework to simulate the "Roma X" game. The API provides random outcomes for wins and losses and awards free games based on the game mechanics.

## 📝 Features

- Randomly determines the number of clears in a game round.
- Awards **free games** based on the number of consecutive clears.
- Returns JSON responses with game results (win/loss), clears, and free games awarded.

---

## ⚙️ Game Rules

1. **Clears and Wins**:
    - If a player clears **4 or more symbols** in a single round, they win.
    - Less than 4 clears results in a loss.

2. **Free Games**:
    - **4 clears**: 3 free games
    - **5 clears**: 5 free games
    - **6 clears**: 10 free games
    - **7 clears**: 20 free games

3. **Free Games Reusability**:
    - During free games, players can win additional free games by clearing symbols.

---

## 🚀 Getting Started

Follow these steps to set up and run the API on your local machine.

### Prerequisites

- [Go](https://go.dev/) installed on your system.
- A package manager like `go mod` for dependencies.

---

### 🛠️ Installation

1. Clone the repository:

   ```bash
   git clone git@github.com:devenock/romax-backend.git
   cd romax-backend
