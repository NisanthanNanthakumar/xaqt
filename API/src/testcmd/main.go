package main

import (
	"log"
	"testbox"
)

func main() {
	box := testbox.New("data/compilers.json")
	compilerTests := make(map[string]string)

	compilerTests["C++"] = "#include <iostream>\nusing namespace std;\n\nint main() {\n\tcout<<\"Hello\";\n\treturn 0;\n}"
	compilerTests["Perl"] = "use strict;\nuse warnings\n;use v5.14; say 'Hello';"
	compilerTests["Java"] = "\n\nimport java.io.*;\n\nclass myCode\n{\n\tpublic static void main (String[] args) throws java.lang.Exception\n\t{\n\t\t\n\t\tSystem.out.println(\"Hello\");\n\t}\n}"
	compilerTests["C#"] = "using System;\n\npublic class Test\n{\n\tpublic static void Main()\n\t{\n\t\t\tConsole.WriteLine(\"Hello\");\n\t}\n}"
	compilerTests["Clojure"] = "(println \"Hello\")"
	compilerTests["Go"] = "package main\nimport \"fmt\"\n\nfunc main(){\n  \n\tfmt.Printf(\"Hello\")\n}"
	compilerTests["Nodejs"] = "console.log(\"Hello\");"
	compilerTests["Python"] = "print \"Hello\""
	compilerTests["Ruby"] = "puts \"Hello\""
	compilerTests["Bash"] = "echo 'Hello' "
	// FIXME
	//compilerTests["Scala"] = "object HelloWorld {def main(args: Array[String]) = println(\"Hello\")}"
	// FIXME compiler
	//compilerTests["Rust"] = "fn main() {\n\tprintln!(\"Hello\");\n}"
	// FIXME
	//compilerTests["PHP"] = "<?php\n$ho = fopen('php://stdout', \"w\");\n\nfwrite($ho, \"Hello\");\n\n\nfclose($ho);\n"

	/*
		"MySQL":"create table myTable(name varchar(10));\ninsert into myTable values(\"Hello\");\nselect * from myTable;",
		"Objective-C": "#include <Foundation/Foundation.h>\n\n@interface Test\n+ (const char *) classStringValue;\n@end\n\n@implementation Test\n+ (const char *) classStringValue;\n{\n    return \"Hey!\";\n}\n@end\n\nint main(void)\n{\n    printf(\"%s\\n\", [Test classStringValue]);\n    return 0;\n}",
		"VB.NET": "Imports System\n\nPublic Class Test\n\tPublic Shared Sub Main() \n    \tSystem.Console.WriteLine(\"Hello\")\n\tEnd Sub\nEnd Class",
	*/

	stdin := "\n"
	expected := "Hello\n"

	for lang, code := range compilerTests {
		out := box.Test(lang, code, stdin, expected)
		//log.Println(out)
		if out[""] == true {
			//log.Printf("%s passed 'Hello' test.", lang)
		} else {
			log.Println(out)
			log.Printf("%s failed 'Hello' test.", lang)
		}
	}
}