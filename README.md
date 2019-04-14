# LPG-Bot
A bot multitask for my Discord Server named : **LaPlanqueGaming**
<img src="/LPGBot/config/LogoBot_alpha.png" alt="lpgbot logo" align="right" width="300" height="244"/>
## Current features

 - When you type `!hello` LPG Bot answers you
 - You can now play sounds on your discord with `!sd` or `!dit` <sound> cmd (see sounds below)
 > boi / bruh / fuck / mgs / nice / ooh / oui / thug and wow
 - You can join the RicardoGame ! With the command `!ricardo` !
 > The ricardo game allows you to earn new server tags !
 - Try your video game words knowledge with the `!hangman` cmd
 > Enter `!h + <guess>` to send a guess
 - LPG Bot send you chuck norris facts if you type `!chuck`
 - Play to head and tails with the `!flip` command
 - Launch a dice with the `!roll` command, good luck !
 - The commands `!weather` / `!wt` / `!meteo` or `!mt` gives you the weather of the 2 next days

 - And you can type `!help` if you need some help

 ## Todo list / Roadmap
- [x] Fix GuildRoleAdd / GuildRoleRemove functions, in RicardoGame
> Solution was to put "Bot role" above others guild's roles
- [x] Fix role earned display, in RicardoGame
- [x] Add a new game (Hangman ?)
- [ ] Add a "Choose your own adventure game ?
- [ ] Refact the code (especially RicardoGame) ?

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