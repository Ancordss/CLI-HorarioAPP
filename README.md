### Preview

![horario_pic.png](https://github.com/Ancordss/CLI-HorarioAPP/blob/main/pics/horario_pic.png)


### Setting-up

Clone this repository or download the package from here: ![repo](https://github.com/Ancordss/CLI-HorarioAPP.git)

enter the folder open terminal and run this command

```
-\.install.ps1
```
now run the following command

```
- [x] Get-Variable Profile
```
will show something like this: ![var_profile.png](https://github.com/Ancordss/CLI-HorarioAPP/blob/main/pics/var_profile.png)

go to the address shown in the terminal

open the file ending with ##profile.ps1

copy and paste the following command at the end of the file
```
- [x] Set-Alias h 'C:\horario\horario.exe'
```
Close and open the terminal again to run the program place h and hit enter


### build from scratch

prerequisites:

| go | 1.18.4|
|----|-------|

clone the repository enter the folder

open a terminal and put

```
- [x] go build .\horario.go
```

run it and it would be

To install it go to the section above.

!note: instructions only for windows sorry maybe later for linux
