package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main(){
	app:=&cli.App{
		Name:"Health checker",
		Usage:"Website check",
		Flags:[]cli.Flag{
			&cli.StringFlag{
				Name:"domain",
				Required:true,
			},
			&cli.StringFlag{
				Name:"Port",
			    Required:false,
			},

		},
		Action: func (c* cli.Context) error{
			port:=c.String("port")
			if port==""{
               port="80"
			}
            status:=Check(c.String("domain"),port)
			fmt.Println(status)
			return nil
		},
	}
	err:=app.Run(os.Args)
	if err!=nil{
		log.Fatal(err)
	}
}