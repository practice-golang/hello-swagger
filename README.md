# Swagger practice using swaggo, echo-swagger

* Swagger `docs` page : http://localhost:4416/docs/index.html -> `Basic auth`: dev / 1234

## Build on Windows
```powershell
# Using MinGW, Powershell
cd bookshelf

# build
mingw32-make.exe

# test
mingw32-make.exe test

# clean
mingw32-make.exe clean
```

## Build on Linux
```sh
cd ./bookshelf

# build
make

# test
make test

# clean
make clean
```
