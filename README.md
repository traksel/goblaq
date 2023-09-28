# GOBLAQ
Goblaq is a simple command-line interface (CLI) for monitoring the status of web applications. Easy to set up and deploy, Goblaq allows you to start tracking the status of web applications even in the early stages of infrastructure development, when more feature-rich monitoring solutions are not yet deployed or configured.

## INSTALLATION
Requirements:
- Linux
- Go 1.19+
- make

To install Goblaq:
1. Clone the repository to your machine and navigate to the directory.
2. Run make install in the console.
3. Add the services you want to monitor. Run goblaq watch help in the console and follow the instructions.
4. Start the Goblaq daemon for real-time operation with goblaq daemon daemon, or run it in the background with make daemon.
5. Wait for some time while Goblaq gathers statistics, then execute goblaq status all.

## EXAMPLE
|NAME |URL|STATUS|TIMESTAMP|MESSAGE|
|-----|---|------|---------|-------|
|app01|https://app01.io/health|200|Fri Sep 29 00:04:27 MSK 2023|200 OK|
|app02|https://app02.io/health|503|Fri Sep 29 00:04:28 MSK 2023|503 Service Unavailable|
|app03|https://app03.io/health|401|Fri Sep 29 00:04:32 MSK 2023|401 Unauthorized|