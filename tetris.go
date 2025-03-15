package main

import (
  "fmt"
  "github.com/mattn/go-tty"
)

const n, m = 10, 10
var board [n][m]int;
var hasGameStarted bool = false;  

func main() {

  // tty setup
  tty, err := tty.Open()
  if err != nil {
    fmt.Printf("UNEXPECTED ERROR")
    return;
  }

  defer tty.Close();

  renderMenu(tty);
  startGame(tty);
}

func startGame(tty *tty.TTY) {
  clearScreen();
  fmt.Printf("in render game");
  amt := 1;

  // game loop
  for {
    inp, err := tty.ReadRune()
    if err != nil {
      fmt.Printf("ERROR")
    }
    fmt.Printf("%c was pressed", inp);
    fmt.Printf("%d", amt);
    amt++;
  }

}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func renderMenu(tty *tty.TTY) {
  fmt.Printf(`
              ░▒▓████████▓▒░▒▓████████▓▒░▒▓████████▓▒░▒▓███████▓▒░░▒▓█▓▒░░▒▓███████▓▒░ 
                 ░▒▓█▓▒░   ░▒▓█▓▒░         ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░        
                 ░▒▓█▓▒░   ░▒▓█▓▒░         ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░        
                 ░▒▓█▓▒░   ░▒▓██████▓▒░    ░▒▓█▓▒░   ░▒▓███████▓▒░░▒▓█▓▒░░▒▓██████▓▒░  
                 ░▒▓█▓▒░   ░▒▓█▓▒░         ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░ 
                 ░▒▓█▓▒░   ░▒▓█▓▒░         ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░ 
                 ░▒▓█▓▒░   ░▒▓████████▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓███████▓▒░  
                                                                                       
                                                                                       `);

  fmt.Printf("\n\n\n");

  fmt.Printf(`
  ▗▄▄▖  ▄▄▄ ▗▞▀▚▖ ▄▄▄  ▄▄▄      ▗▄▄▖▗▄▄▖ ▗▞▀▜▌▗▞▀▘▗▞▀▚▖       ■   ▄▄▄       ▄▄▄  ■  ▗▞▀▜▌ ▄▄▄ ■  
  ▐▌ ▐▌█    ▐▛▀▀▘▀▄▄  ▀▄▄      ▐▌   ▐▌ ▐▌▝▚▄▟▌▝▚▄▖▐▛▀▀▘    ▗▄▟▙▄▖█   █     ▀▄▄▗▄▟▙▄▖▝▚▄▟▌█ ▗▄▟▙▄▖
  ▐▛▀▘ █    ▝▚▄▄▖▄▄▄▀ ▄▄▄▀      ▝▀▚▖▐▛▀▘          ▝▚▄▄▖      ▐▌  ▀▄▄▄▀     ▄▄▄▀ ▐▌       █   ▐▌  
  ▐▌                           ▗▄▄▞▘▐▌                       ▐▌                 ▐▌           ▐▌  
                                                             ▐▌                 ▐▌           ▐▌  
                                                                                                 
                                                                                                 `)

  inp, err := tty.ReadRune()
  if err != nil {
    fmt.Printf("ERROR")
  }

  for inp != ' ' {
    inp, err = tty.ReadRune()
    if err != nil {
      fmt.Printf("ERROR")
    }
  }

  hasGameStarted = true;
}
