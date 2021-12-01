# slic

Show Languages In Code. <br/>
![image](https://user-images.githubusercontent.com/74897008/144322182-11f9809d-8f18-4eab-b3da-ef5065a8602d.png)
![image](https://user-images.githubusercontent.com/74897008/144322301-e625aa07-000b-4775-b27b-4c72c460b243.png)

<!-- A program to generate an image containing stats on the languages used inside a project. -->

## Installation

NOTE: The binaries are only for x86, 64bit devices. For others, please compile from source.

Alternative, <br/>
Get the required files

```
git clone https://github.com/resyfer/slic.git
cd slic
```

### Linux

Install

```
./install.sh
```

Run with <br/>

```
slic
```

### Windows

The bin folder has the executable `slic.exe`. Copy to preferred location and run from there, or add to PATH for enhanced experience.

## Compiing From Source

Pre-requirements: `Go 1.17 or higher`

`cd` into project folder

```
rm -rf bin
mkdir bin
go build -o ./bin
cd bin
```

and run the executable binary

### Usage

Run it with an `-h` flag to list all commands.

`-d` flag can be used to specify the directory of search<br/>
`-i` flag can be used to ignore certain files or directories.
