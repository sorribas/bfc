package main

import "C"
import "bytes"

import "fmt"
import "github.com/sorribas/tcc"
import "github.com/sorribas/bfc/libtccbins"
import "io/ioutil"
import "os"

func compile(source string) string {
	var buff bytes.Buffer
	var programStart = `
	extern void putchar(char);
	extern char getchar();

	void main () {
		char array[30000];
		int index = 0;
	`
	buff.WriteString(programStart)

	bts := []byte(source)
	for _, bt := range bts {
		switch bt {
		case '>':
			buff.WriteString("index++;")
		case '<':
			buff.WriteString("index--;")
		case '+':
			buff.WriteString("array[index]++;")
		case '-':
			buff.WriteString("array[index]--;")
		case '.':
			buff.WriteString("putchar(array[index]);")
		case ',':
			buff.WriteString("array[index] = getchar();")
		case '[':
			buff.WriteString("while (array[index]) {")
		case ']':
			buff.WriteString("}")
		}
	}

	buff.WriteString("}")
	return buff.String()
}

func tmplibfolder() (string, error) {
	var err error
	slash := string(os.PathSeparator)
	tmpDir, err := ioutil.TempDir("", "bfc-tcclibs")

	libtcca, err := libtccbins.Asset("libtcc.a")
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(tmpDir+slash+"libtcc.a", libtcca, 0755)
	if err != nil {
		return "", err
	}

	libtcc1a, err := libtccbins.Asset("libtcc1.a")
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(tmpDir+slash+"libtcc1.a", libtcc1a, 0755)
	if err != nil {
		return "", err
	}

	return tmpDir, err
}

func handlePanic() {
	if err := recover(); err != nil {
		fmt.Println("bfc: error:", err)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("bfc - a brainfuck compiler")
		fmt.Println("")
		fmt.Println("usage:")
		fmt.Println("  bfc source.bf executable")
		fmt.Println("")
		fmt.Println("The executable is optional (defaults to ./a.out)")
		return
	}

	sourceBytes, err := ioutil.ReadFile(os.Args[1])
	source := string(sourceBytes)
	var dest string

	if len(os.Args) > 2 {
		dest = os.Args[2]
	} else {
		dest = "./a.out"
	}

	defer handlePanic()
	cc := tcc.NewTcc()
	cc.SetOutputType(tcc.OUTPUT_EXE)
	cc.CompileString(compile(source))
	libfolder, err := tmplibfolder()
	if err != nil {
		panic(err)
	}
	cc.SetLibPath(libfolder)
	cc.OutputFile(dest)
	err = os.RemoveAll(libfolder)
	if err != nil {
		panic(err)
	}
}
