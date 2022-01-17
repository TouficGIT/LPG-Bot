# LPG-Bot
A bot multitask for my Discord Server named : **LaPlanqueGaming**
<img src="config/LogoBot_alpha.png" alt="lpgbot logo" align="right" width="300" height="244"/>
## Current features

 - When you type `!hello` LPG Bot answers you
 - You can now play sounds on your discord with `!say` <sound> cmd (see sounds below)
 > boi / bruh / coffin / damage / daniel / deja / fuck / krabs / mega
 > mgs / nani / nice / ooh / oui / ricardo / spooky / thug / wow
 - Get your Rocket League stats with the `!rl` command
 - Create a poll with `!poll`
 - LPG Bot send you chuck norris facts if you type `!chuck`
 - Get some memes from Reddit, with the command `!meme`
 - Play to head and tails with the `!flip` command
 - Launch a dice with the `!roll` command, good luck !

 - And you can type `!help` if you need some help

 ## Todo list / Roadmap
- [ ] Add a "Choose your own adventure game ?
- [ ] Add Admin commands
- [ ] Add more esport commands (like LOL stats)

## Install
Requirements: Have Go installed (https://golang.org/)
1. Clone repository
2. Run `Go build` in the clone repository
3. Create a bot on your discord developper account
4. Create the `config.json` file and place it inside the repository, following the template below
5. Execute `LPGBOT.exe`

## config.json
You'll need to create a new bot in your Discord developper account.
Then a bot token will be available. Add it to the config.json file
You need to create a '**config.json**' file where you'll save the Token of your bot.
Discord Developper: https://discordapp.com/developers/applications/

The syntax of your config file should be like :

    {
      "Token": "YOUR TOKEN",
      "BotPrefix": "!"
    }


The BotPrefix is the prefix before the bot command. Like "!command".

## Credits
To create this cool bot, I used the [discordGo](https://github.com/bwmarrin/discordgo) and [Airhorn revived](https://github.com/jbmagination/airhornrevived) librairies.

You should go check those amazing git repos ! :+1:

### @Toufic 
