package main

import "activity-csv/configs"

func main() {
	app := configs.New()

	configs.Catch(app.InitEnv())
	configs.Catch(app.InitService())
	configs.Catch(app.InitServer())

	app.Start()
}
