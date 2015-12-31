# first-go-app
test app for deployment

#Steps

SSH into the machine <br/>
dokku apps:create &ltapp-name&gt <br/>

On local machine in the app folder<br/>
git remote add dokku dokku@&ltmy-IP&gt:&ltapp-name&gt <br/>
git push dokku master
