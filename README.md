# pm

🍀 Run commands from everywhere in your project.

## 🚀 Installation

### Via Homebrew

Wip 🚧

### Manual

You can paste the binary in your `bin` directory (e.g., on mac it's `/usr/bin/local`). \
Don't forget to grant execution permissions to the binary.

```bash
chmox +x pm
```

## 💻 Usage

`pm` allows you to easily run and manage commands on your project, according to your <a>file reference</a> which acts like the root of your project.

```bash
# Will move you to the root of your project (file reference)
# then return you to your current working directory
pm npm install express
```

Use the `--file` / `-f` flag to configure your file reference.

> By default, the value is `package.json`.

```bash
# If you're working on a Cargo-based project
pm --file cargo.lock
```

You can see your reference file by using the flag `--file` / `-f` without argument.

```bash
# Will output your reference file
pm --file
```

## 🧑‍🤝‍🧑 Contributing

To contribute, fork the repository and open a pull request with the details of your changes.

Create a branch with a [conventionnal name](https://tilburgsciencehub.com/building-blocks/collaborate-and-share-your-work/use-github/naming-git-branches/).

- fix: `bugfix/the-bug-fixed`
- features: `feature/the-amazing-feature`
- test: `test/the-famous-test`
- hotfix `hotfix/oh-my-god-bro`
- wip `wip/the-work-name-in-progress`

## 📌 Roadmap

- [ ] Homebrew installation
- [ ] Add options to navigate even faster in your project (wip...)
- [x] Rewrite script in `Go`
- [ ] Implement `pm -` in go
- [ ] Option to go to the root of the project

## 📑 License

This project is under [MIT License](LICENSE).
