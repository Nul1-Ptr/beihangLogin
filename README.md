# beihangLogin

[![Build Status](https://github.com/Izumiko/beihangLogin/actions/workflows/go.yml/badge.svg)](https://github.com/Izumiko/beihangLogin/actions/workflows/go.yml)

北航校园网登录客户端

## 编译
```
git clone https://github.com/Izumiko/beihangLogin
cd beihangLogin
go build -o beihangLogin main.go
```

或访问[Github Actions](https://github.com/Izumiko/beihangLogin/actions/workflows/go.yml) 下载对应平台的预编译程序
## 用法

```
A command line auth tool for BUAA

Usage:
  beihangLogin [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  detect      Detect online status and try to login
  help        Help about any command
  login       Login using your username and password.
  logout      Logout your account
  status      Get current online info.

Flags:
      --debug   Enable debug mode.
  -h, --help    help for beihangLogin

Use "beihangLogin [command] --help" for more information about a command.
```

### 登录

```
Login using your username and password.

Usage:
  beihangLogin login [flags]

Flags:
  -h, --help              help for login
  -p, --password string   Password of your account. (required)
  -u, --username string   Username of your account. (required)

Global Flags:
      --debug   Enable debug mode.
```

### 注销

```
Logout your account

Usage:
  beihangLogin logout [flags]

Flags:
  -h, --help   help for logout

Global Flags:
      --debug   Enable debug mode.
```

### 状态查看

```
Get current online info.

Usage:
  beihangLogin status [flags]

Flags:
  -h, --help   help for status

Global Flags:
      --debug   Enable debug mode.
```

### 检测状态并尝试登录

```
Detect online status and login using your username and password.

Usage:
  beihangLogin detect [flags]

Flags:
  -h, --help              help for detect
  -p, --password string   Password of your account. (required)
  -u, --username string   Username of your account. (required)

Global Flags:
      --debug   Enable debug mode.
```