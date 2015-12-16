package chap1

func Main(args []string) {
	excercise := args[2]
	switch excercise {

	case "echo1":
		echo1(args)

	case "echo2":
		echo2(args)

	case "ex1_1":
		ex1_1(args)

	case "ex1_2":
		ex1_2(args)

	case "ex1_3":
		ex1_3(args)

	case "dup1":
		dup1()

	case "dup2":
		dup2()

	case "dup2_1":
		dup2_1()

	default:
		println("not found: ", excercise)
	}
}
