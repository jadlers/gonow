# gonow

CLI to search for trips with SL in Stockholm

## Installation

To use the tool you'll need two API-keys for "SL Platsuppslag" and "SL
Reseplanerare 3.1". Apply for the keys by creating a project at
[trafiklab.se](https://www.trafiklab.se/node/add/project).

Copy the
[`.env.example`](https://github.com/jadlers/gonow/blob/master/.env.example) file
to a new file called `.env` and add the API-keys from the project you created.

Run `make install` to install or `make build` to build the binary.
