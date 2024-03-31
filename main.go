package main

import (
    "fmt"
    "os"
)

func main() {
    var filePath string
    var choice int

    fmt.Print("filepath:    ")
    fmt.Scan(&filePath)

    rawDoc, err := read(filePath)
    if err != nil {
        fmt.Println("Error processing the file:", err)
        return
    }

    fmt.Println("[1] - encrypt\t[2] - decrypt")
    fmt.Scan(&choice)

    var errT error
    switch choice {
    case 1:
        encDoc := encrypt(rawDoc)
        errT = write(encDoc, filePath)
    case 2:
        decDoc := decrypt(rawDoc)
        errT = write(decDoc, filePath)
    }

    if errT != nil {
        fmt.Println("Error writing data to file:", errT)
        return
    }
}

func runesToString(runes []rune) string {
    return string(runes)
}

func read(path string) (string, error) {
    content, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(content), nil
}

func write(cont string, path string) error {
    return os.WriteFile(path, []byte(cont), 0644)
}

func encrypt(cont string) string {
    runes := []rune(cont)
    var eRunes []rune = make([]rune, len(runes))

    for i, cRune := range runes {
        switch cRune {
        case 'a':
            eRunes[i] = 'd'
        case 'b':
            eRunes[i] = 'e'
        case 'c':
            eRunes[i] = 'f'
        case 'd':
            eRunes[i] = 'g'
        case 'e':
            eRunes[i] = 'h'
        case 'f':
            eRunes[i] = 'i'
        case 'g':
            eRunes[i] = 'j'
        case 'h':
            eRunes[i] = 'k'
        case 'i':
            eRunes[i] = 'l'
        case 'j':
            eRunes[i] = 'm'
        case 'k':
            eRunes[i] = 'n'
        case 'l':
            eRunes[i] = 'o'
        case 'm':
            eRunes[i] = 'p'
        case 'n':
            eRunes[i] = 'q'
        case 'o':
            eRunes[i] = 'r'
        case 'p':
            eRunes[i] = 's'
        case 'q':
            eRunes[i] = 't'
        case 'r':
            eRunes[i] = 'u'
        case 's':
            eRunes[i] = 'v'
        case 't':
            eRunes[i] = 'w'
        case 'u':
            eRunes[i] = 'x'
        case 'v':
            eRunes[i] = 'y'
        case 'w':
            eRunes[i] = 'z'
        case 'x':
            eRunes[i] = 'a'
        case 'y':
            eRunes[i] = 'b'
        case 'z':
            eRunes[i] = 'c'
        case 'A':
            eRunes[i] = 'D'
        case 'B':
            eRunes[i] = 'E'
        case 'C':
            eRunes[i] = 'F'
        case 'D':
            eRunes[i] = 'G'
        case 'E':
            eRunes[i] = 'H'
        case 'F':
            eRunes[i] = 'I'
        case 'G':
            eRunes[i] = 'J'
        case 'H':
            eRunes[i] = 'K'
        case 'I':
            eRunes[i] = 'L'
        case 'J':
            eRunes[i] = 'M'
        case 'K':
            eRunes[i] = 'N'
        case 'L':
            eRunes[i] = 'O'
        case 'M':
            eRunes[i] = 'P'
        case 'N':
            eRunes[i] = 'Q'
        case 'O':
            eRunes[i] = 'R'
        case 'P':
            eRunes[i] = 'S'
        case 'Q':
            eRunes[i] = 'T'
        case 'R':
            eRunes[i] = 'U'
        case 'S':
            eRunes[i] = 'V'
        case 'T':
            eRunes[i] = 'W'
        case 'U':
            eRunes[i] = 'X'
        case 'V':
            eRunes[i] = 'Y'
        case 'W':
            eRunes[i] = 'Z'
        case 'X':
            eRunes[i] = 'A'
        case 'Y':
            eRunes[i] = 'B'
        case 'Z':
            eRunes[i] = 'C'
        default:
            eRunes[i] = cRune
        }
    }

    return runesToString(eRunes)
}

func decrypt(cont string) string {
    runes := []rune(cont)
    var dRunes []rune = make([]rune, len(runes))

    for i, cRune := range runes {
        switch cRune {
        case 'd':
            dRunes[i] = 'a'
        case 'e':
            dRunes[i] = 'b'
        case 'f':
            dRunes[i] = 'c'
        case 'g':
            dRunes[i] = 'd'
        case 'h':
            dRunes[i] = 'e'
        case 'i':
            dRunes[i] = 'f'
        case 'j':
            dRunes[i] = 'g'
        case 'k':
            dRunes[i] = 'h'
        case 'l':
            dRunes[i] = 'i'
        case 'm':
            dRunes[i] = 'j'
        case 'n':
            dRunes[i] = 'k'
        case 'o':
            dRunes[i] = 'l'
        case 'p':
            dRunes[i] = 'm'
        case 'q':
            dRunes[i] = 'n'
        case 'r':
            dRunes[i] = 'o'
        case 's':
            dRunes[i] = 'p'
        case 't':
            dRunes[i] = 'q'
        case 'u':
            dRunes[i] = 'r'
        case 'v':
            dRunes[i] = 's'
        case 'w':
            dRunes[i] = 't'
        case 'x':
            dRunes[i] = 'u'
        case 'y':
            dRunes[i] = 'v'
        case 'z':
            dRunes[i] = 'w'
        case 'a':
            dRunes[i] = 'x'
        case 'b':
            dRunes[i] = 'y'
        case 'c':
            dRunes[i] = 'z'
        case 'D':
            eRunes[i] = 'A'
        case 'E':
            eRunes[i] = 'B'
        case 'F':
            eRunes[i] = 'C'
        case 'G':
            eRunes[i] = 'D'
        case 'H':
            eRunes[i] = 'E'
        case 'I':
            eRunes[i] = 'F'
        case 'J':
            eRunes[i] = 'G'
        case 'K':
            eRunes[i] = 'H'
        case 'L':
            eRunes[i] = 'I'
        case 'M':
            eRunes[i] = 'J'
        case 'N':
            eRunes[i] = 'K'
        case 'O':
            eRunes[i] = 'L'
        case 'P':
            eRunes[i] = 'M'
        case 'Q':
            eRunes[i] = 'N'
        case 'R':
            eRunes[i] = 'O'
        case 'S':
            eRunes[i] = 'P'
        case 'T':
            eRunes[i] = 'Q'
        case 'U':
            eRunes[i] = 'R'
        case 'V':
            eRunes[i] = 'S'
        case 'W':
            eRunes[i] = 'T'
        case 'X':
            eRunes[i] = 'U'
        case 'Y':
            eRunes[i] = 'V'
        case 'Z':
            eRunes[i] = 'W'
        case 'A':
            eRunes[i] = 'X'
        case 'B':
            eRunes[i] = 'Y'
        case 'C':
            eRunes[i] = 'Z'
        default:
            dRunes[i] = cRune
        }
    }

    return runesToString(dRunes)
}
