# STARMAN - A Simple Tool for managing start up applications

```sh
     _
 ___| |_ __ _ _ __ _ __ ___   __ _ _ __
/ __| __/ _` | '__| '_ ` _ \ / _` | '_ \
\__ \ || (_| | |  | | | | | | (_| | | | |
|___/\__\__,_|_|  |_| |_| |_|\__,_|_| |_|

v0.1.2
```

A simple yet intuitive cli tool for managing windows start up applications.

## Install

```sh
 go install github.com/elliot40404/starman@latest
```

[Find latest Release Here](https://github.com/elliot40404/starman/releases/latest)

Or just download the binary from releases page and put it in your PATH.

## Features

- [x] Add an application to startup
- [x] Remove an application from startup
- [x] List all applications in startup
- [x] Enable or disable an application from startup
- [x] Enable or Disable starman startup
- [x] Set delay time for startup

## Usage

### Available Commands

```sh
    starman add <app_name>
    starman rm <app_name>
    starman ls
    starman enable <app_name>
    starman disable <app_name>
    starman start
    starman stop
    starman delay
    starman delay <delay_time>
    starman help
```

### Add an application to startup

`add` is also aliased as `a`

```sh
starman add <app_name> <app_path>
```

```sh
starman a <app_name> <app_path>
```

### Remove an application from startup

```sh
starman rm <app_name>
```

### List all applications in startup

`ls` is also aliased as `l`

```sh
$ starman l
$ starman ls
┌─────────┬─────────┬────────────────────────────────────────┐
│ NAME    │ STATUS  │ PATH                                   │
├─────────┼─────────┼────────────────────────────────────────┤
│ SPOTIFY │ Enabled │ "C:\Users\Avishek\Desktop\Spotify.exe" │
└─────────┴─────────┴────────────────────────────────────────┘
```

### Enable an application from startup

`enable` is also aliased as `e`

```sh
starman enable <app_name>
starman e <app_name>
```

### Disable an application from startup

`disable` is also aliased as `d`

```sh
starman disable <app_name>
starman d <app_name>
```

### Stop starman from running at startup

```sh
starman stop
```

### Start starman to run at startup

```sh
starman start
```

### Check current delay time for startup

```sh
starman delay
```

### Set delay time for startup

delay time is in seconds. Default is 30 seconds

```sh
starman delay <delay_time>
```

## Defaults

- Default delay time is 30 seconds

- Config file location is `HOMEDIR/.sm/sm_startup.bat`

- When a new application is added, starman is automatically added to startup if disabled.

## LICENSE

GPL-3.0
