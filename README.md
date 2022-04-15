# goget

goget is a simple tool which can be used to download a file concurrently with split method.
It splits a single file into multiple parts and downloads all of them concurrently. Once download
is completed, it rejoins them into a single file.

## Commandline structure
This tool has the following command options
```shell
Usage: goget [options]
Options:
-u          Url of the file to be downloaded
-t          Number of concurrent threads to be used
-o          File path where file will be downloaded
-n          Name of file with which it will be created
```