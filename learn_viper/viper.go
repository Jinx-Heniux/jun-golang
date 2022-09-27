package learn_viper

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadIni() {
	v := viper.New()

	v.AddConfigPath("./learn_viper") // 注意：相对路径是以main()所在的路径为基路径
	v.SetConfigName("config")
	v.SetConfigType("ini")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("DB username: %s\n", v.GetString("db.username"))
	fmt.Printf("DB password: %s\n", v.GetString("db.password"))
	fmt.Printf("Web username: %s\n", v.GetString("web.username"))
	fmt.Printf("Web password: %s\n", v.GetString("web.password"))
	fmt.Printf("username: %s\n", v.GetString("default.username"))
	fmt.Printf("password: %v\n", v.GetInt("default.password"))
}

func ReadYaml() {
	v := viper.New()

	v.AddConfigPath("./learn_viper") // 注意：相对路径是以main()所在的路径为基路径
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("DB username: %s\n", v.GetString("db.username"))
	fmt.Printf("DB password: %v\n", v.GetInt("db.password"))
	fmt.Printf("Web username: %s\n", v.GetString("web.username"))
	fmt.Printf("Web password: %v\n", v.GetInt("web.password"))

}
