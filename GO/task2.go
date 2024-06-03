package main

import (
    "fmt"
    "math/rand"
    "time"
)

func generateMasterKey(masterKey *[]byte, mt *rand.Rand) {
    letters := "abcdefghijklmnopqrstuvwxyz"
    numbers := "0123456789"
    for i := 0; i < 8; i++ {
        *masterKey = append(*masterKey, byte(letters[mt.Intn(len(letters))]))
        *masterKey = append(*masterKey, byte(numbers[mt.Intn(len(numbers))]))
    }
}

func shiftRows(line []byte) {
    v := make([]byte, len(line))
    for i := 1; i < len(line); i++ {
        v[i-1] = line[i]
    }
    v[len(line)-1] = line[0]
    copy(line, v)
}

func invShiftRows(line []byte) {
    v := make([]byte, len(line))
    for i := range line {
        v[i] = line[(i+i%(i%4)*i%4)%len(line)]
    }
    copy(line, v)
}

func subBytes(line []byte) {
    sbox := [256]byte{
        0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76,
        // ...
    }
    for i, b := range line {
        line[i] = sbox[b]
    }
}

func invSubBytes(line []byte) {
    invSbox := [256]byte{
        0x52, 0x09, 0x6a, 0xd5, 0x30, 0x36, 0xa5, 0x38, 0xbf, 0x40, 0xa3, 0x9e, 0x81, 0xf3, 0xd7, 0xfb,
        // ...
    }
    for i, b := range line {
        line[i] = invSbox[b]
    }
}

func addRoundKey(first, second []byte) []byte {
    v := make([]byte, len(first))
    for i := range first {
        v[i] = first[i] ^ second[i]
    }
    return v
}

func keyExpansion(masterKey []byte, roundKeys *[][]byte) {
    Nk, Nb, Nr := 4, 4, 10
    *roundKeys = make([][]byte, Nb*(Nr+1))
    for i := range *roundKeys {
        (*roundKeys)[i] = make([]byte, 4)
    }

    i := 0
    for ; i < Nk; i++ {
        copy((*roundKeys)[i], masterKey[4*i:4*i+4])
    }

    rcon := []byte{
        0x00, 0x00, 0x00, 0x00,
        0x01, 0x00, 0x00, 0x00,
        // ...
    }

    for ; i < Nb*(Nr+1); i++ {
        tmp := make([]byte, 4)
        copy(tmp, (*roundKeys)[i-1])
        if i%Nk == 0 {
            shiftRows(tmp)
            subBytes(tmp)
            for j := range tmp {
                tmp[j] ^= rcon[i/Nk*4+j]
            }
        } else if Nk > 6 && i%Nk == 4 {
            subBytes(tmp)
        }
        for j := range tmp {
            (*roundKeys)[i][j] = (*roundKeys)[i-Nk][j] ^ tmp[j]
        }
    }
}

func galoisMultiply(a, b byte) byte {
    var result byte
    var carry byte

    for i := 0; i < 8; i++ {
        if b&1 == 1 {
            result ^= a
        }

        carry = a & 0x80
        a <<= 1
        if carry != 0 {
            a ^= 0x1b
        }

        b >>= 1
    }

    return result
}

func generateMultBy2() []byte {
    v := make([]byte, 256)
    for i := range v {
        v[i] = galoisMultiply(byte(i), 2)
    }
    return v
}

func generateMultBy3() []byte {
    v := make([]byte, 256)
    for i := range v {
        v[i] = galoisMultiply(byte(i), 3)
    }
    return v
}

func generateMultBy14() []byte {
    v := make([]byte, 256)
    for i := range v {
        v[i] = galoisMultiply(byte(i), 14)
    }
    return v
}

func generateMultBy9() []byte {
    v := make([]byte, 256)
    for i := range v {
        v[i] = galoisMultiply(byte(i), 9)
    }
    return v
}

func generateMultBy13() []byte {
    v := make([]byte, 256)
    for i := range v {
        v[i] = galoisMultiply(byte(i), 13)
    }
    return v
}

func generateMultBy11() []byte {
    v := make([]byte, 256)
    for i := range v {
        v[i] = galoisMultiply(byte(i), 11)
    }
    return v
}

func mixColumns(line []byte) {
    multBy2 := generateMultBy2()
    multBy3 := generateMultBy3()

    v := make([]byte, 4)
    v[0] = multBy2[line[0]] ^ multBy3[line[1]] ^ line[2] ^ line[3]
    v[1] = multBy2[line[1]] ^ multBy3[line[2]] ^ line[0] ^ line[3]
    v[2] = multBy2[line[2]] ^ multBy3[line[3]] ^ line[0] ^ line[1]
    v[3] = multBy2[line[3]] ^ multBy3[line[0]] ^ line[1] ^ line[2]
    copy(line, v)
}

func invMixColumns(line []byte) {
    multBy14 := generateMultBy14()
    multBy9 := generateMultBy9()
    multBy13 := generateMultBy13()
    multBy11 := generateMultBy11()

    v := make([]byte, 4)
    v[0] = multBy14[line[0]] ^ multBy9[line[1]] ^ multBy13[line[2]] ^ multBy11[line[3]]
    v[1] = multBy14[line[1]] ^ multBy9[line[2]] ^ multBy13[line[3]] ^ multBy11[line[0]]
    v[2] = multBy14[line[2]] ^ multBy9[line[3]] ^ multBy13[line[0]] ^ multBy11[line[1]]
    v[3] = multBy14[line[3]] ^ multBy9[line[0]] ^ multBy13[line[1]] ^ multBy11[line[2]]
    copy(line, v)
}

func blockGenerate(text string, block *[][][]byte) {
    for len(textfunc blockGenerate(text string, block *[][][]byte) {
    for len(text) % 16 != 0 {
        text += " " // Добавляем пробелы, если длина не кратна 16
    }

    *block = make([][][]byte, 0, len(text)/16)
    sixteen := make([][]byte, 4)
    for i := range sixteen {
        sixteen[i] = make([]byte, 4)
    }

    for i := range text {
        a := i % 16 % 4
        b := i % 16 / 4
        sixteen[a][b] = text[i]

        if (i+1)%16 == 0 {
            *block = append(*block, sixteen)
            sixteen = make([][]byte, 4)
            for i := range sixteen {
                sixteen[i] = make([]byte, 4)
            }
        }
    }
}

func cipher(block, roundKeys [][]byte, masterKey []byte) [][]byte {
    v := make([][]byte, 4)
    for i := range v {
        v[i] = addRoundKey(block[i], roundKeys[i])
    }

    for i := 1; i <= 9; i++ {
        for j := 0; j <= 3; j++ {
            subBytes(v[j])
            shiftRows(v[j])
            mixColumns(v[j])
            v[j] = addRoundKey(v[j], roundKeys[i])
        }
    }

    for j := 0; j <= 3; j++ {
        subBytes(v[j])
        shiftRows(v[j])
        v[j] = addRoundKey(v[j], roundKeys[10])
    }

    return v
}

func decipher(block, roundKeys [][]byte, masterKey []byte) [][]byte {
    roundKeys = make([][]byte, 44)
    keyExpansion(masterKey, &roundKeys)

    v := make([][]byte, 4)
    for j := 0; j <= 3; j++ {
        v[j] = addRoundKey(block[j], roundKeys[10])
    }

    for i := 9; i >= 1; i-- {
        for j := 0; j <= 3; j++ {
            v[j] = addRoundKey(v[j], roundKeys[i])
            invMixColumns(v[j])
            invShiftRows(v[j])
            invSubBytes(v[j])
        }
    }

    for j := 0; j <= 3; j++ {
        v[j] = addRoundKey(v[j], roundKeys[0])
        invShiftRows(v[j])
        invSubBytes(v[j])
    }

    return v
}

func main() {
    rand.Seed(time.Now().UnixNano())
    mt := rand.New(rand.NewSource(time.Now().UnixNano()))

    var text string
    fmt.Print("Input text for cipher >> ")
    fmt.Scanln(&text)

    var block [][][]byte
    blockGenerate(text, &block)
    fmt.Println("----------------------------------------------")
    fmt.Println("The encryption block:")
    for _, b := range block {
        for _, line := range b {
            for _, c := range line {
                fmt.Printf("%4d ", c)
            }
            fmt.Println()
        }
        fmt.Println()
    }
    fmt.Println("----------------------------------------------")

    var masterKey []byte
    generateMasterKey(&masterKey, mt)
    fmt.Print("128-bit master key: ")
    for _, b := range masterKey {
        fmt.Printf("%c", b)
    }
    fmt.Println("\n----------------------------------------------")

    var roundKeys [][]byte
    keyExpansion(masterKey, &roundKeys)
    fmt.Println("Generated keys:")
    for _, line := range roundKeys {
        for _, b := range line {
            fmt.Printf("%04x ", b)
        }
        fmt.Println()
    }
    fmt.Println("----------------------------------------------")

    var prevState [][]byte
    defaultState := prevState
    var encrypted [][][]byte
    for _, b := range block {
        res := make([][]byte, 4)
        for i := range res {
            res[i] = make([]byte, 4)
        }
        cipherState := cipher(prevState, roundKeys, masterKey)
        for i := range b {
            for j := range b[i] {
                res[i][j] = cipherState[i][j] ^ b[i][j]
            }
        }
        encrypted = append(encrypted, res)
        prevState = cipherState
    }

    fmt.Println("The final cipher after encryption:")
    for _, b := range encrypted {
        for _, line := range b {
            for _, c := range line {
                fmt.Printf("%4d ", c)
            }
            fmt.Println()
        }
        fmt.Println()
    }
    fmt.Println("----------------------------------------------")

    prevState = defaultState
    var decrypted [][][]byte
    for _, b := range encrypted {
        res := make([][]byte, 4)
        for i := range res {
            res[i] = make([]byte, 4)
        }
        cipherState := cipher(prevState, roundKeys, masterKey)
        for i := range b {
            for j := range b[i] {
                res[i][j] = b[i][j] ^ cipherState[i][j]
            }
        }
        decrypted = append(decrypted, res)
        prevState = cipherState
    }

    var decrypted2 [][][]byte
    for _, b := range encrypted {
        res := make([][]byte, 4)
        for i := range res {
            res[i] = make([]byte, 4)
        }
        cipherState := decipher(b, roundKeys, masterKey)
        for i := range cipherState {
            for j := range cipherState[i] {
                res[i][j] = cipherState[i][j]
            }
        }
        decrypted2 = append(decrypted2, res)
    }

    fmt.Println("Message after decryption:")
    for _, b := range decrypted {
        for _, line := range b {
            for _, c := range line {
                fmt.Printf("%c", c)
            }
        }
    }
    fmt.Println("\n----------------------------------------------")
}

