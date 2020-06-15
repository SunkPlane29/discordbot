# Discordbot
 A very basic discord music bot

-------------------------------------

### Requirements:

* Go
* Youtube-dl - for dowloading music from youtube (must be in PATH)
* Ffmpeg - for converting youtube videos to .mp3 files (must be in PATH)
* Libraries:
  * DiscordGo - use `go get github.com/bwmarrin/discordgo` (https://github.com/bwmarrin/discordgo)
  * Dca - use `go get github.com/jonas747/dca` (https://github.com/jonas747/dca)
  
  
 ### Getting Started:
 
 * Setting up an application and a bot in discord and also adding it to your server (https://www.digitaltrends.com/gaming/how-to-make-a-discord-bot/)
 * You must have youtube-dl executable installed (https://youtube-dl.org/) and in your PATH environmental variable and also ffmpeg executable installed (https://ffmpeg.org/) and in PATH.
 * The youtube API is used for finding the videos and you need GOOGLE_API_KEY environmental variable with a valid key (https://www.slickremix.com/docs/get-api-key-for-youtube/).
 * For default the audio directory will be on root/assets/audios, it will create that directory if it doesn't exists. To change it you must change the -o argment iin the youtube-dl.conf file and also in the play command (I'll be later creating a configuration global file).
 * If you want to create an executable you can create a bin directory in the root directory and set GOBIN to the directory you just created.
 
 ### Running the code:

 * Type `go run main.go -t="(BOT TOKEN)"` if using an executable `name of executable -t="(BOT TOKEN)"`. The bot token must be inside "".
 * Keep in mind that this bot isn't fast, so it takes some 30s++ to start playing the song as it needs to make downloads and conversions. Also you don't want to play two musics in the same server.
 
 ### List of commands:
 
 * `!summon` - makes the bot join your current voice channel.
 * `!disband` - makes the bot leave it's current voice channel.
 * `!play NAME OF THE SONG` - plays the fist song that apears.
 
 ----------------------------------------------------
 
 #### Special thanks to who's behind DiscordGo library and dca library <3.
