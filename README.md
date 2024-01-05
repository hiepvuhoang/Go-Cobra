# Golang Cobra testing project

Use Golang Cobra to create scripts which using to send API related to scheduler to TestOps

## Installation and starting

Firstly, do install
```bash
go install
```

Then, do corba run. For exp:

```bash
cobra run createAndRunScheduler
```

## Development

First adding a .go file using cmd

```bash
cobra-cli add <file_name>
```
Update cmd which has just been created to put your go script or descriptions of it if needed

After it, do update function init() and use PersistentFlags() to make it happen. Don't forget to install it again and run like the above step.
