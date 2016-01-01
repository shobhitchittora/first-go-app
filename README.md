# first-go-app
test app for deployment

#Steps

SSH into the machine <br/>
> dokku apps:create <app-name> <br/>

On local machine in the app folder<br/>
> git remote add dokku dokku@<my-IP> : <app-name> <br/>
<br/>
> git push dokku master

#Procfile

> web: < app-folder-name>
