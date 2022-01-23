# TELKOM TEST BACKEND DEVELOPER
This repo for answer task backend developer telkom
*note this program just handle file .log nginx (error.log or access.log) for convert json file

### How to use Command Line Tools
To get tools help, we run command `-h`
```
./cli-convert-log.exe -h
```

For the response: 
```
  -o string
        a string with value example [path] -t [json/text] -o [path] || [path] -o [path]
  -t string                                                                
        a string with value json(default) or text example [path] -t [json/text] (default "json")
```

Example use command for `-t`
```
./cli-convert-log.exe ./file-test/error.log -t json 
```
or 
```
./cli-convert-log.exe ./file-test/error.log -t text
```

Example use command for `-o`
```
./cli-convert-log.exe ./file-test/error.log -o ./file-test/error.json
```
with `-t`
```
./cli-convert-log.exe ./file-test/error.log -t json -o ./file-test/error.json
```