package main

func main() {
	loadDerbyConfig()
	ConnectDB()
	test()
	ServeWeb()
}

func test() {
	// log.Fatal("derbyConfig: ", iter.N(derbyConfig["laneCount"].(int)))
	// var i int
	// var r []struct{}
	// r = iter.N(9)
	// for i = range r {
	// 	fmt.Printf("n: %d\n", i+1)
	// }
	// log.Fatal("done")
}
