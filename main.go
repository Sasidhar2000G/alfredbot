package main

import (
   "fmt"
   
   bt "github.com/SakoDroid/telego"
   cfg "github.com/SakoDroid/telego/configs"
   objs "github.com/SakoDroid/telego/objects"
)

func main(){

   bot, err := bt.NewBot(cfg.Default("5655859813:AAEm41cfxwd33OAMxOZgXHa9THBRu3ADm_E"))

   if err == nil{
       err == bot.Run()
       if err == nil{
           go start(bot)
       }
   }
}

func start(bot *bt.Bot){

    //The general update channel.
    updateChannel := bot.GetUpdateChannel()

   //Adding a handler. Everytime the bot receives message "hi" in a private chat, it will respond "hi to you too".
   bot.AddHandler("hi",func(u *objs.Update) {
   	_,err := bot.SendMessage(u.Message.Chat.Id,"hi to you too","",u.Message.MessageId,false,false)
   	if err != nil{
   		fmt.Println(err)
   	}
   },"private")

   //Monitores any other update. (Updates that don't contain text message "hi" in a private chat)
    for {
        update := <- updateChannel

       //Some processing on the update
    }
}
