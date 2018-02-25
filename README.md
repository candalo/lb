# lb
lb (Lazy Backup) is a tool to automate backups

## Important notes

I create this project firstly because I'm lazy and, secondly, to learn Go Programming Language. Because of this, this is just an experimental project.
Even so, I plan to expand the project to cover other cloud storage services.

All suggestions are welcomed!

## How to use

#### 1. Turn on the Drive API
Follow the step 1 of the guide: https://developers.google.com/drive/v3/web/quickstart/go

#### 2. Put the `.client_secret.json` file in the home folder
The file must be hidden

#### 3. Run `lb my_file` or `lb my_folder\`
The first time, it is necessary to authenticate using browser. Copy the link generated after run and paste it on your browser.

#### 4. Give necessary permissions
Just follow the orientations on your browser

#### 5. Paste generated code in the terminal
Copy the generated code and paste it in the terminal

#### 6. Enjoy the powers of laziness!

## How to get

`go get github.com/candalo/lb`

## TODO

[ ] - Add tests

[ ] - Add alarms for backups

[ ] - Cover other cloud storage services
