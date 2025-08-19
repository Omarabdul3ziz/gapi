package main

func main() {
	// logging.Init(zerolog.LevelInfoValue)

	// manager := substrate.NewManager("wss://tfchain.grid.tf/ws")
	// substrateConnection, err := manager.Substrate()
	// if err != nil {
	// 	panic(err)
	// }

	// header, err := substrateConnection.Subscribe()
	// if err != nil {
	// 	panic(err)
	// }

	// for {
	// 	head := <-header.Chan()
	// 	fmt.Printf("New block: %d\n", head.Number)

	// 	events, errs := substrateConnection.GetEventsForBlock(uint32(head.Number))
	// }
}
