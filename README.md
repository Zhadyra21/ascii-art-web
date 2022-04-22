## ASCII Art Web

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of our last project, [ascii-art](https://github.com/01-edu/public/tree/master/subjects/ascii-art).

**The webpage allows the use of the different banners:**

    • shadow
    • standard
    • thinkertoy

**The main page has:**

    • text input
    • radio buttons, select object or anything else to switch between banners
    • button, which sends a POST request to '/ascii-art' and outputs the result on the page.

**Usage to run :**
1. Run on the Terminal:`go run server.go`
2. Open **localhost:8080** on your browser

**Usage to run with docker :**
1. Run on the Terminal: `./run-docker.sh`
2. Open **localhost:8080** on your browser
Warning : you'll need to have running docker desktop on your PC

### Examples
- `standard`:
```brainfuck
 _    _          _   _           _  
| |  | |        | | | |         | | 
| |__| |   ___  | | | |   ___   | | 
|  __  |  / _ \ | | | |  / _ \  | | 
| |  | | |  __/ | | | | | (_) | |_| 
|_|  |_|  \___| |_| |_|  \___/  (_) 
                                    
                                    
```

- `shadow`:
```brainfuck
                                    
_|    _|          _| _|          _| 
_|    _|   _|_|   _| _|   _|_|   _| 
_|_|_|_| _|_|_|_| _| _| _|    _| _| 
_|    _| _|       _| _| _|    _|    
_|    _|   _|_|_| _| _|   _|_|   _| 
                                    
                                    
```

- `thinkertoy`:
```brainfuck
                   
o  o     o o       
|  |     | |     o 
O--O o-o | | o-o | 
|  | |-' | | | | o 
o  o o-o o o o-o   
                 O 
                   
```







