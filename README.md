# MOPS - Maintainance Operations

A simple CLI tool to backup a configuration file.

![Demo MOPS](./mops.gif)

## Features

- Backup a file: `A.foo` -> `A.foo.YYYYMMDD.backup`
  
  E.g: `nginx.conf` -> `nginx.conf.20221016.backup`

- Store multiple versions of the backup files into `backup_versions.json`. This file will be used for the file restore feature, __available soon__!

## Usage

- Clone the repository and `cd` into it:
  ``` bash
  git clone https://github.com/thainmuet/mops.git
  cd mops
  ```

- Back up a file:
  ``` bash
  ./mops backup path/to/file
  ```
  Example:
  ``` bash
  # /etc/nginx/nginx.conf requires`sudo` privileage to make any changes
  sudo ./mops backup /etc/nginx/nginx.conf 
  ```

  Result:
  ``` bash
  $ ls -l /etc/nginx/
  ``` 

  ```
  total 100
  ...
  -rw-r--r--. 1 root root  2334 May 25 19:34 nginx.conf
  -rw-------. 1 root root  2334 Oct 16 03:41 nginx.conf.20221016.backup
  ...
  ```

## Note
- Currently, the default permission for the created backup file is `0600` or `-rw-------`.
