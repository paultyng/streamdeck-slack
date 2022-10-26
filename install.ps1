Get-Process -Name "StreamDeck" | Stop-Process -Force
Start-Sleep -Seconds 2
mkdir -p C:\Users\$env:UserName\AppData\Roaming\Elgato\StreamDeck\Plugins\com.paultyng.slack.sdPlugin
Remove-Item C:\Users\$env:UserName\AppData\Roaming\Elgato\StreamDeck\Plugins\com.paultyng.slack.sdPlugin\* -Recurse -Force
Copy-Item -Path .\com.paultyng.slack.sdPlugin\* -Destination C:\Users\$env:UserName\AppData\Roaming\Elgato\StreamDeck\Plugins\com.paultyng.slack.sdPlugin -PassThru
go build -o C:\Users\$env:UserName\AppData\Roaming\Elgato\StreamDeck\Plugins\com.paultyng.slack.sdPlugin\streamdeck-slack.exe
Start-Process -FilePath "C:\Program Files\Elgato\StreamDeck\StreamDeck.exe"