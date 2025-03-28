# gg 🧠🚀  
A minimal and smart CLI wrapper for Git that makes your workflow faster and cleaner.

# ✨ What is it?

gg is a shorthand Git command-line tool written in Go. It wraps common Git operations into shorter, cleaner commands, and adds smart behavior like:

- 🧼 Auto-cleaning branch names (e.g. feature / Login Page → feature/login-page)
- ✅ Confirmations before destructive actions (push, checkout, etc.)
- 🔍 Viewing recent branches with last commit info
- 📜 Clean, readable log output

## 📥 Installation

```bash
git clone https://github.com/youruser/gg.git
cd gg
go install
```

Make sure `$GOBIN` is in your `$PATH`, then run `gg` from anywhere.


## 🚀 Commands

| Command               | Description                                | Git Equivalent             |
|-----------------------|--------------------------------------------|----------------------------|
| `gg c new feature asd`| Commit with message (`-am`)                | `git commit -am "new feature asd"`|
| `gg p`                | Push current branch                        | `git push`                 |
| `gg a .`              | Add files                                  | `git add .`                |
| `gg co branch`        | Checkout branch                            | `git checkout branch`      |
| `gg nb branch-name`   | Create new branch (auto-cleaned name)      | `git checkout -b name`     |
| `gg s`                | Show working tree status                   | `git status` (parsed)      |
| `gg l`                | Show recent commits                        | `git log` (cleaned view)   |
| `gg rb` / `gg recent` | Show recent branches with latest commit    | _(custom output)_          |


## 🧼 Smart branch names

The `gg nb` command automatically sanitizes your branch names:

```bash
gg nb feature / Login Page
# => git checkout -b feature/login-page

gg nb session timeout
# => git checkout -b session-timeout
```

It removes spaces, replaces underscores and uppercase letters, and ensures the name is valid for Git.



## 🔍 Recent branches

```bash
gg rb
```

Shows your most recently used branches, along with the last commit message and timestamp:

```
Recent branches:
- feature/login-page: Add login form - 2025-03-28 11:23
- fix/session-timeout: Fix bug - 2025-03-27 19:10
```



## 📜 Log viewer

```bash
gg l
```

Displays recent commits in a clean format:

```
Author:  Leandro Tocalini  
Date:    Fri, 28 Mar 2025 13:50:00 UTC  
Message: Fix logout redirect  
Commit:  abc1234
```


## 🛡️ Confirmation prompts

Before running commands like `push`, `commit`, `checkout`, or `add`, `gg` will confirm:

```
Command to execute: git push  
Proceed? [y/N]:
```

Safer, less error-prone Git workflows.

## 🧱 Tech stack

- [Go](https://golang.org/)
- [go-git](https://github.com/go-git/go-git) – Pure Go Git implementation


## 📄 License

MIT

## 👤 Author

Leandro Tocalini Joerg — [@leandrotocalini](https://github.com/leandrotocalini)
